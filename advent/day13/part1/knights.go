package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/skarlso/goutils/arrayutils"
)

var seatingCombinations = make([][]int, 0)
var table = make(map[int][]map[int]int)
var keys = make([]int, 0)
var nameMapping = make(map[string]int)

//Person a person
type Person struct {
	// neighbour *Person
	name string
	like int
}

func main() {
	CalculatePerfectSeating()
}

//CalculatePerfectSeating returns the perfect seating order based on Love/Hate relations
func CalculatePerfectSeating() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		trimmedNeighbour := strings.Trim(split[10], ".")
		like, _ := strconv.Atoi(split[3]) //If lose -> * -1
		if _, ok := nameMapping[split[0]]; !ok {
			nameMapping[split[0]] = id
			id++
		}
		if _, ok := nameMapping[trimmedNeighbour]; !ok {
			nameMapping[trimmedNeighbour] = id
			id++
		}
		if split[2] == "lose" {
			like *= -1
		}
		table[nameMapping[split[0]]] = append(table[nameMapping[split[0]]], map[int]int{nameMapping[trimmedNeighbour]: like})
		if !arrayutils.ContainsInt(keys, nameMapping[split[0]]) {
			keys = append(keys, nameMapping[split[0]])
		}
	}
	generatePermutation(keys, len(keys))
	// fmt.Println("Best seating efficiency:", calculateSeatingEfficiancy())
}

func generatePermutation(s []int, n int) {
	if n == 1 {
		news := make([]int, len(s))
		copy(news, s)
		seatingCombinations = append(seatingCombinations, news)
	}
	for i := 0; i < n; i++ {
		s[i], s[n-1] = s[n-1], s[i]
		generatePermutation(s, n-1)
		s[i], s[n-1] = s[n-1], s[i]
	}
}

func calculateSeatingEfficiancy() int {
	bestSeating := math.MinInt64
	for _, v := range seatingCombinations {
		calculatedOrder := 0

		for i := range v {
			left := (i - 1) % len(v)
			//This is to work around the fact that in Go
			//modulo of a negative number will not return a positive number.
			//So -1 % 4 will not return 3 but -1. In that case we add length.
			if left < 0 {
				left += len(v)
			}
			right := (i + 1) % len(v)
			calculatedOrder += getLikeForTargetConnect(v[i], v[left]) + getLikeForTargetConnect(v[i], v[right])
		}
		if calculatedOrder > bestSeating {
			bestSeating = calculatedOrder
		}
	}

	return bestSeating
}

func getLikeForTargetConnect(name int, neighbour int) int {
	neighbours := table[name]
	for _, t := range neighbours {
		if v, ok := t[neighbour]; ok {
			return v
		}
	}
	return 0
}
