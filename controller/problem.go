package controller

import (
	"fmt"
	"net/http"

	"github.com/zc-staff/dboj/model"
	"github.com/zc-staff/dboj/util"
	"github.com/zc-staff/dboj/view"
)

func showProblem(w http.ResponseWriter, r *http.Request) {
	var pid int
	util.Ensure(util.ParseForm(r, "id", &pid))
	uid := checkUser(r, 0)

	info, err := model.GetProblemInfo(pid)
	util.Ensure(err)

	page := 0
	util.ParseForm(r, "page", &page)
	count, ac, err := model.CountSubmit(uid, pid)
	util.Ensure(err)
	pages := paginize(page, count)
	list, err := model.ListSubmit(uid, pid, page)
	util.Ensure(err)

	data := map[string]interface{}{
		"problem": info,
		"edit":    uid == info.Owner.ID,
		"submits": list,
		"count":   count, "ac": ac,
		"page": pages,
	}

	util.Ensure(view.ShowProblem(w, data))
}

func createProblem(w http.ResponseWriter, r *http.Request) {
	checkUser(r, 0)

	data := map[string]string{"post": "/problem/create"}
	util.Ensure(view.EditProblem(w, data))
}

func doCreateProblem(w http.ResponseWriter, r *http.Request) {
	var info model.ProblemInfo
	util.Ensure(util.ParseForm(r, "title", &info.Title, "desc", &info.Desc))
	info.Owner.ID = checkUser(r, 0)

	pid, err := model.CreateProblem(info)
	util.Ensure(err)

	redirect(fmt.Sprintf("/problem?id=%v", pid))
}

func updateProblem(w http.ResponseWriter, r *http.Request) {
	var pid int
	util.Ensure(util.ParseForm(r, "id", &pid))

	info, err := model.GetProblemInfo(pid)
	util.Ensure(err)
	checkUser(r, info.Owner.ID)

	data := map[string]interface{}{
		"post":    fmt.Sprintf("/problem/edit?id=%v", pid),
		"problem": info,
	}
	util.Ensure(view.EditProblem(w, data))
}

func doUpdateProblem(w http.ResponseWriter, r *http.Request) {
	var info model.ProblemInfo
	util.Ensure(util.ParseForm(r, "id", &info.ID, "title", &info.Title, "desc", &info.Desc))

	owner, err := model.GetProblemOwner(info.ID)
	util.Ensure(err)
	checkUser(r, owner)

	util.Ensure(model.UpdateProblem(info))
	redirect(fmt.Sprintf("/problem?id=%v", info.ID))
}

func deleteProblem(w http.ResponseWriter, r *http.Request) {
	var pid int
	util.Ensure(util.ParseForm(r, "id", &pid))

	owner, err := model.GetProblemOwner(pid)
	util.Ensure(err)
	checkUser(r, owner)

	util.Ensure(model.DeleteProblem(pid))
	redirect("/")
}

func editDataset(w http.ResponseWriter, r *http.Request) {
	var pid int
	util.Ensure(util.ParseForm(r, "id", &pid))

	owner, err := model.GetProblemOwner(pid)
	util.Ensure(err)
	checkUser(r, owner)

	data := map[string]interface{}{"pid": pid}
	util.Ensure(view.Dataset(w, data))
}

func addDataset(w http.ResponseWriter, r *http.Request) {
	var (
		pid  int
		data model.DatasetInfo
	)
	util.Ensure(util.ParseForm(r, "id", &pid, "score", &data.Score, "input", &data.Input, "answer", &data.Answer))

	owner, err := model.GetProblemOwner(pid)
	util.Ensure(err)
	checkUser(r, owner)

	util.Ensure(model.AddDataset(pid, data))
	redirect(fmt.Sprintf("/problem?id=%v", pid))
}

func deleteDataset(w http.ResponseWriter, r *http.Request) {
}

func init() {
	util.SafeHandle("/problem", showProblem)
	util.SafeHandle("/problem/create", createProblem).Methods("GET")
	util.SafeHandle("/problem/create", doCreateProblem).Methods("POST")
	util.SafeHandle("/problem/edit", updateProblem).Methods("GET")
	util.SafeHandle("/problem/edit", doUpdateProblem).Methods("POST")
	util.SafeHandle("/problem/delete", deleteProblem).Methods("GET")
	util.SafeHandle("/problem/dataset", editDataset).Methods("GET")
	util.SafeHandle("/problem/dataset", addDataset).Methods("POST")
	util.SafeHandle("/problem/dataset/delete", deleteDataset).Methods("GET")
}
