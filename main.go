package main

import (
	"fmt"
	"math/rand"
	"time"
)

func choice() {
	fmt.Println("choisir : 1 => pierre , 2 => feuille , 3 => ciseaux")
}

func inputPlayer() int {
	choice()
	var inputPlayer int
	fmt.Scanln(&inputPlayer)
	switch inputPlayer {
	case 1:
		fmt.Println("vous avez choisi : pierre")
	case 2:
		fmt.Println("vous avez choisi : feuille")
	case 3:
		fmt.Println("vous avez choisi : ciseaux")
	}
	return inputPlayer
}

func iaPlayer() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	ia := r1.Intn(3)
	switch ia {
	case 0:
		fmt.Println("ia à choisi : pierre")
	case 1:
		fmt.Println("ia à choisi : feuille")
	case 2:
		fmt.Println("ia à choisi : ciseaux")
	}
	return ia
}

func resultat(inputPlayer int, ia int) {
	if inputPlayer == 1 && ia == 0 || inputPlayer == 2 && ia == 1 || inputPlayer == 3 && ia == 2 {
		fmt.Println("match null")
	} else if inputPlayer == 1 && ia == 2 || inputPlayer == 2 && ia == 0 || inputPlayer == 3 && ia == 1 {
		fmt.Println("vous avez gagné")
	} else {
		fmt.Println("vous avez perdu")
	}
}

func main() {
	var inputPlayer int = inputPlayer()
	var ia int = iaPlayer()
	resultat(inputPlayer, ia)
}
