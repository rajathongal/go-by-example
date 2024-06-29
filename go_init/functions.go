package main
import "fmt"

// Functions are central in go

// Here is a function that takes two ints and returns 
// their sum as int
func plus(a int, b int) int {
  // Go requires explicit returns i.e. it wont automatically 
  // return the value of last expression
  return a + b
}

// when you have multiple consecutive parameters of the same type 
// you may omit the type name for the like- typed parameters
// up to the final parameter that declares the type 
func plusPlus(a,b,c int) int {
  return a + b + c
}

// Multiple Return Values
// Go has built-in support for multiple return values.
// This feature is used often in idomatic Go,
// for example to return both result and error values from a function

// The (int, int) in this function signature shows that the function returns 2 int values
func vals() (int, int) {
  return 3, 7
}

func main() {
  // call a function just as you'd expect with func-name(args)
  res := plus(1,2)
  fmt.Println("1+2 =", res)

  res = plusPlus(1,2,3)
  fmt.Println("1+2+3 =", res)

  // Here we use the 2 different return values from the call
  // with multiple assignment
  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  // if you want a subset of the retured values, use blank identifier _
  _, c := vals()
  d, _ := vals()
  fmt.Println(c)
  fmt.Println(d)
}

