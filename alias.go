package aliasrand

import "math/rand"

type ele struct {
	prob float64
	idx  int
}

type Alias struct {
	table []ele
}

//Pick pick a random value sampled from the underlying distribution.
func (a *Alias) Pick() int {
	return a.PickWithRand(nil)
}

//PickWithRand pick a random value sampled from the underlying distribution with the specific random generator
func (a *Alias) PickWithRand(r *rand.Rand) int {

	var rnd float64

	if r == nil {
		rnd = rand.Float64() * float64(len(a.table))
	} else {
		rnd = r.Float64() * float64(len(a.table))
	}

	cloumn := int(rnd)

	if rnd-float64(cloumn) < a.table[cloumn].prob {
		return cloumn
	}

	return a.table[cloumn].idx
}
