package invade

import (
	"fmt"

	"os"
)

var alienMoves = 10000

func StartGame(alienNum int, file string) int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("This file probably doesnt exist, Try creating a new map file with the 'map' command...")

	}
	defer f.Close()

	return alienNum
}
