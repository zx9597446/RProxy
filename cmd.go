package main

import "flag"

type Cmd struct {
	bind string
	//remote string
	//ip     string
	queryStringKey string
}

func parseCmd() Cmd {
	var cmd Cmd
	flag.StringVar(&cmd.bind, "l", "0.0.0.0:8888", "listen on ip:port")
	//flag.StringVar(&cmd.remote, "r", "http://idea.lanyus.com:80", "reverse proxy addr")
	//flag.StringVar(&cmd.ip, "ip", "", "reverse proxy addr server ip")
	flag.StringVar(&cmd.queryStringKey, "k", "backend", "reverse proxy backend url")
	flag.Parse()
	return cmd
}
