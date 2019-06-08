package float

import (
	"math/rand"
	"strconv"
)

type Float64Array []float64

func (p Float64Array) Len() int {
	return len(p)
}

func (p Float64Array) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Float64Array) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Float64Array) Fill(length int) {
	*p = make(Float64Array, length)
	for i := 0; i < length; i++ {
		(*p)[i] = rand.Float64() * 10
	}
}

func (p Float64Array) List() string {
	var formatString string
	for i := 0; i < p.Len(); i++ {
		formatString += strconv.FormatFloat(p[i], 'f', 2, 64) + " : "
	}

	return formatString
}

func (p Float64Array) String() string {
	return p.List()
}