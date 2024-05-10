package main

import (
	"CRUD-SERVER/init/cmd"
	"flag"
)

//서버시작 (cmd.go 받아서 구성)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	cmd.NewCmd(*configPathFlag)
}
