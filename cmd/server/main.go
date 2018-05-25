package main

import (
	"flag"

	_ "github.com/gyf1214/dboj/controller"
	"github.com/gyf1214/dboj/util"
)

func main() {
	flag.Parse()
	util.ListenAndServe(util.ListenAddr)
}
