package invade

import (
	"bufio"

	"fmt"

	"log"

	"math/rand"

	"os"

	"strconv"

	"strings"

	"time"
)

type city struct {
	id          int
	name        string
	directions  map[int]*city //north = 1 south =2 East = 3 West = 4
	isDestroyed bool
}
type alien struct {
	isActive bool
	location int
	turns    int
}

var cities map[int]*city
var aliens []alien
var iterator int                    // the alien counter
var citytoalien = make(map[int]int) //maps alien to a city

func MoveDirections(direction string) int { //transalator for city building

	switch direction {
	case "north":
		return 1
	case "south":
		return 2
	case "east":
		return 3
	case "west":
		return 4
	default:
		return 0
	}
}
func MoveDirectionsBack(direction int) string { //transalator for city building
	switch direction {
	case 1:
		return "north"
	case 2:
		return "south"
	case 3:
		return "east"
	case 4:
		return "west"
	default:
		return "No opposite directions"
	}
}

//This function generates aliens and randomly assings them to citites
func GenerateAlienOnMap(count int) {
	aliens = make([]alien, count+1)
	x := 0
	for i := 1; i <= count; i++ {

		//ensuring that two aliens are not assigned to the same city intially
		for citytoalien[x] != 0 || x == 0 {
			x = rand.Intn(iterator)
		}
		if x == 0 {
			x++
		}
		aliens[i] = alien{true, x, 0}
		//attach alien to city
		citytoalien[x] = i
		fmt.Println("Alien ", i, " at ", cities[x].name)
	}
}

//this generates a new city map/object
func GenerateOrGetCity(cityname string) *city {

	var gtcity city
	iterator++
	gtcity = city{iterator, cityname, nil, false}
	gtcity.directions = make(map[int]*city)
	cities[iterator] = &gtcity
	return (&gtcity)
}
func GenerateCityMap(fileName string) { //generates the city map based on the input file

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file) //reading the input file
	var line string
	cities = make(map[int]*city)
	var citiesread = make(map[string]int)

	for {

		line, err = reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		str := strings.Split(line, " ")
		if citiesread[str[0]] == 0 { //check if the city has been encountered before
			GenerateOrGetCity(str[0]) //generate if not
			citiesread[str[0]] = iterator
		}
		//handling directions

		for i := 1; i < len(str); i++ {
			st := strings.Split(str[i], "=")
			st[1] = strings.TrimRight(st[1], "\r\n") //represents city
			var neig *city
			if citiesread[st[1]] == 0 { //check and generate new city neighbor
				neig = GenerateOrGetCity(st[1])
				citiesread[st[1]] = iterator
			} else {
				neig = cities[citiesread[st[1]]]
			}
			cities[citiesread[str[0]]].directions[MoveDirections(st[0])] = neig //assigning the neighbor city
		}
		if err != nil {
			break
		}
	}
	fmt.Println("Generated Map!!") //print the generated city map
	for k, v := range cities {
		fmt.Print("key: ", k, " value: ( ", v.name)
		for s, t := range v.directions {
			fmt.Print(" ", s, " : ", t.name)
		}
		fmt.Println(")")
	}
}

func CheckEnd() bool { //checks if every alien is isActive or not

	for i := 0; i < len(aliens); i++ {
		if aliens[i].isActive == true {
			return true
		}
	}
	return false
}

//Runs simulation of alien movement
func MoveAliens() {

	//var ctemp city
	for i := 0; i < len(aliens); i++ {
		if aliens[i].isActive {
			aliens[i].turns++
			mvdir := rand.Intn(4)
			if mvdir == 0 {
				mvdir++
			}

			if cities[aliens[i].location].directions[mvdir] != nil && !cities[aliens[i].location].directions[mvdir].isDestroyed { //check if the city neighbor exist and that it is not isDestroyed
				cid := cities[aliens[i].location].directions[mvdir].id
				//if two aliens are present in the same city
				if citytoalien[cid] != 0 {
					cities[aliens[i].location].directions[mvdir].isDestroyed = true
					fmt.Println("City ", cities[aliens[i].location].directions[mvdir].name, " isDestroyed by alien", citytoalien[cid], " and alien", i)
					//render alien inisActive
					aliens[i].isActive = false
					//render  alien inisActive already present in the city
					aliens[citytoalien[cid]].isActive = false
					//this resets the number of aliens in the city to 0 after former aliens have been isDestroyed
					citytoalien[aliens[i].location] = 0
				} else {
					citytoalien[aliens[i].location] = 0
					aliens[i].location = cities[cid].id
					citytoalien[cid] = i
				}

			}

			//checking and deactivating alien when they reach 10000 moves
			if aliens[i].turns == 10000 {
				aliens[i].isActive = false
			}

		}
	}

}
func MoveTillEnd() {

	for CheckEnd() != false {
		MoveAliens()
	}
}
func Run(alienNum string) {
	//seeding the random number
	rand.Seed(time.Now().UTC().UnixNano())
	noaliens, _ := strconv.Atoi(alienNum)
	fmt.Println("alien Number:", noaliens)
	fmt.Println("city count:", iterator)
	if noaliens < iterator {
		GenerateAlienOnMap(noaliens)
		MoveTillEnd()
		fmt.Println("Remaining cities !!")
		for k, v := range cities {
			if !v.isDestroyed {
				k = k + 1
				fmt.Print(v.name)
				for s, t := range v.directions {
					fmt.Print(" ", MoveDirectionsBack(s), "=", t.name)
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Error!! Number of aliens must be less than number of cities!!")
	}

}
