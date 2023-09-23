package models

type Member struct {
	Name   string
	Weight int
	Roles  Roles
}

type Roles struct {
	Top     bool
	Jungle  bool
	Middle  bool
	ADC     bool
	Support bool
}
