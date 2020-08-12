package main

import "fmt"

func main(){
//  printStar(5);
//  printReverseStar(5);
//  printDiamond(5);
//  fmt.Println(getFactorial(5));
  fmt.Println(sumOfDigits(10));
}

func printStar(n int){
  for i:=1; i <= n; i++ {
    for j:=0; j<i; j++ {
      fmt.Print("*");
    }
    fmt.Println();
  }
}

func printReverseStar(n int) {
  for i:=1; i<=n; i++ {
    // print spaces
    for j:= n; j>i; j-- {
      fmt.Print(" ");
    }
    for j:= 0; j<i; j++ {
      fmt.Print("*");
    } 
    fmt.Println();
  }  
}

func printDiamond(n int){
  for i:=1; i<=n; i++ {
    // print spaces
    for j:= n; j>i; j-- {
      fmt.Print(" ");
    }
    for j:= 0; j< 2*i; j++ {
      fmt.Print("*");
    } 
    fmt.Println();
  }  
  for i:=n; i>0; i-- {
    // print spaces
    for j:= i; j<n; j++ {
      fmt.Print(" ");
    }
    for j:= 0; j< 2*i; j++ {
      fmt.Print("*");
    } 
    fmt.Println();
  }  
}

func getFactorial(n int) int {
  if n == 0 { return 1 }
  return n*getFactorial(n-1);
}

func sumOfDigits(n int) int {

  var m map[int]string = map[int]string{
    1: "One",
    2: "Two",
    3: "Three",
  }
 
  for k, v := range m {
      fmt.Println(k, v)
  }
}

//   var res int 
//   while n > 0: {
//     d = n % 10
//     res = res + d
//     n = n/10
//   }
//   return res
// }
