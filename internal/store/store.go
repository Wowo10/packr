package store

import (
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var pkgInts = []int{}

func ImportPacks(packsStr string) {
	pkgInts = []int{}

	if packsStr == "" {
		pkgInts = append(pkgInts, 5000, 2000, 1000, 500, 250)
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

func getOptimalPacks(packSizes []int, orderAmount int) map[int]int {
	return make(map[int]int)
}
