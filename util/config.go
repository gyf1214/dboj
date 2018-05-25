package util

import "time"

// global config
const (
	ListenAddr    = ":8087"
	Database      = "dboj:dboj@tcp(127.0.0.1:3306)/dboj?charset=utf8&parseTime=true"
	SessionExpire = 3600 * time.Second
)
