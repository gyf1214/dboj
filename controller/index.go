package controller

import (
	"net/http"

	"github.com/gyf1214/dboj/model"
	"github.com/gyf1214/dboj/util"
	"github.com/gyf1214/dboj/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	uid, err := model.Authenticate(util.GetSession(r))
	util.Ensure(err)

	if uid == 0 {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := model.GetUserInfo(uid)
		util.Ensure(err)
		util.Ensure(view.Index(w, user))
	}
}

func init() {
	util.SafeHandle("/", index)
}
