package main

import (
	"fmt"

	"github.com/rafaelmardones/go-tour/helpers"
)

const numPool = 1000

func GetRandomNum(channel chan int) {
	value := helpers.RandomNumber(numPool)
	channel <- value
}

func main() {
	intChannel := make(chan int) // channel creation
	defer close(intChannel)
	go GetRandomNum(intChannel)
	random := <-intChannel
	fmt.Println(random)
}
