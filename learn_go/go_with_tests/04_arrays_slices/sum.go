package main 

// Sum calculates the total from slices of numbers 
func Sum(numbers []int) int {
  add := func(acc, x int) int {return acc+x}
  return Reduce(numbers, add, 0)
}
// Sum all Trails calculates the sum of all but the first number give a collection of slices 
func SumAllTails(numbers ...[]int) []int {
  sum := func(acc. x[]int) []int {
    if len(x) == 0 {
      return append(acc, 0)
    }else {
      tail := x[1:]
      return append(acc, Sum(tail))
    }
  }
  return Reduce(numbers, sumTail, []int{})
}

