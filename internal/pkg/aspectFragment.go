package pkg

import "time"

type AspectFragment struct {
	Aspect string    `json:"aspect"`
	Value  string    `json:"value"`
	Time   time.Time `json:"time"`
}
