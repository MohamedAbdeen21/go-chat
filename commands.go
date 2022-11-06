package main

type commandID int

const (
	CMD_NICK commandID = iota // creates incremental constants
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
	CMD_HISTORY
)

type command struct {
	id     commandID
	client *client
	args   []string
}
