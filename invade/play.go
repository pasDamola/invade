package invade

import (

	"bufio"

	"fmt"

	"os"
)

func BuildWorld(path string) ([]string, error) {
	file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
	}

	fmt.Println(len(lines))

	return lines, scanner.Err()
}
