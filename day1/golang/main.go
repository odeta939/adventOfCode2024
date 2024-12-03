package main

/*
read file
split to chars
1. make 2 arrays
2. sort
3. compare
4. sum the difference
*/

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left := []int{}
	right := []int{}

	file, err := os.Open("../input.txt") // Replace with your file path
		if err != nil {
			log.Fatal(err)
		}

	defer file.Close()

	// need to separate string into 2 arrays
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		
		leftInt, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatal(err)
		}
		rightInt, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
		}

		left = append(left,leftInt)
		right = append(right, rightInt)

	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(sum)
}

