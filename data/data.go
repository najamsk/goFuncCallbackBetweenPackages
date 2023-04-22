package data

import (
	"fmt"
)

// import "calls/utils"

type repo struct {
}

func NewRepo() *repo {
	return &repo{}
}

func (r *repo) A() error {
	return nil
}

// func (r *repo) BFromUtils() error {
// 	j := utils.NewJira()
// 	j.B()
// 	fmt.Println("utils.B called")
// 	return nil
// }

func (r *repo) UpdateTeam(teamid int, projectid int) error {
	fmt.Println("data.UpdateTeam called")
	return nil
}
