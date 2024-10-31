package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

var (
	welterWeights         = [...]string{"Canelo", "Crawford", "Lenard", "Porter", "Gotti", "Ward", "Robinson"}
	welterWeightsFighters = make([]string, cap(welterWeights))
	loserBracket          []string
	player1               string
	player2               string
	winner                string
	fighterNames          string
	keepRunnig            = true
)

func main() {
	tourtamentIntoroduction()
	appendFighters()
	addCustomFighterChoice()
	for {
		fighterChoice()
		returnWinner()
		if keepRunnig == false {
			break
		}
	}

}

func tourtamentIntoroduction() {
	fmt.Printf("########################\n")
	fmt.Printf("Lets Get ready to rumble\n")
	fmt.Printf("########################\n")
}

func appendFighters() []string {
	fmt.Printf("These are the fighters that have entered the tourtament\n")
	fmt.Printf("##########################################################\n")
	copy(welterWeightsFighters, welterWeights[0:7])
	for boxers := range welterWeightsFighters {
		fmt.Printf("%v\n", welterWeightsFighters[boxers])
	}
	return welterWeightsFighters
}

func fighterChoice() (string, string) {
	fmt.Printf("###########################################\n")
	fmt.Printf("These are the Two Fighters that are up next\n")
	fmt.Print("##############################################\n")
	player1 = welterWeightsFighters[rand.IntN(len(welterWeightsFighters))]
	player2 = welterWeightsFighters[rand.IntN(len(welterWeightsFighters))]
	if strings.Contains(player1, player2) {
		fmt.Printf("These two are the same name:%+v", player2)
		player2 = welterWeightsFighters[rand.IntN(len(welterWeightsFighters))]
		fmt.Printf("%v: vs %v: \n", player1, player2)
		return player1, player2
	}
	fmt.Printf("%v: vs %v: \n", player1, player2)
	return player1, player2
}

func returnWinner() (string, bool) {
	fmt.Printf("######################################\n")
	fmt.Printf("who won: %v or %v\n", player1, player2)
	fmt.Printf("Press the number 1 for %v, press the number 2 for %v\n", player1, player2)
	fmt.Printf("#######################################\n")
	fmt.Scan(&winner)
	if winner == "1" {
		loserBracket = slices.Insert(loserBracket, len(loserBracket), player2)
		fmt.Printf("################################\n")
		fmt.Printf(" Losers so far are: %v\n", loserBracket)
		var checkit = slices.Index(welterWeightsFighters, player2)
		welterWeightsFighters = append(welterWeightsFighters[:checkit], welterWeightsFighters[checkit+1:]...)
		fmt.Printf("###################################\n")
		fmt.Printf(" figters left so far: %v\n", welterWeightsFighters)
	}
	if winner == "2" {
		loserBracket = slices.Insert(loserBracket, len(loserBracket), player1)
		fmt.Printf("################################\n")
		fmt.Printf(" Losers so far are: %v\n", loserBracket)
		var checkit = slices.Index(welterWeightsFighters, player1)
		welterWeightsFighters = append(welterWeightsFighters[:checkit], welterWeightsFighters[checkit+1:]...)
		fmt.Printf("###################################\n")
		fmt.Printf(" figters left so far: %v\n", welterWeightsFighters)
	}
	if len(welterWeightsFighters) == 1 {
		fmt.Printf("----------------------------------------------------\n")
		fmt.Printf("The new Champion is: %v\n", welterWeightsFighters[0])
		fmt.Printf("----------------------------------------------------\n")
		fmt.Printf("----------------------------------------------------\n")
		fmt.Printf("Thanks to all the fighters that entered the tourtament!!!!!!!\n")

		keepRunnig = false
	}

	return winner, keepRunnig

}

func addCustomFighterChoice() []string {
	for {
		fmt.Printf("##########################################################\n")
		fmt.Printf("Do You want to add fighters? type 1 for yes, type 2 for no\n")
		fmt.Scan(&fighterNames)
		if fighterNames == "1" {
			fmt.Printf("Add your fighters name to the tourtament\n")
			fmt.Scan(&fighterNames)
			welterWeightsFighters = append(welterWeightsFighters, fighterNames)
			for fighters := range welterWeightsFighters {
				fmt.Printf("%v\n", welterWeightsFighters[fighters])
			}
		}
		if fighterNames == "2" {
			break
		}
	}
	return welterWeightsFighters
}
