package main

import "fmt"
import "reflect"

func main() {

  var x float64 = 3.4
  v := reflect.ValueOf(x)
  fmt.Println("v:", v)
  fmt.Println("type:", v.Type())
  fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
  fmt.Println("value:", v.Float())

}

