package models

import "time"

type Msg struct {
	Time time.Time
	tag  string
	msg  string
}
