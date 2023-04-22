package main

import (
	"calls/data"
	"calls/utils"
	"fmt"
)

func main() {
	fmt.Println("hi")
	r := data.NewRepo()
	j := utils.NewJira()
	id, err := j.UpdateJira("jira-001", r.UpdateTeam)
	if err != nil {
	}
	fmt.Println("jira id is =", id)
	_ = r.UpdateTeam(1, 33)
	j.AFromData()

}
