package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Spotlight struct {
	X int // x coordinate
	Y int // y coordinate
	T int // Total of squares
}

func main()  {
	var room [][]int

	// Read matrix file
	f, err := os.Open("in.txt")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error to try read in.txt [%v]", err))
		return
	}
	defer f.Close()

	line := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		var mIntArray []int
		mStringArray := strings.Split(s.Text(), ",")
		for k, s := range mStringArray {
			if i, err := strconv.Atoi(s); err != nil {
				fmt.Println(fmt.Sprintf("El valor [%v,%v] No es un entero", line, k))
				return
			} else {
				mIntArray = append(mIntArray, i)
			}
		}
		room = append(room, mIntArray)
		line++
	}

	fmt.Println("Matriz de entrada:")
	printMatrix(room)

	// Calculate spots
	fmt.Println("\nMatriz de Salida:")
	printMatrix(SetSpotlights(room, 0, len(room) * (len(room[0]))))
}

func NewSpot(x, y, t int) Spotlight {
	return Spotlight{
		X: x,
		Y: y,
		T: t,
	}
}

func BlankSpots (s []Spotlight) bool {
	blank := true
	
	for _, v := range s {
		if v.T != 0 { blank = false }
	}
	
	return blank
}

func GetTopOfSpots(s []Spotlight) Spotlight {
	sort.Slice(s, func(i, j int) bool {
		return s[i].T > s[j].T
	})

	return s[0]
}

func SetSpotlights(room [][]int, currentI int, maxI int) [][]int {
	if currentI >= maxI {
		return room
	}

	spots := CalculateSpotlights(room)
	if BlankSpots(spots) {
		return room
	} else {
		s := GetTopOfSpots(spots)
		room[s.X][s.Y] = 2
		toRight(s.X, s.Y, room, true)
		toLeft(s.X, s.Y, room, true)
		toTop(s.X, s.Y, room, true)
		toBottom(s.X, s.Y, room, true)
		return SetSpotlights(room, currentI + 1, maxI)
	}

	return room
}

func CalculateSpotlights(room [][]int) []Spotlight {
	var spots []Spotlight

	for i, row := range room {
		for j, v := range row {
			switch v {
			case 0:
				spots = append(spots, NewSpot(i, j, CheckTotalOfIlluminatedSquares(i, j, room)))
			}
		}
	}

	return spots
}

func CheckTotalOfIlluminatedSquares(x, y int, room [][]int) int {
	return toRight(x, y, room, false) +
		toLeft(x, y, room, false) +
		toTop(x, y, room, false) +
		toBottom(x, y, room, false)
}

func toRight(x, y int, room [][]int, illuminate bool) (t int) {
	for i := y; i < len(room[x]); i++ {
		switch room[x][i] {
		case 0:
			t++
			if illuminate { room[x][i] = 3 }
		case 1:
			return t
		}
	}

	return t
}

func toLeft(x, y int, room [][]int, illuminate bool) (t int) {
	for i := y; i >= 0; i-- {
		switch room[x][i] {
		case 0:
			t++
			if illuminate { room[x][i] = 3 }
		case 1:
			return t
		}
	}

	return t
}

func toBottom(x, y int, room [][]int, illuminate bool) (t int) {
	for i := x; i < len(room); i++ {
		switch room[i][y] {
		case 0:
			t++
			if illuminate { room[i][y] = 3 }
		case 1:
			return t
		}
	}

	return t
}

func toTop(x, y int, room [][]int, illuminate bool) (t int) {
	for i := x; i >= 0; i-- {
		switch room[i][y] {
		case 0:
			t++
			if illuminate { room[i][y] = 3 }
		case 1:
			return t
		}
	}

	return t
}

func printMatrix(m [][]int) {
	for _, row := range m {
		var rowData string

		for _, v := range row {
			if v == 3 { v = 0 }
			rowData += fmt.Sprintf("| %v ", v)
		}

		fmt.Println(fmt.Sprintf("%v|", rowData))
	}
}