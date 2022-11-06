package main

import (
	"net"
	"strings"
	"time"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	db := establish()

	if len(strings.Split(msg, ":")) == 1 {
		go db.Insert("server", r.name, msg, time.Now().Format(time.RFC3339Nano))
	} else {
		go db.Insert(sender.nick, r.name, msg, time.Now().Format(time.RFC3339Nano))
	}

	for addr, mem := range r.members {
		if addr != sender.conn.RemoteAddr() {
			mem.msg(msg)
		}
	}
}
