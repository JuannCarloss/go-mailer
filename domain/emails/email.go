package domain

import "time"

type Emails struct {
	To        string
	Type      string
	Timestamp time.Time
}
