package controller

import (
	"fmt"
	"net/http"

	"github.com/zc-staff/dboj/model"
	"github.com/zc-staff/dboj/util"
	"github.com/zc-staff/dboj/view"
	"github.com/zc-staff/dboj/worker"
)

func showSubmit(w http.ResponseWriter, r *http.Request) {
	var id int
	util.Ensure(util.ParseForm(r, "id", &id))
	uid := checkUser(r, 0)

	info, err := model.GetSubmitInfo(id)
	util.Ensure(err)
	evals, err := model.ListEvalution(id)
	util.Ensure(err)

	data := map[string]interface{}{
		"submit": info,
		"self":   uid == info.User.ID,
		"evals":  evals,
	}

	util.Ensure(view.ShowSubmit(w, data))
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
	util.Ensure(util.ParseForm(r, "id", &info.Problem.ID, "code", &info.Code, "language", &info.Language))

	var err error
	info.User.ID = checkUser(r, 0)
	info.ID, err = model.CreateSubmit(info)
	util.Ensure(err)

	util.Ensure(worker.RunSubmition(info))
	redirect(fmt.Sprintf("/submit?id=%v", info.ID))
}

func init() {
	util.SafeHandle("/submit", showSubmit)
	util.SafeHandle("/submit/create", createSubmit).Methods("GET")
	util.SafeHandle("/submit/create", doCreateSubmit).Methods("POST")
}
