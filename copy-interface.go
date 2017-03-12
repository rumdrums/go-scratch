package main

import "fmt"

type Thing interface {
  Speak() string
  Copy(Thing)
}

type Animal struct {
  utterance string
}

type Person struct {
  utterance string
  dance	string
}

func Say(t Thing) {
  fmt.Println(t.Speak())
}

func (a Animal) Speak() string {
 return a.utterance
}

func (a Animal) Copy(t Thing) {
  u := t.(*Animal)
  u.utterance = a.utterance
}

func (p Person) Speak() string {
 return p.utterance
}

func (p Person) Copy(t Thing) {
  u := t.(*Person)
  u.dance = p.dance
  u.utterance = p.utterance
}

func (p Person) Dance() string {
 return p.dance
}

func main() {
  p := Person{utterance: "hello", dance: "dancing"}
  a := Animal{"woof"}
  var q Person
  r := &p
  var b Animal
  p.Copy(&q)
  a.Copy(&b)
  p.utterance = "goodbye"
  a.utterance = "meow"
  Say(p)
  Say(q)
  Say(r)
  Say(a)
  Say(b)
}
