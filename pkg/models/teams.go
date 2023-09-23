package models

type Teams struct {
	Team1 Team
	Team2 Team
}

type Team struct {
	Top     Member
	Jungle  Member
	Middle  Member
	ADC     Member
	Support Member
	Weight  int
}
