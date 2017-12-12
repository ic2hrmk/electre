package electre

import (
	"electre/utils"

	"math"
	"fmt"
)

type ElectreType func(kPlus, kZero, kMinus float32, w []float32) float32

func Electre1(pPlus, pZero, pMinus float32, w []float32) float32 {
	var sum float32 = 0.0
	for i := range w {
		sum += w[i]
	}

	return (pPlus + pZero) / sum
}

func Electre2(pPlus, pZero, pMinus float32, w []float32) float32 {
	return pPlus / pMinus
}

func GetIndices(variant *Variant, electreMethod ElectreType) (f [][]float32, d [][]float32) {
	X := variant.X
	W := variant.W

	indicesNumber := len(variant.X)
	koefNumber := len(variant.X[0])

	//	Memory management
	f = make([][]float32, indicesNumber)
	d = make([][]float32, indicesNumber)

	for i := 0; i < indicesNumber; i++ {
		f[i] = make([]float32, indicesNumber)
		d[i] = make([]float32, indicesNumber)
	}

	//	Iterations
	for i := 0; i < indicesNumber; i++ {
		for j := 0; j < indicesNumber; j++ {

			if i == j {
				continue
			}

			mainRow := X[i]
			compRow := X[j]

			var kPlus []int
			var kEq []int
			var kMinus []int

			for k := 0; k < koefNumber; k++ {
				switch {
				case mainRow[k] > compRow[k]:
					kPlus = append(kPlus, k)
				case mainRow[k] < compRow[k]:
					kMinus = append(kMinus, k)
				default:
					kEq = append(kEq, k)
				}
			}

			fmt.Println("X[", i+1, ",", j+1, "]")
			fmt.Println("   K(+) = {", utils.GetKIndexArray(kPlus), "}")
			fmt.Println("   K(=) = {", utils.GetKIndexArray(kEq), "}")
			fmt.Println("   K(-) = {", utils.GetKIndexArray(kMinus), "}")

			fmt.Println("   --- ")

			pPlus := indexSum(W, kPlus)
			pEq := indexSum(W, kEq)
			pMinus := indexSum(W, kMinus)

			fmt.Println("   P(+) = ", pPlus)
			fmt.Println("   P(=) = ", pEq)
			fmt.Println("   P(-) = ", pMinus)

			f[i][j] = electreMethod(pPlus, pEq, pMinus, W)
			d[i][j] = calcNonComplianceIndices(kMinus, W, mainRow, compRow)
		}
	}

	return
}

func GetRequiredConditionMatrix(f [][]float32, c float32) (required [][]float32) {
	indicesNumber := len(f)

	//	Memory management
	required = make([][]float32, indicesNumber)
	for i := 0; i < indicesNumber; i++ {
		required[i] = make([]float32, indicesNumber)
	}

	for i := 0; i < indicesNumber; i++ {
		for j := 0; j < indicesNumber; j++ {
			if f[i][j] > c {
				required[i][j] = float32(1.0)
			}
		}
	}

	return
}

func GetSufficientConditionMatrix(d [][]float32, s float32) (sufficient [][]float32) {
	indicesNumber := len(d)

	//	Memory management
	sufficient = make([][]float32, indicesNumber)
	for i := 0; i < indicesNumber; i++ {
		sufficient[i] = make([]float32, indicesNumber)
	}

	for i := 0; i < indicesNumber; i++ {
		for j := 0; j < indicesNumber; j++ {
			if d[i][j] < s {
				sufficient[i][j] = float32(1.0)
			}
		}
	}

	return
}

func CombineConditions(requiredCondition [][]float32, sufficientConditions [][]float32) (combinedRequirements [][]float32) {
	indicesNumber := len(requiredCondition)

	//	Memory management
	combinedRequirements = make([][]float32, indicesNumber)
	for i := 0; i < indicesNumber; i++ {
		combinedRequirements[i] = make([]float32, indicesNumber)
	}

	for i := 0; i < len(requiredCondition); i++ {
		for j := 0; j < len(requiredCondition); j++ {
			if requiredCondition[i][j] == 1.0 && sufficientConditions[i][j] == 1.0 {
				combinedRequirements[i][j] = 1.0
			}
		}
	}

	return
}

func GetRelation(combinedConditions [][]float32) (relations [][]float32) {
	indicesNumber := len(combinedConditions)

	for i := 0; i < indicesNumber; i++ {
		for j := 0; j < indicesNumber; j++ {
			if combinedConditions[i][j] == 1.0 {
				relations = append(relations, []float32{float32(i), float32(j)})
			}
		}
	}

	return
}

//	Helpers

func indexSum(w []float32, k []int) (sum float32) {
	sum = 0

	for _, index := range k {
		sum += w[index]
	}

	return
}

func calcNonComplianceIndices(kMinus []int, w []float32, x1 []float32, x2 []float32) (d float32) {
	if len(kMinus) == 0 {
		d = 0
	} else {
		var maxByKMinus float32
		var maxByAll float32

		//	Indexes of smaller values
		for _, i := range kMinus {
			temp := w[i] * float32(math.Abs(float64(x1[i]-x2[i])))

			if temp > maxByKMinus {
				maxByKMinus = temp
			}
		}

		//	Indexes of all values
		for i, _ := range w {
			temp := w[i] * float32(math.Abs(float64(x1[i]-x2[i])))

			if temp > maxByAll {
				maxByAll = temp
			}
		}

		d = maxByKMinus / maxByAll
	}

	return
}