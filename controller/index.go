package controller

import (
	"net/http"

	"github.com/gyf1214/dboj/model"
	"github.com/gyf1214/dboj/util"
	"github.com/gyf1214/dboj/view"
)

func redirect(url string) {
	panic(util.Redirect{URL: url, Code: util.RedirectCode})
}

func index(w http.ResponseWriter, r *http.Request) {
	uid := checkUser(r, 0)
	user, err := model.GetUserInfo(uid)
	util.Ensure(err)
	util.Ensure(view.Index(w, user))
}

func init() {
	util.SafeHandle("/", index)
}
