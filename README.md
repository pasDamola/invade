# Invade

## Introduction
This is a simulation of an alien invasion where aliens start out randomly distributed in different cities, taking different routes. If two aliens arrive at a city, the aliens fight and that city is destroyed.
This project was completed with the https://github.com/urfave/cli, simple package for completing CLI applications in Golang.
Tests were also written to capture some edge cases.


This project runs under the following assumptions:

1) An alien becomes inactive(destroyed) when it has reached the end of the steps or two are found in the same city.
2) The number of aliens should be less than the number of cities.
3) The system running the program has go installed in them.


## Run the Program
To run the program,

1) git clone https://github.com/pasDamola/invade
2) cd invade
3) go run main.go start files/testMapfile.map numAliens

## Tests
To run the tests, assuming you are already in the project's root directory:
1) cd invade
2) go test -v


