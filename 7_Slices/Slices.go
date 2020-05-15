package main

import "fmt"

func main() {
	// Okey, so first of all we need to know the difference between an array and an slice:

	// An slice is basiclay an array but it works different because is a pointer to and array (justo a reference to an array), we will later talk about pointer but for know we are going to learn about slices

	//var nilSlice []int
	//fmt.Printf("When a slice is declared it doesnt get the 0 value because it is just a reference, insted it gets null: %v \n", nilSlice == nil)

	//ourFirstSlice()
	//refereceToSlice()
	//capacityAndLength()
	//explainingAppend()
	//explainingAppendTwo()
	//makeExample()
	//spreadOperator()
	readValue()

}

func ourFirstSlice() {
	var s []int
	fmt.Printf("Is nil? %v \n", s == nil)

	normalArray := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = normalArray[2:4]

	fmt.Println(s)
}

func refereceToSlice() {
	var s []int
	normalArray := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = normalArray[2:4]

	s[0] = 33
	s[1] = 44

	fmt.Println(normalArray)
}

func capacityAndLength() {
	var s []int
	normalArray := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = normalArray[2:4]

	fmt.Printf("len: %v, cap: %v", len(s), cap(s))
}

func explainingAppend() {
	normalArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := normalArray[2:4]
	newSlice := append(s, 55, 66)

	fmt.Printf("slice = %v , newSlice = %v \n", s, newSlice)
	fmt.Printf("slice len = %v , slice cap = %v \n", len(s), cap(s))
	fmt.Printf("newSlice len = %v , newSlice cap = %v \n", len(newSlice), cap(newSlice))
	fmt.Println(normalArray)

}

func explainingAppendTwo() {

	//This is how we declare an anonymus slice

	aS := []int{1, 2, 3, 4, 5}

	fmt.Printf("This is the Slice: %v\n", aS)
	fmt.Printf("Slice len: %v, cap %v\n", len(aS), cap(aS))

	aS = append(aS, 66, 77, 88)

	fmt.Printf("This is the Slice: %v\n", aS)
	fmt.Printf("Slice len: %v, cap %v\n", len(aS), cap(aS))

	aS = append(aS, 99, 100, 110)

	fmt.Printf("This is the Slice: %v\n", aS)
	fmt.Printf("Slice len: %v, cap %v\n", len(aS), cap(aS))

}

func makeExample() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6, 7}
	s4 := []int{1, 2, 3}

	n1 := copy(s1, s2)
	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("s1 == nil", s1 == nil)

	n2 := copy(s2, s3)
	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3)

	n3 := copy(s3, s4)
	fmt.Printf("n3=%d, s3=%v, s4=%v\n", n3, s3, s4)
}

func spreadOperator() {

	s1 := make([]int, 0, 10)
	fmt.Println("before -> s1=", s1)

	s2 := []int{1, 2, 3}
	s1 = append(s1, s2...)
	fmt.Println("after -> s1=", s1)
}
