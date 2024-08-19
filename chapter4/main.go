package main

import "fmt"

func main() {
	// slice()
	sliceByLiteral()

	// array()
	// arrayPointer()
	// arrayAssisnment()
	// arrayCopyPointer()
	// arrayMultiDim()
}

func slice(){
	slice := make([]string, 3, 5)
	slice = append(slice, "hello")
	fmt.Println(slice)
}

func sliceByLiteral(){
	slice := []string{"hello", "world"} 
	fmt.Println(slice)
}


func arrayMultiDim(){
	var array [4][2]int 
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, { 40, 41}};
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	array = [4][2]int{1: {0:20}, 3:{1:41}}
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
