package controller

import (
	"fmt"
	"net/http"

	"github.com/gyf1214/dboj/model"
	"github.com/gyf1214/dboj/util"
	"github.com/gyf1214/dboj/view"
)

func listDiscussion(w http.ResponseWriter, r *http.Request) {
	checkUser(r, 0)
	pid := 0
	page := 0
	util.ParseForm(r, "page", &page)
	util.ParseForm(r, "id", &pid)

	var (
		err  error
		prob model.ProblemInfo
	)
	if pid != 0 {
		prob, err = model.GetProblemInfo(pid)
		util.Ensure(err)
	}

	count, err := model.CountDiscussion(0)
	util.Ensure(err)
	pages := paginize(page, count)
	list, err := model.ListDiscussion(pid, page)
	util.Ensure(err)

	data := map[string]interface{}{
		"problem": prob,
		"list":    list,
		"page":    pages,
	}
	util.Ensure(view.Discussion(w, data))
}

func createDiscussion(w http.ResponseWriter, r *http.Request) {
	uid := checkUser(r, 0)
	pid := 0
	util.Ensure(util.ParseForm(r, "id", &pid))

	prob, err := model.GetProblemInfo(pid)
	util.Ensure(err)

	data := map[string]interface{}{
		"uid": uid, "problem": prob,
		"pid":  prob.ID,
		"post": "/discussion/create",
	}

	util.Ensure(view.EditPost(w, data))
}

func doCreateDiscussion(w http.ResponseWriter, r *http.Request) {
	var info model.DiscussionInfo
	info.User.ID = checkUser(r, 0)
	util.Ensure(util.ParseForm(r, "id", &info.Problem.ID, "title", &info.Title, "content", &info.Content))

	id, err := model.CreateDiscussion(info)
	util.Ensure(err)
	redirect(fmt.Sprintf("/post?id=%v", id))
}

func showPost(w http.ResponseWriter, r *http.Request) {
	var id int
	checkUser(r, 0)
	util.Ensure(util.ParseForm(r, "id", &id))

	info, err := model.GetDiscussionInfo(id)
	util.Ensure(err)
	list, err := model.ListPost(id)
	util.Ensure(err)

	data := map[string]interface{}{"list": list, "info": info}
	util.Ensure(view.Post(w, data))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var info model.PostInfo
	info.User.ID = checkUser(r, 0)
	util.Ensure(util.ParseForm(r, "id", &info.Discussion.ID))

	var err error
	info.Discussion, err = model.GetDiscussionInfo(info.Discussion.ID)
	util.Ensure(err)

	data := map[string]interface{}{
		"post": "/post/create",
		"pid":  info.Discussion.ID,
		"info": info,
	}
	util.Ensure(view.EditPost(w, data))
}

func doCreatePost(w http.ResponseWriter, r *http.Request) {
	var info model.PostInfo
	info.User.ID = checkUser(r, 0)
	util.Ensure(util.ParseForm(r, "id", &info.Discussion.ID, "content", &info.Content))

	_, err := model.CreatePost(info)
	util.Ensure(err)
	redirect(fmt.Sprintf("/post?id=%v", info.Discussion.ID))
}

func init() {
	util.SafeHandle("/discussion", listDiscussion)
	util.SafeHandle("/discussion/create", createDiscussion).Methods("GET")
	util.SafeHandle("/discussion/create", doCreateDiscussion).Methods("POST")

	util.SafeHandle("/post", showPost)
	util.SafeHandle("/post/create", createPost).Methods("GET")
	util.SafeHandle("/post/create", doCreatePost).Methods("POST")
}
