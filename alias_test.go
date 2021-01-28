package aliasrand

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const bound = 0.001
const count = 10000000

type params struct {
	weights  []uint64
	expected []float64
}

func build() (*Alias, params) {
	var w int64
	var p params

	n := rand.Intn(8) + 3
	p.weights = make([]uint64, 0, n)
	p.expected = make([]float64, 0, n)

	for j := 0; j < n; j++ {
		weight := rand.Int63n(500) + 1
		w += weight
		p.weights = append(p.weights, uint64(weight))
	}

	for _, weight := range p.weights {
		p.expected = append(p.expected, float64(weight)/float64(w))
	}

	a, err := NewWeight(p.weights)
	if err != nil {
		panic(err)
	}

	return a, p
}

func TestAlias(t *testing.T) {

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for i := 0; i < 10; i++ {
		a, p := build()
		n := len(a.table)
		m := make([]int, n)
		for k := 0; k < count; k++ {
			m[a.Pick()]++
		}

		for k := 0; k < n; k++ {
			current := float64(m[k]) / count
			if math.Abs(current-p.expected[k]) > bound {
				t.Fatal("Distribution did not match, seed", seed, "- got ", current, "expected", p.expected[k], "==>", i, p)
			}
		}
	}
}

func BenchmarkAlias(b *testing.B) {
	b.StopTimer()
	a, _ := build()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		a.Pick()
	}
}
