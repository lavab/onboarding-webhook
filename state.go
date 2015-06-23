package main

import (
	"time"
)

type Timer struct {
	ID      string    `gorethink:"id"`
	Time    time.Time `gorethink:"time"`
	To      []string  `gorethink:"to"`
	From    string    `gorethink:"from"`
	Subject string    `gorethink:"subject"`
	Body    string    `gorethink:"body"`
}

type State []*Timer

func (s State) Len() int {
	return len(s)
}

func (s State) Less(i, j int) bool {
	return s[i].Time.Before(s[j].Time)
}

func (s State) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
