package sort 

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 要排序的数据必须实现 Sorter 接口
func Sort(data Sorter) {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len() - i; j++ {
			if data.Less(j + 1, j) {
				data.Swap(j, j + 1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	for i := 0; i < data.Len() - 1; i++ {
		if data.Less(i + 1, i) {
			return false
		}
	}

	return true
}

// Convenience types for common cases
type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}

func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortInts(a []int) {
	Sort(IntArray(a))
}

func SortStrings(a []string) {
	Sort(StringArray(a))
}

func IntsAreSorted(a []int) bool {
	return IsSorted(IntArray(a))
}

func StringsAreSorted(a []string) bool {
	return IsSorted(StringArray(a))
}
