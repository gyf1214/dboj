package controller

import (
	"net/http"

	"github.com/gyf1214/dboj/model"
	"github.com/gyf1214/dboj/util"
	"github.com/gyf1214/dboj/view"
)

func login(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"login": true, "post": "/login"}
	util.Ensure(view.Login(w, data))
}

func register(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"login": true, "post": "/register"}
	util.Ensure(view.Login(w, data))
}

func doRegister(w http.ResponseWriter, r *http.Request) {
	var name, passwd string
	util.Ensure(util.ParseForm(r, "name", &name, "passwd", &passwd))

	_, sid, err := model.Register(name, passwd)
	util.Ensure(err)
	util.SetSession(w, sid)

	redirect("/")
}

func doLogin(w http.ResponseWriter, r *http.Request) {
	var name, passwd string
	util.Ensure(util.ParseForm(r, "name", &name, "passwd", &passwd))

	sid, err := model.Login(name, passwd)
	util.Ensure(err)
	util.SetSession(w, sid)

	redirect("/")
}

func doLogout(w http.ResponseWriter, r *http.Request) {
	util.Ensure(model.Logout(util.GetSession(r)))

	redirect("/login")
}

func profile(w http.ResponseWriter, r *http.Request) {
	mid := checkUser(r, 0)
	uid := mid
	util.ParseForm(r, "id", &uid)

	user, err := model.GetUserInfo(uid)
	util.Ensure(err)

	page := 0
	util.ParseForm(r, "page", &page)
	count, err := model.CountSubmit(uid, 0)
	util.Ensure(err)
	pages := paginize(page, count)
	list, err := model.ListSubmit(uid, 0, page)
	util.Ensure(err)

	data := map[string]interface{}{
		"self":    uid == mid,
		"user":    user,
		"submits": list,
		"page":    pages,
	}
	util.Ensure(view.Profile(w, data))
}

func init() {
	util.SafeHandle("/login", login).Methods("GET")
	util.SafeHandle("/login", doLogin).Methods("Post")
	util.SafeHandle("/register", register).Methods("GET")
	util.SafeHandle("/register", doRegister).Methods("POST")
	util.SafeHandle("/logout", doLogout)
	util.SafeHandle("/profile", profile)
}
