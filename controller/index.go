package controller

import (
	"math"
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
	count = int(math.Ceil(float64(count) / float64(util.PageSize)))

	pages := []int{}
	ll, rr := -1, -1
	for i := page - util.PageDelta; i < page+util.PageDelta; i++ {
		if i >= 0 && i < count {
			pages = append(pages, i)
			if i < page {
				ll = page - 1
			}
			if i > page {
				rr = page + 1
			}
		}
	}

	list, err := model.ListProblem(page, 0)
	util.Ensure(err)

	data := map[string]interface{}{
		"list": list,
		"page": map[string]interface{}{
			"ll": ll, "rr": rr,
			"pages": pages, "cur": page,
		},
	}

	util.Ensure(view.Index(w, data))
}

func init() {
	util.SafeHandle("/", index)
}
