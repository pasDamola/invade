package main

import (
	"log"

	"github.com/urfave/cli"

	"os"

	"fmt"

	"strconv"

	"github.com/pasDamola/invade/invade"

	
)

var app = cli.NewApp()

func info() {
	app.Name = "Alien Invasion"
	app.Usage = "Simulating the 'Alien Invasion Saga' on Planet Earth"
	app.Author = "Oyindamola"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:  "map",
			Usage: "Create your map file",
			Action: func(c *cli.Context) {
				//Create a map file with the map command, passing the name of file as an argument
				mapFile := fmt.Sprintf("%s.map", c.Args().Get(0))
				if _, err := os.Stat(mapFile); err == nil {
					fmt.Printf("This map file %s already exists, try using another file name...\n", mapFile)
					return
				}

			},
		},
		{
			Name:  "aliens",
			Usage: "Create ",
			Action: func(c *cli.Context) {

			},
		},
		{
			Name:  "start",
			Usage: "Start the  Invasion!!",
			Action: func(c *cli.Context) {
				//Create a map file on the start command
				f, err := os.Open(c.Args().Get(0))
				if err != nil {
					fmt.Printf("This file probably doesnt exist, Try creating a new map file with the 'map' command...\n")

				}
				defer f.Close()

				//parse the number of aliens for this invasion
				alienNum, err := strconv.ParseInt(c.Args().Get(1), 10, 16)
				if err != nil {
					fmt.Printf("'%v' is not a valid integer", c.Args().Get(1))
					fmt.Println(alienNum)
					return
				}
				//continuePlay := true;

				invade.GenerateCityMap(c.Args().Get(0)) 


				invade.Run(c.Args().Get(1))

			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
