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
				mapFile := fmt.Sprintf("%s.map", c.Args().Get(0))
				if _, err := os.Stat(mapFile); err == nil {
					fmt.Printf("This map file %s already exists, try another file name...\n", mapFile)
					return
				}

				f, _ := os.OpenFile(mapFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				defer f.Close()
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
				// Input the amount of aliens you want to use
				alienNum, err := strconv.ParseInt(c.Args().Get(0), 10, 16)
				if err != nil {
					fmt.Printf("'%v' is not a valid integer", c.Args().Get(0))
					fmt.Println(alienNum)
					return
				}
			   
				 invade.StartGame(int(alienNum))

		
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
