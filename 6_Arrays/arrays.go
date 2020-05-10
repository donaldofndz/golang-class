package main

import "fmt"

/*

	First of all an array is a collection of the same data type

	The arrays has fixed lengh, so once defined, the siz cannot be increased

*/

func main() {

	//var ourFirstArray [3]int
	var ourSecondarray [3]int
	otherArray := [5]int{0, 9, 8, 7, 6}
	multiLineArray := [3]string{
		"This is the first array",
		"Second array",
		"obviusly the third",
	}
	multiLineDeclarationValue := [...]string{
		"One",
		"Two",
		"Three",
		"emoji?",
		"🤪",
	}

	ourSecondarray[0] = 1
	ourSecondarray[1] = 2
	ourSecondarray[2] = 3

	fmt.Println(ourSecondarray)
	fmt.Println(otherArray)
	fmt.Println(multiLineArray)
	fmt.Print(multiLineDeclarationValue)
	fmt.Println(len(multiLineDeclarationValue))

	a := [3]int{1, 2, 3}
	b := [3]int{1, 3, 3}
	c := [3]int{1, 1, 1}
	d := [3]int{1, 2, 3}
	//e := [...]int{1, 2, 3, 4}

	fmt.Println("a == b", a == b)
	fmt.Println("a == c", a == c)
	fmt.Println("a == b", a == d)
	//fmt.Println("a == e", a == e)

	iteratingOver := [...]int{1, 2, 3, 4, 5, 6, 7}

	for index := 0; index < len(iteratingOver); index++ {
		fmt.Printf("iteratingOver[%v] = %v \n", index, iteratingOver[index])
	}

	for index, value := range iteratingOver {
		fmt.Printf("index: %v, value: %d\n", index, value)
	}

	for _, value := range iteratingOver {
		fmt.Printf("value: %d\n", value)
	}

	multidimensional := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	fmt.Printf("The arry is %v and type o array element is %T \n", multidimensional, multidimensional[0])

	multiShortDeclaration := [...][2]int{{1, 2}, {3, 4}, {5, 6}}

	fmt.Println(multiShortDeclaration)

	for indexDad, child := range multiShortDeclaration {
		for indexSon, elem := range child {
			fmt.Printf("dad: %v, son: %d, elem: %v", indexDad, indexSon, elem)
		}
	}

	//Anotation when you pass an array to a function you are do it by value so you wont see the change outside the function
}
