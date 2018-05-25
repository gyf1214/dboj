package main

import (
	"net/http"

	"github.com/gyf1214/dboj/util"
	"github.com/gyf1214/dboj/view"
)

func main() {
	view.ServeStatic()
	http.ListenAndServe(util.ListenAddr, nil)
}
