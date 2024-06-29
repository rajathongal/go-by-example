package main

import "fmt"

func main() {
  // range iterates over elements in a variety of data structures.

  // Here we use range to sum the numbers in a slice.
  // Arrays work like this too.
  nums := []int{2,3,4}
  sum := 0

  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)

  // range on arrays and slices provides both the index and value for each entry
  // Above we did not need index, so we ignored it with the blank identifier _.
  for i, num := range nums {
    if(num == 3) {
      fmt.Println("index:", i)
    }
  }

  // range on map iterates over key/value pairs
  kvs := map[string]string{"a": "apple", "b": "banana"}

  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k,v)
  }

  // range can also iterate over just keys of a map
  for k := range kvs {
    fmt.Println("key:", k)
  }

  // range can also iterate over just values of a map
  for _, v := range kvs {
    fmt.Println("value:", v)
  }

  // range on strings iterates over unicode code points.
  // The first value is the starting byte index of the rune and
  // the second rune itself
  for i, c := range "Ago" {
    fmt.Println(i,c)
  }
}