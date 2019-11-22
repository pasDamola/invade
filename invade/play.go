package invade

import (
	"strings"

	"os"

	"bufio"

	"fmt"

	"time"

	"math/rand"
)

var alienMoves = 10000

const alienNameLength = 3
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// City represents a city with roads going out of it
type City struct {
	Name  string
	Roads []*Road
}

// Road represents a connection between cities
type Road struct {
	Direction string
	City      string
}

// Alien represents an alien
type Alien struct {
	Name  string
	City  string
	Turns int
}

// Map represents the game map. It contains an array of cities and some aliens
type Map struct {
	Cities []*City
	Aliens []*Alien

	rand *rand.Rand
}

// Play runs the game
func (m *Map) Play() {
	for i := 0; i < alienMoves; i++ {
		m.MoveAliens()
		m.DestroyCitiesAndAliens(i)
		// If there are no remaining aliens then end the game
		if len(m.Aliens) == 0 {
			break
		}
	}
}

// NewMap creates a new map initialized with a source of randomness
func NewMap() *Map {
	return &Map{rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// NewMapFromFile returns a new instance of map based on the file
func NewMapFromFile(file string) (*Map, error) {
	// Instantiate the map
	m := NewMap()

	// Open the file and return any errors
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("This file probably doesnt exist, Try creating a new map file with the 'map' command...\n")
	}
	defer f.Close()

	// Scan the file by lines
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		city := &City{Roads: make([]*Road, 0)}
		// Split the lines on spaces
		l := strings.Split(strings.TrimSpace(scan.Text()), " ")
		// If the line is empty or too long ignore it
		if len(l) > 0 && len(l) < 6 && !strings.Contains(l[0], "=") {
			city.Name = l[0]
			l = l[1:]
			for _, rd := range l {
				// Split the roads on "="
				kv := strings.Split(rd, "=")
				city.Roads = append(city.Roads, &Road{Direction: kv[0], City: kv[1]})
			}
			m.Cities = append(m.Cities, city)
			continue
		}
	}
	fmt.Printf("%v", m)
	return m, nil
}

// NewAliens creates n new Alien objects with random names
func (m *Map) newAliens(n int) []*Alien {
	out := make([]*Alien, 0)
	names := m.randStrings(n, alienNameLength)
	for _, name := range names {
		out = append(out, &Alien{Name: name, Turns: 0})
	}
	return out
}

// randStrings returns a list of n random strings of length l
func (m *Map) randStrings(n, l int) []string {
	cityNames := make([]string, 0, n)
	for i := 0; i < n; i++ {
		cityNames = append(cityNames, m.randStringOfLength(l))
	}
	return cityNames
}

// Generate a random string of a particular length
func (m *Map) randStringOfLength(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, m.rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = m.rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// MoveAliens moves all of the aliens
func (m *Map) MoveAliens() {
	for _, a := range m.Aliens {
		m.moveAlien(a)
	}
}

// DestroyCitiesAndAliens reconciles the map at the end of the turn
// by removing the aliens, cities, and roads that have been destroyed
func (m *Map) DestroyCitiesAndAliens(turn int) {
	// First create a mapping of city -> aliens
	out := make(map[string][]string, 0)
	for _, a := range m.Aliens {
		if _, pres := out[a.City]; pres {
			out[a.City] = append(out[a.City], a.Name)
			continue
		}
		out[a.City] = []string{a.Name}
	}

	// Then loop over and destroy all the cities > 2 aliens
	for k, v := range out {
		if len(v) > 1 {
			msg := fmt.Sprintf("%s has been destroyed by ", k)
			for _, a := range v {
				msg += fmt.Sprintf("alien %s and ", a)
			}
			msg = strings.TrimSuffix(msg, " and ")
			msg += fmt.Sprintf("! (turn %d)", turn)
			fmt.Println(msg)
			
		}
	}
}


// Return a new city for the alien, if there are no roads return an empty string
func (m *Map) moveAlien(a *Alien) {
	a.Turns++
	for _, c := range m.Cities {
		if c.Name == a.City && len(c.Roads) > 0 {
			a.City = c.RandRoad().City
			return
		}
	}
}

// RandRoad returns a random road from a city
func (c *City) RandRoad() *Road {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return c.Roads[r.Intn(len(c.Roads))]
}