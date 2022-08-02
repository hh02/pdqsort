package pdqsort

import "fmt"

type test interface {
	Len() int
}

type data struct {
	arr []int
}

func (d data) Len() int {
	return len(d.arr)
}

func testData(d test) {
	fmt.Println(d.Len())
}

func Run() {
	var da data
	testData(da)

}

