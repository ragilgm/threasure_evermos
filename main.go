package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {

	// define reader stdin
	reader := bufio.NewReader(os.Stdin)

	// map data
	var mapData = "########\n" +
		"#......#\n" +
		"#.###..#\n" +
		"#...#.##\n" +
		"#X#....#\n" +
		"########\n"

	// define probable treasure

	// how to play
	var howToPlay = "- Press 'A' for move Up\n" +
		"- Press 'B' for move Right\n" +
		"- Press 'C' for move Bottom\n" +
		"- Press 'D' for move Left "

	var treasureFound = false
	var tryAgain = false

	for {

		// begin
		fmt.Println("Treasure Hunt")
		fmt.Println("---------------------")
		fmt.Println("Press Y for start")
		fmt.Println("Press X for end")

		fmt.Print("-> ")
		var text, _, err = reader.ReadLine()

		if err != nil {
			fmt.Println(err)
		}

		// if user press X then program will close
		if strings.Compare("X", strings.ToUpper(string(text))) == 0 {
			fmt.Println("I hope you come back again ..")
			break
		} else if strings.Compare("Y", strings.ToUpper(string(text))) == 0 {
			var mapGame = mapData

			for {
				var probableTreasure = defineTreasure(mapGame)
				for {
					fmt.Println(mapGame)
					fmt.Println(howToPlay)

					fmt.Print("-> ")
					text, _, err = reader.ReadLine()

					if err != nil {
						fmt.Println(err)
					}

					if strings.Compare("X", strings.ToUpper(string(text))) == 0 {
						fmt.Println("I hope you come back again ..")
						break
					}
					mapGame = controlData(mapGame, strings.ToUpper(string(text)), probableTreasure)

					for i := 0; i < len(mapGame); i++ {
						if mapGame[i] == '$' {

							fmt.Println(mapGame)

							fmt.Println("Congratulation The treasure was found. ")
							treasureFound = true
							break
						}
					}

					if treasureFound {
						break
					}

				}

				for {
					fmt.Println("Are you want to try again ?")
					fmt.Println("Press Y for Yes")
					fmt.Println("Press X for Close")
					text, _, err = reader.ReadLine()
					if strings.Compare("Y", strings.ToUpper(string(text))) == 0 {
						treasureFound = false
						tryAgain = true
						break
					}
					if strings.Compare("X", strings.ToUpper(string(text))) == 0 {
						tryAgain = false
						break
					}
				}

				if tryAgain {
					mapGame = mapData
				} else {
					fmt.Println("I hope you come back again ..")
					return
				}

			}

		}

	}

}

// control user
func controlData(mapData string, move string, treasure int) string {

	var currentPosition = strings.Index(mapData, "X")
	var nextPosition = 0
	switch move {
	case "A":
		nextPosition = currentPosition - 9
		break
	case "B":
		nextPosition = currentPosition + 1
		break
	case "C":
		nextPosition = currentPosition + 9
	case "D":
		nextPosition = currentPosition - 1
	default:
		break
	}
	if strings.Compare("#", string(mapData[nextPosition])) == 0 {
		return mapData
	} else if nextPosition == treasure {
		mapData = strings.Replace(mapData, "X", ".", currentPosition)
		mapData = MovePerson(mapData, '$', nextPosition)
	} else {
		mapData = strings.Replace(mapData, "X", ".", currentPosition)
		mapData = MovePerson(mapData, 'X', nextPosition)
	}
	return mapData

}

// move person
func MovePerson(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}

// define treasure
func defineTreasure(data string) int {
	var min = 1
	var max = 54
	probable := rand.Intn(54)
	var obstacle = []int{}
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "." {
			obstacle = append(obstacle, i)
		}
	}
	for {
		for _, value := range obstacle {
			if value == probable {
				return probable
			}
		}
		probable = rand.Intn(max - min)
	}
}
