package main

import (
	"fmt"

	"github.com/valerii-smirnov/my-pet-project/structs/2/models"
)

type age int

func (a age) isAdult() bool {
	return a >= 18
}

func main() {
	myAge := age(20)
	myAge.isAdult()

	fmt.Println(myAge)

	data := models.Collection{7, 88, 33,23, 52, 64, 67, 99, 3, 112, 1}
	fmt.Println(data.Evens().Sort().Max())
}
