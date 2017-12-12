package utils

import (
	"strconv"
	"strings"
	"fmt"
)

func GetKIndexArray(indexes []int) (kIndexes string) {
	buffer := []string{}

	for _, index := range indexes {
		buffer = append(buffer, "K" + strconv.Itoa(index + 1))
	}

	return strings.Join(buffer, ",")
}

func PrintMatrix(name string, matrix [][]float32) {
	fmt.Println("=== ", name, " ===")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%10.3f", matrix[i][j])
		}

		fmt.Println()
	}
}

func PrintRelations(relations [][]float32) {
	fmt.Println("=== ", "Relations", " ===")
	for i := 0; i < len(relations); i++ {
		fmt.Println("[", relations[i][0] + 1,",", relations[i][1] + 1 ,"]")
	}
}