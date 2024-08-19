package main

import "fmt"

func main() {
	// map1()
	mapIterator()

	// slice()
	// sliceByLiteral()
	// sliceNil()
	// sliceUse1()
	// sliceAppend()
	// sliceAppend2()
	// sliceAppend3_WithSameLenAndCap()
	// sliceAppend4_expand()
	// sliceIterator()
	// sliceIterator_copy()
	// sliceMutilDim()

	// array()
	// arrayPointer()
	// arrayAssisnment()
	// arrayCopyPointer()
	// arrayMultiDim()
}

func map1() {
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	fmt.Println(colors)

	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	}
}

func mapIterator() {
	colors := map[string]string{"AliceBlue": "#f0f8ff",
		"Coral": "#ff7F50", "DarkGray": "#a9a9a9", "ForestGreen": "#228b22"}
	for key, value := range colors {
		fmt.Println(key, value)
	}
	// delete(colors, "Coral")
	removeColor(colors, "DarkGray")
	fmt.Println()
	for key, value := range colors {
		fmt.Println(key, value)
	}
}

func removeColor(colors map[string]string, key string){
	delete(colors, key)
}


func sliceMutilDim() {
	slice := [][]int{{10}, {20}}
	slice[0] = append(slice[0], 11)
	fmt.Println(slice)
}

func sliceIterator() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Index:%d Value:%d\n", index, value)
	}
}

func sliceIterator_copy() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		//value 拷贝的 slice[index] 值
		fmt.Printf("Value:%d ValueAddr:%X ElementAddr:%X\n", value, &value, &slice[index])
	}
}

func sliceAppend4_expand() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))
}

func sliceAppend3_WithSameLenAndCap() {
	source :=
		[]string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice := source[2:3:3]
	fmt.Println(len(slice), cap(slice))
	fmt.Printf("%p\n", &slice)
	slice2 := append(slice, "kivi")
	fmt.Printf("%p\n", &slice2)

}

func sliceAppend2() {
	source :=
		[]string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice := source[2:3:4]
	fmt.Println(len(slice), cap(slice))

	//panic: runtime error: slice bounds out of range [::6] with capacity 5
	slice2 := source[2:3:6]
	fmt.Println(slice2)
}

func sliceAppend() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]

	newSlice = append(newSlice, 60)
	fmt.Println(slice)
	fmt.Println(newSlice)

	// slice cap grow
	slice = append(slice, 60)
	fmt.Println(slice, len(slice), cap(slice))
}

func sliceUse1() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	newSlice := slice[1:3]
	fmt.Println(newSlice)

	newSlice[1] = 35
	fmt.Println(slice)
	fmt.Println(newSlice)
	//newSlice length is 2
	//panic: runtime error: index out of range [3] with length 2
	// newSlice[3] = 45

}

func slice() {
	slice := make([]string, 3, 5)
	slice = append(slice, "hello")
	fmt.Println(slice)
}

func sliceByLiteral() {
	slice := []string{"hello", "world"}
	fmt.Println(slice)

	slice2 := []string{99: "test 99"}
	fmt.Println(slice2)
}

func sliceNil() {
	slice := make([]int, 0)
	fmt.Println(len(slice), cap(slice))
	slice2 := []int{}
	fmt.Println(len(slice2), cap(slice2))
}

func arrayMultiDim() {
	var array [4][2]int
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array)
}

func arrayCopyPointer() {
	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"
	array1 = array2

	fmt.Println(array1)
	fmt.Println(array2)
}

func arrayAssisnment() {
	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array1 = array2
	fmt.Printf("%p %v", &array1, array1)
	fmt.Printf("%p %v", &array2, array2)

	// var array3 [4]string
	// array3 =  array2
}

func arrayPointer() {
	array := [5]*int{0: new(int), 1: new(int)}
	*array[0] = 10
	*array[1] = 20
	var i int = 30
	array[4] = &i
	for index, value := range array {
		if value != nil {
			fmt.Println(index, *value)
		}
	}
}

func array() {
	array := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array)

	array[2] = 35
	fmt.Println(array)

	// array1 := [...]int{10, 20, 30, 40, 50}
	// fmt.Println(array1)

	// array2 := [5]int{1:10, 2:20}
	// fmt.Println(array2)
}
