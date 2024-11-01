package main

import "fmt"

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	dpa := make([]int64, len(energyDrinkA))
	dpb := make([]int64, len(energyDrinkB))

	dpa[0] = int64(energyDrinkA[0])
	dpb[0] = int64(energyDrinkB[0])

	for i := 1; i < len(energyDrinkB); i++ {
		dpa[i] = max(dpa[i-1]+int64(energyDrinkA[i]), dpb[i-1])
		dpb[i] = max(dpb[i-1]+int64(energyDrinkB[i]), dpa[i-1])
	}

	return max(dpa[len(energyDrinkA)-1], dpb[len(energyDrinkB)-1])
}

func main() {
	fmt.Println(maxEnergyBoost([]int{1, 3, 1}, []int{3, 1, 1}))
	fmt.Println(maxEnergyBoost([]int{4, 1, 1}, []int{1, 1, 3}))
}
