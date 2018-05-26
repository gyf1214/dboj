package controller

import (
	"net/http"

	"github.com/gyf1214/dboj/model"
	"github.com/gyf1214/dboj/util"
	"github.com/gyf1214/dboj/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	checkUser(r, 0)

	page := 0
	util.ParseForm(r, "page", &page)
	count, err := model.CountProblem(0)
	util.Ensure(err)
	pages := paginize(page, count)
	list, err := model.ListProblem(page, 0)
	util.Ensure(err)

	data := map[string]interface{}{"list": list, "page": pages}
	util.Ensure(view.Index(w, data))
}

func init() {
	util.SafeHandle("/", index)
}
