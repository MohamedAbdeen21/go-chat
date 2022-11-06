package main

import (
	"log"

	"github.com/gocql/gocql"
)

type connection struct {
	session *gocql.Session
}

var singleton *connection = nil

func (c *connection) Select(room string) (logs [][]string) {
	log := make([]string, 3)
	scanner := c.session.Query("SELECT sender, date, msg FROM messages WHERE room = ? ORDER BY date", room).Iter().Scanner()
	for scanner.Next() {
		err := scanner.Scan(&log[0], &log[1], &log[2])
		if err != nil {
			println(err.Error())
		}
		logs = append(logs, log)
	}
	return logs
}

func (c *connection) Insert(sender, room, msg, date string) {
	err := c.session.Query(`INSERT INTO messages (sender, room, msg, date)
	VALUES (?, ?, ?, ?)`, sender, room, msg, date).Exec()
	if err != nil {
		log.Println(err.Error())
	}
}

func establish() *connection {
	if singleton != nil {
		return singleton
	}

	const HOST string = "172.17.0.2"
	cluster := gocql.NewCluster(HOST)
	cluster.Keyspace = "go"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Established connection to Cassandra on host: %s", HOST)

	err = session.Query(`CREATE TABLE IF NOT EXISTS messages (
		sender text,
		room text,
		msg text,
		date text,
		PRIMARY KEY(room, date));`).Exec()
	if err != nil {
		log.Println(err.Error())
	}
	singleton = &connection{session: session}
	return singleton
}
