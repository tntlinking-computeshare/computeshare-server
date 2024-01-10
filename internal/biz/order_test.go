package biz

import (
	"fmt"
	"testing"
)

type Cycle2 struct {
	ID       string `json:"id,omitempty"`
	FkUserId string `json:"fkUserId,omitempty"`
	Cycle    int    `json:"cycle,omitempty"`
}

func Test(t *testing.T) {
	cycle := Cycle2{
		ID:       "1",
		FkUserId: "1",
		Cycle:    1,
	}
	cycle1 := &cycle
	cycle1.ID = "2"
	cycle1.Cycle = 2
	fmt.Println(cycle)
}
