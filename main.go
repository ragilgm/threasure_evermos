package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main()  {

	// map data
	var mapData =
			"########\n" +
			"#......#\n" +
			"#.###..#\n" +
			"#...#.##\n" +
			"#X#....#\n" +
			"########\n"

	// define probable treasure
	var probableTreasure = defineTreasure(mapData)
	

	// how to play
	var howToPlay = "- Press 'A' for move Up\n" +
		"- Press 'B' for move Right\n" +
		"- Press 'C' for move Bottom\n" +
		"- Press 'D' for move Left "


	// begin
	fmt.Println("Treasure Hunt")
	fmt.Println("---------------------")
	fmt.Println("Press Y for start")
	fmt.Println("Press X for end")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")

		// read string from user input
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		// if user press X then program will close
		if strings.Compare("X", strings.ToUpper(text)) == 0 {
			fmt.Println("I hope you come back again ..")
			break
		}else if strings.Compare("Y", strings.ToUpper(text)) == 0{
			for {
				fmt.Println(mapData)
				fmt.Println(howToPlay)

				fmt.Print("-> ")
				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)

				if strings.Compare("X", strings.ToUpper(text)) == 0 {
					fmt.Println("I hope you come back again ..")
					break
				}
				mapData = controlData(mapData,strings.ToUpper(text),probableTreasure)

				for i := 0; i < len(mapData); i++ {
					if mapData[i]== '$' {
						fmt.Println("Congratulation The treasure was found. ")
						return
					}
				}
			}
		}

	}

}

// control user
func controlData(mapData string,move string, treasure int) string{

 var currentPosition = strings.Index(mapData,"X")
	var nextPosition = 0;
	switch move {
	case "A":
		nextPosition= currentPosition -9
		break
	case "B":
		 nextPosition = currentPosition +1
		break
	case "C":
		 nextPosition = currentPosition +9
	case "D":
		 nextPosition = currentPosition -1
	default:
		break
	}
	if strings.Compare("#", string(mapData[nextPosition])) == 0{
	}else if nextPosition==treasure {
		mapData = strings.Replace(mapData,"X",".",currentPosition)
		mapData = MovePerson(mapData,'$',nextPosition)
	}else{
		mapData = strings.Replace(mapData,"X",".",currentPosition)
		mapData = MovePerson(mapData,'X',nextPosition)
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
func defineTreasure (data string) int {
	var min = 1
	var max = 54
	probable:= rand.Intn(54)
	var obstacle = []int{}
	for i := 0; i <len(data); i++ {
		if string(data[i])=="."{
			obstacle = append(obstacle, i)
		}
	}
	for  {
		for _,value :=range obstacle{
			if value==probable {
				return probable
			}
		}
		probable = rand.Intn(max-min)
	}
}