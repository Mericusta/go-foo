package main

import (
	"flag"
	mysqlfoo "go-foo/src/mysql-foo"
	redisfoo "go-foo/src/redis-foo"
)

var (
	redisUrl      = flag.String("redis_url", "", "redis url")
	redisPassword = flag.String("redis_password", "", "redis password")
	redisDBidx    = flag.Int("redis_db", 0, "redis db index")
	mysqlUser     = flag.String("mysql_user", "", "mysql user")
	mysqlPassword = flag.String("mysql_password", "", "mysql password")
	mysqlUrl      = flag.String("mysql_url", "", "mysql url")
	mysqlDB       = flag.String("mysql_db", "", "mysql table")
)

func init() {
	flag.Parse()
}

func main() {
	redisfoo.SearchAndFixFromRedis(*redisUrl, *redisPassword, *redisDBidx)
	mysqlfoo.SearchAndFixFromMySQL(*mysqlUser, *mysqlPassword, *mysqlUrl, *mysqlDB)
}
