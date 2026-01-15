package main

import "fmt"

func main() {
	fmt.Println("Soal Nomor 6")
	fmt.Println()
	var matrixTheSixth [5][5]int
	result := matrixNumberSix(matrixTheSixth)
	displayMatrixFiveSize(result)

	fmt.Println()

	fmt.Println("Soal Nomor 7")
	fmt.Println()
	var matrixTheSeventh [5][5]int
	secondResult := matrixNumberSeven(matrixTheSeventh)
	displayMatrixFiveSize(secondResult)

	fmt.Println()

	fmt.Println("Soal Nomor 8")
	fmt.Println()
	var matrixTheEight [7][7]int
	thirdResult := matrixNumberEight(matrixTheEight)
	displayMatrixSevenSize(thirdResult)

	fmt.Println()

	fmt.Println("Soal Nomor 9")
	fmt.Println()
	var matrixTheNine [7][7]int
	fourthResult := matrixNumberNine(matrixTheNine)
	displayMatrixSevenSize(fourthResult)

	fmt.Println()

	fmt.Println("Soal Nomor 10")
	fmt.Println()
	matrixNumberTen()
}

func matrixNumberSix(matrix [5][5]int) [5][5]int {
	diagonalCounter := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				matrix[i][j] = diagonalCounter
				diagonalCounter++
			} else if j > i {
				matrix[i][j] = 10
			} else {
				matrix[i][j] = 20
			}
		}
	}
	return matrix
}

func matrixNumberSeven(matrix [5][5]int) [5][5]int {
	diagonalCounter := len(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				matrix[i][j] = diagonalCounter
				diagonalCounter--
			} else if j > i {
				matrix[i][j] = 20
			} else {
				matrix[i][j] = 10
			}
		}
	}
	return matrix
}

func matrixNumberEight(matrix [7][7]int) [7][7]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				matrix[i][j] = j
			} else if j == 0 {
				matrix[i][j] = i
			} else if i == len(matrix)-1 || j == len(matrix)-1 {
				matrix[i][j] = i + j
			}
		}
	}
	return matrix
}

func matrixNumberNine(matrix [7][7]int) [7][7]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = i + j
		}
	}

	return matrix
}

func matrixNumberTen() {
	matrixTenthSize := [8][10]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
	}

	keyAnswer := [10]string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}

	for i := 0; i < len(matrixTenthSize); i++ {
		counter := 0
		for j := 0; j < len(matrixTenthSize[i]); j++ {
			if matrixTenthSize[i][j] == keyAnswer[j] {
				counter++
			}
		}
		fmt.Printf("Jawaban Siswa %d yang benar : %d\n", i, counter)
	}
}

func displayMatrixFiveSize(matrix [5][5]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%5d", matrix[i][j])
		}
		fmt.Println()
	}
}

func displayMatrixSevenSize(matrix [7][7]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%5d", matrix[i][j])
		}
		fmt.Println()
	}
}
