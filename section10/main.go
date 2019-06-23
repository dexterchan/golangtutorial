package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(len(x))
	slicefunc()

	sliceAslice()

	sliceAppend()

	sliceMake()
	multiDimSlice()
	mapFunc()
	mapAddElement()
}

func slicefunc() {
	//x := type {values} //composite literal
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)

	for i, v := range x[2:] {
		fmt.Printf("%d : %v\n", i, v)
	}
}

//a SLICE allows you to group values of same type

func sliceAslice() {
	//SLICE make on top of array.
	//SLICEN is dynamic
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x[1:])
	fmt.Println(x[1:4])
}

func sliceAppend() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	y := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(append(x, 100, 200, 300))
	fmt.Println(append(x, y...))

	//Delete
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func sliceMake() {
	x := make([]int, 12, 12)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < 10; i++ {
		x[i] = i
	}
	x = append(x, 11, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 13, 14)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func multiDimSlice() {
	var row1 []int
	row1 = []int{0, 1, 2, 3, 4, 5}
	row2 := []int{3, 4, 5, 6, 7, 8}
	m := [][]int{row1, row2}

	fmt.Println(m)
}

func mapFunc() {
	m := map[string]int{
		"apple":  1,
		"orange": 2,
	}
	fmt.Println(m["apple"])
	fmt.Println(m["banana"])

	if v, ok := m["banana"]; ok {
		fmt.Printf("value of found key'banana': %d , flag:%v\n", v, ok)
	} else {
		fmt.Printf("value of not found key'banana': %d , flag:%v\n", v, ok)
	}

	if v, ok := m["orange"]; ok {
		fmt.Printf("value of found key'orange': %d , flag:%v\n", v, ok)
	} else {
		fmt.Printf("value of not found key'orange': %d , flag:%v\n", v, ok)
	}
}

func mapAddElement() {
	m := map[string][]int{
		"apple":  []int{1, 2},
		"orange": []int{2, 4},
	}
	m["berry"] = []int{9, 10}

	for k, v := range m {
		fmt.Printf("key:%v \t value:%v\n", k, v)
	}

	fmt.Println("demo delete")

	delete(m, "berry")
	for k, v := range m {
		fmt.Printf("key:%v \t value:%v\n", k, v)
	}
}
