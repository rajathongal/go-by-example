package main

import "fmt"

// go supports anonymous function ex func() int{}
// which can form closures
// Anonymous functions are useful when you want 
// to define a function inline without having to name it

// This function intSeq returns another function which 
// we define anonymously in the body of intSeq
// The returned function closes over the variable i to
// form a closure

func intSeq() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
  // We call intSeq, assigning the result(a function is returned)
  // to nextInt. This function value captures its own i value
  // which will be updated each time we call nextInt
  nextInt := intSeq() 

  // See the effect of closure by calling nextInt a few times
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  // To confirm the state is unique to that particular function
  // create and test a new one
  fmt.Println()

  newInts := intSeq()
  fmt.Println(newInts())
  fmt.Println(newInts())
  fmt.Println(newInts())

}