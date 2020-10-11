package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	exercises := []string{"pushups", "situps", "crunches", "squats", "lunges", "pullups"}
	fmt.Println(exercises[rand.Intn(len(exercises))])
}
