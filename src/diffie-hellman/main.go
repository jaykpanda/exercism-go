package main

import (
	"math"
	"math/rand"
	"time"
)

func main() {

}

func gen(n int) int {
	min := math.Pow(2, n-1)
	max := math.Pow(2, n)
	rand.Seed(time.Now().UTC().UnixNano())
	rnum := min + rand.Intn(max-min)
	check := 0
	count := 0
	for check == 0 {
		for i := 2; i < rnum; i++ {
			if (math.Mod(rnum, i)) != 0 {
				count++
			}
		}
		if count == 0 {
			isPrime = true
		} else {
			count = 0
			rnum = min + rand.Intn(max-min)
		}
	}
	return rnum
}
