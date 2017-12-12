package main

import (
	"flag"
	"log"
	"electra/electre"
	"fmt"
)

var fileName = flag.String("file","data.json","income data")

func printMatrix(name string , matrix [][]float32) {
	fmt.Println("=== ", name, " ===")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%10.3f", matrix[i][j])
		}

		fmt.Println()
	}
}

func main() {
	variant, err := readFile(*fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	f, d := electre.GetIndices(variant, electre.Electre2)

	printMatrix("f", f)
	printMatrix("d", d)

	requiredCondition := electre.GetRequiredConditionMatrix(f, variant.C)
	sufficientConditions := electre.GetSufficientConditionMatrix(d, variant.D)

	printMatrix("requieredCondition", requiredCondition)
	printMatrix("sufficientConditions", sufficientConditions)

	combinedConditions := electre.CombineConditions(requiredCondition, sufficientConditions)

	printMatrix("combinedConditions", combinedConditions)

	relations := electre.GetRelation(combinedConditions)
	fmt.Println("=== Relations ===")
	fmt.Println(relations)
}