package main

import (
	"net/http"

	"github.com/zc-staff/dboj/util"
	"github.com/zc-staff/dboj/view"
)

func main() {
	view.ServeStatic()
	http.ListenAndServe(util.ListenAddr, nil)
}
