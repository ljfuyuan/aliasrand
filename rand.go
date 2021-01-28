//Package aliasrand a faster random method implemented using Vose's algorithm

//For a complete writeup on the alias method, including the intuition and
//important proofs, please see the article "Darts, Dice, and Coins: Smpling
//from a Discrete Distribution" at

//http://www.keithschwarz.com/darts-dice-coins/
package aliasrand

import (
	"errors"
)

//NewProb constructs a new AliasMethod from the given probabilities
func NewProb(probabilities []float64) (*Alias, error) {

	eles, err := setup(probabilities)

	if err != nil {
		return nil, err
	}

	return &Alias{table: eles}, nil
}

//NewWeight constructs a new AliasMethod from the given weights
func NewWeight(weights []uint64) (*Alias, error) {

	var total uint64
	var probs []float64 = make([]float64, len(weights))

	for _, weight := range weights {
		total += weight
	}

	if total == 0 {
		return nil, errors.New("weights must be greater than zero")
	}

	for k := range probs {
		probs[k] = float64(weights[k]) / float64(total)
	}

	return NewProb(probs)
}

func setup(probabilities []float64) ([]ele, error) {

	n := len(probabilities)

	if n == 0 {
		return nil, errors.New("probabilities must not be nil")
	}

	avg := 1.0 / float64(n)
	l, m := 0, n-1
	worklist := make([]int, n)

	for i, prob := range probabilities {
		if prob < 0 {
			return nil, errors.New("probability should not be negative")
		}
		if prob > avg {
			worklist[m] = i
			m--
		} else {
			worklist[l] = i
			l++
		}
	}

	eles := make([]ele, n)
	prob := make([]float64, n)

	copy(prob, probabilities)

	for l != 0 && m != n-1 {
		less, more := worklist[l-1], worklist[m+1]
		eles[less] = ele{
			prob: prob[less] * float64(n),
			idx:  more,
		}

		prob[more] = prob[more] + prob[less] - avg
		l--

		if prob[more] < avg {
			worklist[l] = more
			l++
			m++
		}
	}

	for ; l != 0; l-- {
		eles[worklist[l-1]] = ele{prob: 1}
	}

	for ; m != n-1; m++ {
		eles[worklist[m+1]] = ele{prob: 1}
	}

	return eles, nil
}
