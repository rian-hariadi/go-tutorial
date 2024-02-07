package main

import (
	"fmt"
)

func main() {
	var myString string = "resume"
	var indexex = myString[0]

	fmt.Printf("myString: %v \n", myString)

	fmt.Printf("%v, %T\n", indexex, indexex)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Println(myString[1]) //why Not 'e' ??

	var myRun = 'a'
	fmt.Printf("\n myRun : %v", myRun)

	var strSlice = []string{"a", "b", "c", "d", "e"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
		fmt.Println("\n catStr :", catStr)
	}

	fmt.Printf("\n %v", catStr)
	fmt.Printf("\n strSlice[2] %v \n", strSlice[2])

}
