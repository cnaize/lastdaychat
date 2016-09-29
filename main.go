package main

import (
	"flag"
	"github.com/cnaize/lastdaychat/chat"
)

var defaultPort string

func init() {
	flag.StringVar(&defaultPort, "port", "8080", "default port")
}

func main() {
	flag.Parse()
	panic(chat.NewChat().Run(defaultPort))
}
