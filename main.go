package main

import (
	"fmt"
	"math"
	"math/rand"
)

var action map[int]string //maps number to action (stone/paper/scissors)

var canDefeat map[int]int //map to score who can defeat whom (map[winner]=loser)
var freq map[int]float64  // map to store frequency of stones papers and scissors in each round

var p1TotalScore float64 // total scaore of 1st player
var p2TotalScore float64 // total scaore of 2nd player
var p3TotalScore float64 // total scaore of 3rd player
var p4TotalScore float64 // total scaore of 4th player

// to initialize all the global varialbles
func initialise() {
	action[0] = "stone"
	action[1] = "paper"
	action[2] = "scissors"
	canDefeat[0] = 2
	canDefeat[1] = 0
	canDefeat[2] = 1
	p1TotalScore = 0
	p2TotalScore = 0
	p3TotalScore = 0
	p4TotalScore = 0

}

func main() {
	action = make(map[int]string)
	canDefeat = make(map[int]int)
	freq = make(map[int]float64)

	initialise()
	// set number of rounds
	numberOfRounds := 50

	for i := 1; i <= numberOfRounds; i++ {
		// initialize frequency of stones papers and scissors to zero
		freq[0] = 0
		freq[1] = 0
		freq[2] = 0
		// choose randomly among stone paper and scissors for each player
		choice0 := rand.Intn(3)
		choice1 := rand.Intn(3)
		choice2 := rand.Intn(3)
		choice3 := rand.Intn(3)
		// print the choice of player in each round
		fmt.Printf("Round:%v\n", i)
		fmt.Printf("Player1:%v\n", action[choice0])
		fmt.Printf("Player1:%v\n", action[choice1])
		fmt.Printf("Player2:%v\n", action[choice2])
		fmt.Printf("Player3:%v\n", action[choice3])
		//increase the frequency o choice chosen
		freq[choice0] = freq[choice0] + 1
		freq[choice1] = freq[choice1] + 1
		freq[choice2] = freq[choice2] + 1
		freq[choice3] = freq[choice3] + 1
		// print score of each player in each round
		fmt.Printf("Score of Player1:%v\n", freq[canDefeat[choice0]])
		fmt.Printf("Score of Player2:%v\n", freq[canDefeat[choice1]])
		fmt.Printf("Score of Player3:%v\n", freq[canDefeat[choice2]])
		fmt.Printf("Score of Player3:%v\n", freq[canDefeat[choice3]])
		//add present score to total score
		p1TotalScore = p1TotalScore + freq[canDefeat[choice0]]
		p2TotalScore = p2TotalScore + freq[canDefeat[choice1]]
		p3TotalScore = p3TotalScore + freq[canDefeat[choice2]]
		p4TotalScore = p4TotalScore + freq[canDefeat[choice3]]
	}

	// find and print the winner

	maxScore := math.Max(math.Max(p1TotalScore, math.Max(p2TotalScore, p3TotalScore)), p4TotalScore)
	if p1TotalScore == maxScore {
		fmt.Printf("Player1 wins with a score of :%v\n", p1TotalScore)
	}
	if p2TotalScore == maxScore {
		fmt.Printf("Player2 wins with a score of :%v\n", p2TotalScore)
	}
	if p3TotalScore == maxScore {
		fmt.Printf("Player3 wins with a score of :%v\n", p3TotalScore)
	}
	if p4TotalScore == maxScore {
		fmt.Printf("Player4 wins with a score of :%v\n", p4TotalScore)
	}

}
