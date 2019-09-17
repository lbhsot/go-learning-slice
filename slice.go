package main

import "fmt"

func simpleSlice() {
//     Why code below not work???????
//     myArray [10]int := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//     mySlice []int = myArray[:5]
    var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var mySlice []int = myArray[:5]
    fmt.Println("Elements of myArray: ")
    for _, v := range myArray {
        fmt.Println(v, " ")
    }

    fmt.Println("\nElements of mySlice: ")
    for _, v := range mySlice {
        fmt.Println(v, " ")
    }

    fmt.Println()
}

func sliceFunc() {
    mySlice := make([]int, 5, 10)
    fmt.Println("len(mySlice): ", len(mySlice))
    fmt.Println("cap(mySlice): ", cap(mySlice))
    fmt.Println("mySlice: ", mySlice)
    mySlice2 := []int{8, 9, 10}
    mySlice = append(mySlice, mySlice2...)
    fmt.Println("new mySlice: ", mySlice)
    fmt.Println("new len(mySlice): ", len(mySlice))
    fmt.Println("new cap(mySlice): ", cap(mySlice))
    mySlice = append(mySlice, 12, 12, 32)
    fmt.Println("Memory not enough issue will be managed by slice automatically. And now new mySlice: ", mySlice)
    fmt.Println("And len(mySlice): ", len(mySlice))
    fmt.Println("And cap(mySlice): ", cap(mySlice))
//     Code below not work
//     mySlice = append(mySlice, 12, 12, 32, mySlice2...)
}

func createSliceFromOldOne() {
    oldSlice := [10]int{1, 2, 3, 4, 5}
    fmt.Println("len(oldSlice): ", len(oldSlice))
    fmt.Println("cap(oldSlice): ", cap(oldSlice))
    newSlice := oldSlice[:3]
    fmt.Println("newSlice: ", newSlice)
    newSlice = oldSlice[:6]
    fmt.Println("newSlice: ", newSlice)
//     Code below will not work. Error: invalid slice index 11 (out of bounds for 10-element array)
//     newSlice = oldSlice[:11]
//     fmt.Println("newSlice: ", newSlice)
}

func main() {
//     simpleSlice()
//     sliceFunc()
    createSliceFromOldOne()
}

