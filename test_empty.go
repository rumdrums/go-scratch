package main

import "fmt"
//import "reflect"

func main() {

  var v []interface{}
//  v := reflect.ValueOf(x)
  v = append(v, 1)
  fmt.Println("v:", v[0])
//  fmt.Println("type:", v.Type())
//  fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
//  fmt.Println("value:", v.Float())

}

