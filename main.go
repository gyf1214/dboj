package main

import (
	"flag"

	_ "github.com/gyf1214/dboj/controller"
	"github.com/gyf1214/dboj/util"
)

var (
	listen = flag.String("listen", ":8087", "listen port")
)

func main() {
	flag.Parse()
	util.ListenAndServe(*listen)
}
