package store

import (
	"log"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var pkgInts = []int{}

func ImportPacks(packsStr string) {
	pkgInts = []int{}

	if packsStr == "" {
		pkgInts = append(pkgInts, 250, 500, 1000, 2000, 5000)
		return
	}

	// Parse Packs
	packsStr = strings.ReplaceAll(packsStr, " ", "")
	initPacksSlice := strings.SplitSeq(packsStr, ",")

	for v := range initPacksSlice {
		intV, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("warning: initPacks contains non-integer value, %v", v)
			continue
		}
		pkgInts = append(pkgInts, intV)
	}

	// Sort the pkgInts slice in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(pkgInts)))
}

func GetPacks() []int {
	return pkgInts
}

func AddPack(pack int) {
	for i, v := range pkgInts {
		if v == pack {
			return
		}
		if v < pack {
			pkgInts = append(pkgInts[:i], append([]int{pack}, pkgInts[i:]...)...)
			return
		}
	}

	pkgInts = append(pkgInts, pack)
}

func RemovePack(pack int) {
	for i, v := range pkgInts {
		if v == pack {
			pkgInts = slices.Delete(pkgInts, i, i+1)
			return
		}
	}
}

func Solve(amount int) map[int]int {
	return getOptimalPacks(pkgInts, amount)
}

// TODO: Need another approach, doesn`t work for the edge case`
func getOptimalPacks(packSizes []int, orderAmount int) map[int]int {
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})

	result := make(map[int]int)

	//try greedy approach
	for _, pack := range packSizes {
		count := orderAmount / pack
		result[pack] = count
		orderAmount %= pack
	}

	//find overfill if needed
	if orderAmount > 0 {
		incrementMinimalOverfill(result, packSizes)
	}

	return result
}

func incrementMinimalOverfill(solution map[int]int, packSizes []int) {
	// Find the highest filled pack
	// Try incrementing the next larger pack size and zero out smaller ones
	// Choose the configuration with the smallest total sum,
	// breaking ties using the lowest number of packs

	minSum := math.MaxInt
	minNumber := math.MaxInt
	minIndex := 0

	startIndex := len(packSizes) - 1

	for i, pack := range packSizes {
		if solution[pack] > 0 {
			if i == 0 {
				startIndex = 0
			} else {
				startIndex = i - 1
			}
			break
		}
	}

	for i := startIndex; i < len(packSizes); i++ {
		sum := 0
		number := 0
		for j := range i {
			sum = sum + solution[packSizes[j]]*packSizes[j]
			number = number + solution[packSizes[j]]
		}

		nextCount := solution[packSizes[i]] + 1
		sum += nextCount * packSizes[i]
		number += nextCount

		if sum < minSum {
			minSum = sum
			minNumber = number
			minIndex = i
		} else if sum == minSum && number < minNumber {
			minNumber = number
			minIndex = i
		}
	}

	solution[packSizes[minIndex]] = solution[packSizes[minIndex]] + 1

	for i := minIndex + 1; i < len(packSizes); i++ {
		solution[packSizes[i]] = 0
	}

	return
}
