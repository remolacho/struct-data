package functions

import (
	"fmt"
	tm "github.com/buger/goterm"
	"strconv"
	"strings"
	"time"
)

var matrix [3][3]string
var charPlayer1 string
var charPlayer2 string
var playerName1 string
var playerName2 string
var hasValues int
var top int

func ThreeInLine() {
	clearTerm()
	settingPlayers()
	play()
	assignIndex()
	printMatrix()
}

func settingPlayers() {
	fmt.Printf("Elegir un nombre para el jugador 1 \n")
	fmt.Scanf("%s\n", &playerName1)

	fmt.Printf("Elegir un nombre para el jugador 2 \n")
	fmt.Scanf("%s\n", &playerName2)

	char1 := "X"
	char2 := "O"

	for charPlayer1 != char1 && charPlayer1 != char2 {
		fmt.Printf("%s Elige entre %s y %s \n", playerName1, char1, char2)
		fmt.Scanf("%s\n", &charPlayer1)
		charPlayer1 = strings.ToUpper(charPlayer1)
	}

	if charPlayer1 == char1 {
		charPlayer2 = char2
	} else {
		charPlayer2 = char1
	}

	clearTerm()

	fmt.Printf("---- Al %s se asigna %s \n", playerName1, charPlayer1)
	fmt.Printf("---- Al %s se asigna %s \n", playerName2, charPlayer2)
}

func play() {
	var acum int
	top = 3

	for i := 0; i < top; i++ {
		for j := 0; j < top; j++ {
			acum++
			hasValues += acum
			matrix[i][j] = strconv.Itoa(acum)
		}
	}
}

func assignIndex() {
	if playPlayer(charPlayer1, playerName1) || playPlayer(charPlayer2, playerName2) {
		return
	}

	assignIndex()
}

func playPlayer(player string, playerName string) bool {
	var index string
	var indexValid bool

	for !indexValid {
		fmt.Printf("%s = %s elige un numero disponible en el tablero \n", playerName, player)
		printMatrix()
		fmt.Scanf("%s\n", &index)
		indexValid = findValue(index, player)
		clearTerm()
	}

	if win(player) {
		fmt.Printf("En hora buena %s = %s haz ganado!!!!! \n", playerName, player)
		return true
	}

	if tie() {
		fmt.Println("Termina en empate!!!!!")
		return true
	}

	return false
}

func findValue(index string, player string) bool {
	for i := 0; i < top; i++ {
		for j := 0; j < top; j++ {
			if matrix[i][j] == index {
				value, _ := strconv.Atoi(matrix[i][j])
				hasValues -= value
				matrix[i][j] = player
				return true
			}
		}
	}
	return false
}

func win(charPlayer string) bool {
	if findRow(charPlayer) || findColumn(charPlayer) || findLeft(charPlayer) || findRight(charPlayer) {
		return true
	} else {
		return false
	}
}

func findRow(charPlayer string) bool {
	for row := 0; row < top; row++ {
		inLine := true

		for col := 0; col < top && inLine; col++ {
			if charPlayer != matrix[row][col] {
				inLine = false
			}
		}

		if inLine {
			return true
		}
	}

	return false
}

func findColumn(charPlayer string) bool {
	for col := 0; col < top; col++ {
		inLine := true

		for row := 0; row < top && inLine; row++ {
			if charPlayer != matrix[row][col] {
				inLine = false
			}
		}

		if inLine {
			return true
		}
	}

	return false
}

func findLeft(charPlayer string) bool {
	for i := 0; i < top; i++ {
		if charPlayer != matrix[i][i] {
			return false
		}
	}

	return true
}

func findRight(charPlayer string) bool {
	col := top - 1
	for row := 0; row < top; row++ {
		if charPlayer != matrix[row][col] {
			return false
		}
		col--
	}

	return true
}

func tie() bool {
	if hasValues == 0 {
		return true
	} else {
		return false
	}
}

func printMatrix() {
	fmt.Println("-------------")
	for i := 0; i < top; i++ {
		for j := 0; j < top; j++ {
			fmt.Printf("| %s ", matrix[i][j])
		}
		fmt.Println("|")
		fmt.Println("-------------")
	}
}

func clearTerm() {
	tm.Clear() // Clear current screen
	tm.MoveCursor(1, 1)
	tm.Flush() // Call it every time at the end of rendering
	time.Sleep(time.Second / 2)
}
