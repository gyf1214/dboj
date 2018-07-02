package main

import (
	"flag"
	"math/rand"
	"time"

	_ "github.com/zc-staff/dboj/controller"
	"github.com/zc-staff/dboj/util"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())
	util.ListenAndServe(util.ListenAddr)
}
