package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sort"
)

func main() {
	
	f, err := os.Open("calories.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	elfCalories := []int{}
  scanner := bufio.NewScanner(f)
	calorieCounter := 0

  for scanner.Scan() {
		calories, err := strconv.Atoi(scanner.Text())

		if err != nil {
			elfCalories = append(elfCalories, calorieCounter)
			calorieCounter = 0
			} else {
				calorieCounter = calorieCounter + calories
			}
  }

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	max := 0
	elfIndex := 0

	for index, value := range elfCalories {
		if value > max {
			max = value
		  elfIndex = index
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	fmt.Println ("The top three elves have the following number of calories:", elfCalories[0:3])
	fmt.Println ("The total number calories", elfCalories[0] + elfCalories[1] + elfCalories[2])
	fmt.Printf("Elf of index %v has the most calories with %v\n", elfIndex, max)
}

