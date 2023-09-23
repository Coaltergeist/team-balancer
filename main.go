package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Coaltergeist/team-balancer/pkg/models"
)

const (
	Top     = "top"
	Jungle  = "jungle"
	Middle  = "middle"
	ADC     = "adc"
	Support = "support"
)

// TODO implement role

func main() {
	// set seed
	rand.Seed(time.Now().UnixNano())

	var pool models.Pool
	content, err := os.ReadFile("json/pool.json")
	if err != nil {
		log.Fatal("Error when opening file pool: ", err)
	}

	err = json.Unmarshal(content, &pool)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	var gameday models.Gameday
	content, err = os.ReadFile("json/gameday.json")
	if err != nil {
		log.Fatal("Error when opening file gameday: ", err)
	}

	err = json.Unmarshal(content, &gameday)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	var memberList []models.Member
	for _, v1 := range gameday.Players {
		for _, v2 := range pool.Members {
			if v1 == v2.Name {
				memberList = append(memberList, v2)
			}
		}
	}

	if len(memberList) != 10 {
		log.Fatal("Need 10 members")
	}

	rolenum := make([]int, len(memberList))
	for p := 0; p < len(rolenum); p++ {
		// determine how many roles a player is comfortable with
		for r := 0; r < 5; r++ {
			switch r {
			case 0:
				if memberList[p].Roles.Top {
					rolenum[p]++
				}
				break
			case 1:
				if memberList[p].Roles.Jungle {
					rolenum[p]++
				}
				break
			case 2:
				if memberList[p].Roles.Middle {
					rolenum[p]++
				}
				break
			case 3:
				if memberList[p].Roles.ADC {
					rolenum[p]++
				}
				break
			case 4:
				if memberList[p].Roles.Support {
					rolenum[p]++
				}
				break
			}
		}
		if rolenum[p] == 0 {
			log.Fatal("A member doesn't have any roles: " + memberList[p].Name)
		}
	}

	teamList := []models.Teams{}

	for i := 0; i < 1000; i++ {

		var team1, team2 models.Team
		var teams models.Teams

		// TODO: implement some amount of randomization when starting these next loops
		// solo roles
		for i := 0; i < len(memberList); i++ {
			if rolenum[i] == 1 {
				team, role := placePlayer(team1, team2, memberList[i], rolenum[i])
				switch role {
				case Top:
					if team == 1 {
						team1.Top = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Top = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Jungle:
					if team == 1 {
						team1.Jungle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Jungle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Middle:
					if team == 1 {
						team1.Middle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Middle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case ADC:
					if team == 1 {
						team1.ADC = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.ADC = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Support:
					if team == 1 {
						team1.Support = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Support = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				}
			}
		}

		// 2 roles
		for i := 0; i < len(memberList); i++ {
			if rolenum[i] == 2 {
				team, role := placePlayer(team1, team2, memberList[i], rolenum[i])
				switch role {
				case Top:
					if team == 1 {
						team1.Top = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Top = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Jungle:
					if team == 1 {
						team1.Jungle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Jungle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Middle:
					if team == 1 {
						team1.Middle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Middle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case ADC:
					if team == 1 {
						team1.ADC = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.ADC = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Support:
					if team == 1 {
						team1.Support = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Support = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				}
			}
		}

		// 3 roles
		for i := 0; i < len(memberList); i++ {
			if rolenum[i] == 3 {
				team, role := placePlayer(team1, team2, memberList[i], rolenum[i])
				switch role {
				case Top:
					if team == 1 {
						team1.Top = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Top = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Jungle:
					if team == 1 {
						team1.Jungle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Jungle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Middle:
					if team == 1 {
						team1.Middle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Middle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case ADC:
					if team == 1 {
						team1.ADC = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.ADC = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Support:
					if team == 1 {
						team1.Support = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Support = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				}
			}
		}

		// 4 roles
		for i := 0; i < len(memberList); i++ {
			if rolenum[i] == 4 {
				team, role := placePlayer(team1, team2, memberList[i], rolenum[i])
				switch role {
				case Top:
					if team == 1 {
						team1.Top = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Top = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Jungle:
					if team == 1 {
						team1.Jungle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Jungle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Middle:
					if team == 1 {
						team1.Middle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Middle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case ADC:
					if team == 1 {
						team1.ADC = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.ADC = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Support:
					if team == 1 {
						team1.Support = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Support = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				}
			}
		}

		// 5 roles
		for i := 0; i < len(memberList); i++ {
			if rolenum[i] == 5 {
				team, role := placePlayer(team1, team2, memberList[i], rolenum[i])
				switch role {
				case Top:
					if team == 1 {
						team1.Top = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Top = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Jungle:
					if team == 1 {
						team1.Jungle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Jungle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Middle:
					if team == 1 {
						team1.Middle = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Middle = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case ADC:
					if team == 1 {
						team1.ADC = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.ADC = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				case Support:
					if team == 1 {
						team1.Support = memberList[i]
						team1.Weight += memberList[i].Weight
					} else {
						team2.Support = memberList[i]
						team2.Weight += memberList[i].Weight
					}
				}
			}
		}

		teams.Team1 = team1
		teams.Team2 = team2
		if areTeamsValid(teams) {
			teamList = append(teamList, teams)
			//fmt.Printf("team " + strconv.Itoa(i) + " is valid\n")
		} else {
			//fmt.Printf("team " + strconv.Itoa(i) + " is invalid\n")
		}
	}

	bestDraft := teamList[0]
	bestDraftWeight := math.Abs(float64(bestDraft.Team1.Weight) - float64(bestDraft.Team2.Weight))

	for i := 1; i < len(teamList); i++ {
		currentDraftWeight := math.Abs(float64(teamList[i].Team1.Weight) - float64(teamList[i].Team2.Weight))
		if currentDraftWeight < bestDraftWeight {
			bestDraft = teamList[i]
			bestDraftWeight = currentDraftWeight
		}
	}

	fmt.Printf("Best Draft Found:\n" + formatTeams(bestDraft))
}

func areTeamsValid(t models.Teams) bool {
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			if t.Team1.Top.Name == "" {
				return false
			}
		case 1:
			if t.Team1.Jungle.Name == "" {
				return false
			}
		case 2:
			if t.Team1.Middle.Name == "" {
				return false
			}
		case 3:
			if t.Team1.ADC.Name == "" {
				return false
			}
		case 4:
			if t.Team1.Support.Name == "" {
				return false
			}
		}
	}

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			if t.Team2.Top.Name == "" {
				return false
			}
		case 1:
			if t.Team2.Jungle.Name == "" {
				return false
			}
		case 2:
			if t.Team2.Middle.Name == "" {
				return false
			}
		case 3:
			if t.Team2.ADC.Name == "" {
				return false
			}
		case 4:
			if t.Team2.Support.Name == "" {
				return false
			}
		}
	}

	return true
}

func placePlayer(team1, team2 models.Team, player models.Member, numRoles int) (int, string) {
	teamNum := rand.Intn(1) + 1
	teamRole := ""
	// which of their preferred roles to check 1st
	var givenRole int

	switch numRoles {
	case 1:
		// yucky onetrick
		givenRole = 0
		break
	case 2:
		// 2 role gamers
		givenRole = rand.Intn(1)
		break
	case 3:
		// 3 role kings
		givenRole = rand.Intn(2)
		break
	case 4:
		// 4 role chads
		givenRole = rand.Intn(3)
		break
	case 5:
		// 5 role gigachads
		givenRole = rand.Intn(4)
		break
	}

	placed := false

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			if player.Roles.Top {
				if givenRole == 0 {
					teamRole = Top
					placed = true
					break
				} else {
					givenRole--
				}
			}
		case 1:
			if player.Roles.Jungle {
				if givenRole == 0 {
					teamRole = Jungle
					placed = true
					break
				} else {
					givenRole--
				}
			}
		case 2:
			if player.Roles.Middle {
				if givenRole == 0 {
					teamRole = Middle
					placed = true
					break
				} else {
					givenRole--
				}
			}
		case 3:
			if player.Roles.ADC {
				if givenRole == 0 {
					teamRole = ADC
					placed = true
					break
				} else {
					givenRole--
				}
			}
		case 4:
			if player.Roles.Support {
				if givenRole == 0 {
					teamRole = Support
					placed = true
					break
				} else {
					givenRole--
				}
			}
		}
		if placed {
			break
		}
	}

	// if !placed {
	// 	return teamNum, teamRole
	// }

	switch teamRole {
	case Top:
		if isTopOpen(team1, team2) {
			switch teamNum {
			case 1:
				if team1.Top.Name != "" {
					teamNum = 2
				}
			case 2:
				if team2.Top.Name != "" {
					teamNum = 1
				}
			}
		} else {
			player.Roles.Top = false
			numRoles--
			if numRoles == 0 {
				// log.Fatal("idek top " + player.Name)
			}
			teamNum, teamRole = placePlayer(team1, team2, player, numRoles)
		}
		break
	case Jungle:
		if isJungleOpen(team1, team2) {
			switch teamNum {
			case 1:
				if team1.Jungle.Name != "" {
					teamNum = 2
				}
			case 2:
				if team2.Jungle.Name != "" {
					teamNum = 1
				}
			}
		} else {
			player.Roles.Jungle = false
			numRoles--
			if numRoles == 0 {
				// log.Fatal("idek jg " + player.Name)
			}
			teamNum, teamRole = placePlayer(team1, team2, player, numRoles)
		}
		break
	case Middle:
		if isMiddleOpen(team1, team2) {
			switch teamNum {
			case 1:
				if team1.Middle.Name != "" {
					teamNum = 2
				}
			case 2:
				if team2.Middle.Name != "" {
					teamNum = 1
				}
			}
		} else {
			player.Roles.Middle = false
			numRoles--
			if numRoles == 0 {
				// log.Fatal("idek mid " + player.Name)
			}
			teamNum, teamRole = placePlayer(team1, team2, player, numRoles)
		}
		break
	case ADC:
		if isADCOpen(team1, team2) {
			switch teamNum {
			case 1:
				if team1.ADC.Name != "" {
					teamNum = 2
				}
			case 2:
				if team2.ADC.Name != "" {
					teamNum = 1
				}
			}
		} else {
			player.Roles.ADC = false
			numRoles--
			if numRoles == 0 {
				// log.Fatal("idek adc " + player.Name)
			}
			teamNum, teamRole = placePlayer(team1, team2, player, numRoles)
		}
		break
	case Support:
		if isSupportOpen(team1, team2) {
			switch teamNum {
			case 1:
				if team1.Support.Name != "" {
					teamNum = 2
				}
			case 2:
				if team2.Support.Name != "" {
					teamNum = 1
				}
			}
		} else {
			player.Roles.Support = false
			numRoles--
			if numRoles == 0 {
				// log.Fatal("idek supp " + player.Name)
			}
			teamNum, teamRole = placePlayer(team1, team2, player, numRoles)
		}
		break
	}

	return teamNum, teamRole
}

func isTopOpen(team1, team2 models.Team) bool {
	if team1.Top.Name != "" {
		if team2.Top.Name != "" {
			return false
		}
	}

	return true
}

func isJungleOpen(team1, team2 models.Team) bool {
	if team1.Jungle.Name != "" {
		if team2.Jungle.Name != "" {
			return false
		}
	}

	return true
}

func isMiddleOpen(team1, team2 models.Team) bool {
	if team1.Middle.Name != "" {
		if team2.Middle.Name != "" {
			return false
		}
	}

	return true
}

func isADCOpen(team1, team2 models.Team) bool {
	if team1.ADC.Name != "" {
		if team2.ADC.Name != "" {
			return false
		}
	}

	return true
}

func isSupportOpen(team1, team2 models.Team) bool {
	if team1.Support.Name != "" {
		if team2.Support.Name != "" {
			return false
		}
	}

	return true
}

func formatTeams(t models.Teams) string {
	return "Teams:\n\nTeam 1 weight: " + strconv.Itoa(t.Team1.Weight) + "\t\t" + formatTeam(t.Team1) + "\n\nTeam 2 weight: " + strconv.Itoa(t.Team1.Weight) + "\t\t" + formatTeam(t.Team2)
}

func formatTeam(t models.Team) string {
	return "\n\t" + t.Top.Name + " weight: " + strconv.Itoa(t.Top.Weight) + "\n\t" + t.Jungle.Name + " weight: " + strconv.Itoa(t.Jungle.Weight) + "\n\t" + t.Middle.Name + " weight: " + strconv.Itoa(t.Middle.Weight) + "\n\t" + t.ADC.Name + " weight: " + strconv.Itoa(t.ADC.Weight) + "\n\t" + t.Support.Name + " weight: " + strconv.Itoa(t.Support.Weight)
}
