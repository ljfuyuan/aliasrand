# aliasrand
A Go implementation of the alias method implemented using Vose's algorithm

## Usage

```go
package main

import (
	"fmt"
	"math/rand"

	"github.com/ljfuyuan/aliasrand"
)

func main() {
	weights := []uint64{1000, 2000, 3000}
	if alias, err := aliasrand.NewWeight(weights); err != nil {
		panic(err)
	} else {
		fmt.Println(alias.Pick())
		//you will get one of the 0,1,2 in O(1) time
		fmt.Println(alias.PickWithRand(rand.New(rand.NewSource(1))))
		//and you can also pick with a private random number generator
	}

	probabilities := []float64{0.01, 0.23, 0.55, 0.21}
	if alias, err := aliasrand.NewProb(probabilities); err != nil {
		panic(err)
	} else {
		fmt.Println(alias.Pick())
		//you will get one of the 0,1,2,3 in O(1) time
	}
}```
