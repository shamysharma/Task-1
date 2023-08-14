package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id              int    `json:id gorm:"primary_key"`
	Name            string `json: "name"`
	Leave_type      string `json: "leave_type"`
	From_date       string `json: "from_date"`
	To_date         string `json: "to_date"`
	Team_name       string `json: "team_name"`
	Sick_leave_file string `json: "sick_leave_file"`
	Reporter        string `json: "reporter"`
}