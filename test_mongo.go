package main

import (
        "fmt"
        "gopkg.in/mgo.v2"
)

type Record struct {
        URL string
}

func main() {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        c := session.DB("top_sites").C("top_sites")

        result := []Record{}
        c.Find(nil).All(&result)
		for _, v := range result {
			fmt.Println(v.URL)
		}
}
