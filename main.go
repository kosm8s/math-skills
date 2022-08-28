package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Can not open file")
		return
	}
	defer file.Close()
	var a []int
	var curr int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		a = append(a, curr)
	}
	res1 := Average(a)
	res2 := Median(a)
	res3 := Variance(a)
	res4 := StandardDeviation(a)
	fmt.Printf("Average: %.f\n", res1)
	fmt.Printf("Median: %.f\n", res2)
	fmt.Printf("Variance: %.f\n", res3)
	fmt.Printf("Standard Deviation: %.f\n", res4)
}

func Average(x []int) float64 {
	var result, count float64
	for i := 0; i < len(x); i++ {
		result += float64(x[i])
		count++
	}
	// for _, w := range x {
	// 	result += w
	// 	count++
	// }
	result /= count
	return float64(result)
}

func Median(x []int) float64 {
	sort.Ints(x)
	if len(x)%2 != 0 {
		return math.Round(float64(x[len(x)/2]))
	} else {
		return math.Round(float64((x[len(x)/2] + x[len(x)/2-1])) / 2)
	}
}

func Variance(x []int) float64 {
	aver := Average(x)
	z := 0.0
	for i := 0; i < len(x); i++ {
		z += math.Pow(float64(x[i])-aver, 2)
	}
	return z / float64(len(x))
}

func StandardDeviation(x []int) float64 {
	return math.Sqrt(Variance(x))
}
