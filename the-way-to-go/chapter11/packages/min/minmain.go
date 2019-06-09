package min

type Element interface{}

type Miner interface {
	Get(index int) Element
	Len() int
	Less(i, j Element) bool
}

func Min(m Miner) Element {
	min := m.Get(0)
	for i := 1; i < m.Len(); i++ {
		if m.Less(m.Get(i), min) {
			min = m.Get(i)
		}
	}

	return min
}

// int
type IntArray []int

func (ia IntArray) Len() int {
	return len(ia)
}

func (ia IntArray) Less(i, j Element) bool {

	if i.(int) < j.(int) {
		return true
	}
	
	return false
}

func (ia IntArray) Get(index int) Element {
	var el Element = ia[index]

	return el
}

func MinInts(ints []int) int {
	ia := IntArray(ints)

	return Min(ia).(int)
}

// float
type FloatArray []float64

func (f FloatArray) Len() int {
	return len(f)
}

func (f FloatArray) Less(i, j Element) bool {

	if i.(float64) < j.(float64) {
		return true
	}
	
	return false
}

func (f FloatArray) Get(index int) Element {
	var el Element = f[index]

	return el
}

func MinFloats(floats []float64) float64 {
	f := FloatArray(floats)

	return Min(f).(float64)
}