package utils

import (
	"calls/data"
	"fmt"
)

type jira struct {
}

func NewJira() *jira {
	return &jira{}
}

func (j *jira) B() error {
	return nil
}

func (j *jira) AFromData() {
	r := data.NewRepo()
	r.A()
	fmt.Println("called data.A")
}

func (j *jira) UpdateJira(jira string, fn func(tid, pid int) error) (int, error) {
	fmt.Println("before callback")
	_ = fn(12, 21)
	fmt.Println("after callback function")
	return 1221, nil
}
