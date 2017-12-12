package main

import (
	"electre/electre"
	"electre/utils"

	"flag"
	"log"
	"fmt"
)

var fileName = flag.String("file", "data.json", "File with formatted data (json format)")
var algorithm = flag.Int("electre", electreOne, "Electre algorithm [1 or 2]")

func init() {
	flag.Parse()
}

const (
	electreOne = 1
	electreTwo = 2
)

func main() {
	fmt.Println("FILE:    ", *fileName)
	fmt.Println("Electre: ", *algorithm)

	variant, err := readFile(*fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	var f, d [][]float32

	switch *algorithm {
	case electreOne:
		f, d = electre.GetIndices(variant, electre.Electre1)
	case electreTwo:
		f, d = electre.GetIndices(variant, electre.Electre2)
	default:
		log.Fatal("unknown electre chosen")
	}

	utils.PrintMatrix("f matrix", f)
	utils.PrintMatrix("d marix", d)

	requiredCondition := electre.GetRequiredConditionMatrix(f, variant.C)
	utils.PrintMatrix("Requiered conditions", requiredCondition)

	sufficientConditions := electre.GetSufficientConditionMatrix(d, variant.D)
	utils.PrintMatrix("Sufficient conditions", sufficientConditions)

	combinedConditions := electre.CombineConditions(requiredCondition, sufficientConditions)
	utils.PrintMatrix("Combined conditions", combinedConditions)

	relations := electre.GetRelation(combinedConditions)
	utils.PrintRelations(relations)
}
