package controller

import (
	"errors"
	"net/http"

	"github.com/gyf1214/dboj/model"
	"github.com/gyf1214/dboj/util"
)

func redirect(url string) {
	panic(util.Redirect{URL: url, Code: util.RedirectCode})
}

func checkUser(r *http.Request, uid int) int {
	id, err := model.Authenticate(util.GetSession(r))
	util.Ensure(err)
	if id == 0 {
		redirect("/login")
	}
	if uid != 0 && id != uid {
		panic(errors.New("forbidden"))
	}
	return id
}
