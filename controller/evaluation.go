package controller

import (
	"fmt"
	"net/http"

	"github.com/gyf1214/dboj/model"
	"github.com/gyf1214/dboj/util"
	"github.com/gyf1214/dboj/view"
)

func showSubmit(w http.ResponseWriter, r *http.Request) {
}

func createSubmit(w http.ResponseWriter, r *http.Request) {
	var pid int
	util.Ensure(util.ParseForm(r, "id", &pid))

	checkUser(r, 0)
	info, err := model.GetProblemInfo(pid)
	util.Ensure(err)

	langs, err := model.ListLanguage()
	util.Ensure(err)

	data := map[string]interface{}{"problem": info, "langs": langs}
	util.Ensure(view.Submit(w, data))
}

func doCreateSubmit(w http.ResponseWriter, r *http.Request) {
	var info model.SubmitInfo
	util.Ensure(util.ParseForm(r, "id", &info.Problem, "code", &info.Code, "language", &info.Language))

	info.User = checkUser(r, 0)
	id, err := model.CreateSubmit(info)
	util.Ensure(err)

	redirect(fmt.Sprintf("/submit?id=%v", id))
}

func init() {
	util.SafeHandle("/submit", showSubmit)
	util.SafeHandle("/submit/create", createSubmit).Methods("GET")
	util.SafeHandle("/submit/create", doCreateSubmit).Methods("POST")
}
