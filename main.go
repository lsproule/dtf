package main


import (
	"fmt"
	"math/rand"
)

type Testing struct {
	seed int64;
}


func main() { 
	test := Testing{seed: 1} 
	rand.Seed(test.seed) 
	fmt.Print("Hello, World!") 
}
