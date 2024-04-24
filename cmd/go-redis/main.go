package main

import (
	"flag"
	redisfoo "go-foo/src/redis-foo"
)

var (
	url      = flag.String("url", "", "redis url")
	password = flag.String("password", "", "redis password")
	dbIdx    = flag.Int("db", 0, "redis db index")
)

func init() {
	flag.Parse()
}

func main() {
	redisfoo.SearchAndFix(*url, *password, *dbIdx)
}
