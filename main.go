package main

import (
	"log"

	"github.com/urfave/cli"

	"os"

	"fmt"

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
			Name:  "start",
			Usage: "Start the  Invasion!!",
			Action: func(c *cli.Context) {
				//Create a map file on the start command
				f, err := os.Open(c.Args().Get(0))
				if err != nil {
					fmt.Printf("This file probably doesnt exist, Try using the 'testMapFile.map'...\n")

				}
				defer f.Close()

				//Generate cities and map from map file
				invade.GenerateCityMap(c.Args().Get(0))

				//simulates the invasion by passing in number of aliens
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
