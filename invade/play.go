package invade

import "fmt"

var alienMoves = 10000


func StartGame(alienNum int) int  {
	for i := 0; i < alienMoves; i++ {
		m.MoveAliens()
		m.DestroyCitiesAndAliens(i)
		// If there are no remaining aliens then end the game
		if len(m.Aliens) == 0 {
			break
		}
	}
}