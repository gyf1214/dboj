package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	cacheMem = 32 << 20
	sidName  = "dboj-session"
)

var (
	errParse = errors.New("parse form error")
	router   = mux.NewRouter()
)

// ParseForm parse values from POST request
func ParseForm(r *http.Request, params ...interface{}) error {
	err := r.ParseMultipartForm(cacheMem)
	if err != nil {
		return err
	}

	for i := 0; i < len(params); i += 2 {
		key, ok := params[i].(string)
		if !ok {
			return errParse
		}

		vs := r.Form[key]
		if len(vs) == 0 {
			return errParse
		}
		val := vs[0]

		switch ptr := params[i+1].(type) {
		case *int:
			*ptr, err = strconv.Atoi(val)
			if err != nil {
				return errParse
			}
		case *string:
			*ptr = val
		default:
			return errParse
		}
	}

	return nil
}

// SafeHandle deals with panic
func SafeHandle(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				http.Error(w, fmt.Sprint(p), 400)
			}
		}()
		f(w, r)
	})
}

// Ensure panics if err != nil
func Ensure(err error) {
	if err != nil {
		panic(err)
	}
}

// ListenAndServe starts the server
func ListenAndServe(addr string) {
	srv := &http.Server{Handler: router, Addr: addr}
	log.Fatal(srv.ListenAndServe())
}

// GetSession gets sid from cookie
func GetSession(r *http.Request) string {
	cookie, err := r.Cookie(sidName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

// SetSession sets sid via cookie
func SetSession(w http.ResponseWriter, sid string) {
	http.SetCookie(w, &http.Cookie{Name: sidName, Value: sid})
}
