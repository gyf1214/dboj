package controller

import (
	"errors"
	"math"
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

func paginize(page int, count int) map[string]interface{} {
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

	return map[string]interface{}{
		"ll": ll, "rr": rr,
		"pages": pages, "cur": page,
	}
}
