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
	id        int
	name      string
	directions map[int]*city //north = 1 south =2 East = 3 West = 4
	destroyed bool
}
type alien struct { 
	active     bool 
	location   int
	turns int
}

var cities map[int]*city            
var aliens []alien                  
var iterator int      // the alien counter              
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

func GenerateAlienOnMap(count int) { //generating aliens and assigning to random city
	aliens = make([]alien, count+1) //count of alien + 1 to avoid zero conflict
	x := 0
	for i := 1; i <= count; i++ {

		for citytoalien[x] != 0 || x == 0 { //ensuring that two aliens are not assigned to the same citi intially
			x = rand.Intn(iterator) //randomly select the aliens location out of available cities
		}
		if x == 0 { //skipping index zero
			x++
		}
		aliens[i] = alien{true, x, 0}
		citytoalien[x] = i //mapping alien to city
		fmt.Println("Alien ", i, " at ", cities[x].name)
	}
}
func GenerateOrGetCity(cityname string) *city { //generates new city object

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

func CheckEnd() bool { //checks if every alien is active or not

	for i := 0; i < len(aliens); i++ {
		if aliens[i].active == true {
			return true
		}
	}
	return false
}

func MoveAliens() {

	//var ctemp city
	for i := 0; i < len(aliens); i++ {
		if aliens[i].active {
			aliens[i].turns++
			mvdir := rand.Intn(4)
			if(mvdir == 0){
				mvdir++
			}
			
            

			if cities[aliens[i].location].directions[mvdir] != nil && !cities[aliens[i].location].directions[mvdir].destroyed { //check if the city neighbor exist and that it is not destroyed
				cid := cities[aliens[i].location].directions[mvdir].id //name of the neigbouring city
				if citytoalien[cid] != 0 {                            //two aliens in the same
					cities[aliens[i].location].directions[mvdir].destroyed = true
					fmt.Println("City ", cities[aliens[i].location].directions[mvdir].name, " Destroyed by alien", citytoalien[cid], " and alien", i)
					aliens[i].active = false                //destroy the alien
					aliens[citytoalien[cid]].active = false //destroy the alien already present in the city
					citytoalien[aliens[i].location] = 0     //reset the number of aliens in the current city
				} else { //if there are no conflicts move the alien to the new location and reset the old locations
					citytoalien[aliens[i].location] = 0
					aliens[i].location = cities[cid].id
					citytoalien[cid] = i
				}

			}

			if aliens[i].turns == 10000 { //checking and deactivating alien when they reach 10000 moves
				aliens[i].active = false
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
	rand.Seed(time.Now().UTC().UnixNano()) //seeding the random number
	noaliens, _ := strconv.Atoi(alienNum)
	fmt.Println("city count:", iterator)
	if noaliens < iterator {
		GenerateAlienOnMap(noaliens)       //generating the aliens
		MoveTillEnd()                      //running simulation
		fmt.Println("Remaining cities !!") //printing the remaining cities in the input format
		for k, v := range cities {
			if !v.destroyed {
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