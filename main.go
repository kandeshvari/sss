package main

import (
	"log"
)

func main() {
	//str := "abc[de[]]f]ab]]][sdf]s"
	s := "a[cd"
	str := s + s

	trees, err := BuildTrees(&str)
	if err != nil {
		log.Fatal(err)
	}

	for i, t := range trees {
		s := ""
		GetString(t.Root, &s)
		log.Printf(">>> %d: %s", i, s)
	}
}
