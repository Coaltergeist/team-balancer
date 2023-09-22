package models

type Pool struct {
	Members []Member
}

type Member struct {
	Name   string
	Weight int
	Roles  []bool
}
