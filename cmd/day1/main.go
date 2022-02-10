package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int) {
	*i = Int(n)
}

func main() {
	funny(2)
}

func funny(n int) {
	defer fmt.Println(n)
	defer func() {
		fmt.Println(n)
	}()

	n *= n

	fmt.Println(n)
}

func openCSVAndCount() {
	f, err := os.Open("./oscar_male.csv")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	const nameColumn = 3
	nameCount := map[string]int{}
	for _, record := range records {
		nameCount[record[nameColumn]]++
	}

	for name, count := range nameCount {
		if count > 1 {
			fmt.Println(name)
		}
	}
}

func mappingtype() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	fmt.Println(m == nil)

	for k, v := range m {
		fmt.Println(k, v)
	}

	if v, ok := m["b"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("nothing")
	}

}

func varidic(d ...int) {
	for _, i := range d {
		fmt.Println(i)
	}
}

func couple(s string) []string {
	r := []string{}
	for s += "*"; len(s) > 1; s = s[2:] {
		r = append(r, s[:2])
	}

	return r
}

func printPrime(n int) {
	for i := 2; i <= n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Println(i)
		}
	}
}

func IsCorrect() bool {
	return true
}

func swap(a, b int) (int, int) {
	return b, a
}

func squareArea(a float64) float64 {
	return a * a
}
