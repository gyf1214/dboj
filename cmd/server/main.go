package main

import (
	"flag"
	"math/rand"
	"time"

	_ "github.com/gyf1214/dboj/controller"
	"github.com/gyf1214/dboj/util"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())
	util.ListenAndServe(util.ListenAddr)
}
