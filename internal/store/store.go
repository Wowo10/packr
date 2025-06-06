package store

import (
	"log"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var pkgInts = []int{}

const stateFile = `./data/save`

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

func LoadFile() {
	log.Println("Loading file...")
	if _, err := os.Stat(stateFile); err == nil {
		data, err := os.ReadFile(stateFile)
		if err != nil {
			log.Fatalf("error reading file: %v", err)
		}

		if len(data) > 0 {
			log.Println("Correct data found...")
			ImportPacks(string(data))
		}
	}
}

func SavePacks() {
	log.Println("Saving file...")
	strNums := make([]string, len(pkgInts))
	for i, num := range pkgInts {
		strNums[i] = strconv.Itoa(num)
	}

	err := os.MkdirAll(filepath.Dir(stateFile), 0755)
	if err != nil {
		log.Println("error creating directories:", err)
		return
	}

	err = os.WriteFile(stateFile, []byte(strings.Join(strNums, ",")), 0644)
	if err != nil {
		log.Println("error writing file:", err)
	}
	log.Println("File saved...")
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
			SavePacks()
			return
		}
	}

	pkgInts = append(pkgInts, pack)
	SavePacks()
}

func RemovePack(pack int) {
	for i, v := range pkgInts {
		if v == pack {
			pkgInts = slices.Delete(pkgInts, i, i+1)
			SavePacks()
			return
		}
	}
}

func Solve(amount int) map[int]int {
	return getOptimalPacks(pkgInts, amount)
}

type dpState struct {
	sumItems  int
	packCount int
	packMap   map[int]int
}

func getOptimalPacks(packSizes []int, orderAmount int) map[int]int {
	// assume descending order
	maxPack := packSizes[0]
	limit := orderAmount + maxPack

	dp := make(map[int]*dpState, limit+1)
	dp[0] = &dpState{0, 0, map[int]int{}}

	queue := []int{0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, pack := range packSizes {
			newTotal := current + pack
			if newTotal > limit {
				continue
			}

			// Copy previous state, increment pack count and packMap[pack]
			newPackMap := make(map[int]int)
			maps.Copy(newPackMap, dp[current].packMap)
			newPackMap[pack]++

			newPackCount := dp[current].packCount + 1

			// Save to dp table if new state is better
			if dp[newTotal] == nil ||
				newTotal < dp[newTotal].sumItems ||
				(newTotal == dp[newTotal].sumItems && newPackCount < dp[newTotal].packCount) {

				dp[newTotal] = &dpState{
					sumItems:  newTotal,
					packCount: newPackCount,
					packMap:   newPackMap,
				}

				queue = append(queue, newTotal)
			}
		}
	}

	// Return the best valid option => orderQty
	for i := orderAmount; i <= limit; i++ {
		if dp[i] != nil {
			return dp[i].packMap
		}
	}

	// Can't happen
	return nil

}
