package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	oddCh := make(chan int)
	evenCh := make(chan int)

	go even(evenCh)
	go odd(oddCh)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
}

func even(ch chan int) {
	for {
		fmt.Println("even of", <-ch)
	}
}
func odd(ch chan int) {
	for {
		fmt.Println("odd of", <-ch)
	}
}

type Golang struct {
	gorm.Model
	Title string
}

func mainDBConnect() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=myapp port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Golang{})
}

func mainGoFibo() {
	ch := make(chan int)
	quitCh := make(chan struct{})

	go fibonacciGo(ch, quitCh)

	for i := 0; i < 12; i++ {
		fmt.Println(<-ch)
	}

	quitCh <- struct{}{}
}

func fibonacciGo(ch chan int, quitCh chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-quitCh:
			fmt.Println("stop")
			return
		}
	}
}

func mainOfFibonacci() {
	fn := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}

func fibonacci() func() int {
	f0, f1 := 0, 1
	return func() int {
		defer func() {
			f0, f1 = f1, f0+f1
		}()
		return f0
	}
}

type IntnFunc func(int) int

func (fn IntnFunc) Intn(n int) int {
	return fn(n)
}

func add(a, b int) int {
	return a + b
}

func newAddFunc() func(int, int) int {
	return add
}

type CancelFunc func()

func WithCancelFunc() CancelFunc {
	return func() {
		fmt.Println("cancel")
	}
}

func boom() (func() int, func()) {
	i := 0
	return func() int {
			return i
		}, func() {
			i++
		}
}

func printSum(add func(int, int) int) {
	fmt.Println(add(1, 2))
}

type rectangle struct {
	w, h float64
}

func (r rectangle) area() float64 {
	return r.w * r.h
}

type triangle struct {
	b, h float64
}

func (t triangle) area() float64 {
	return t.b * t.h * 0.5
}

type areaer interface {
	area() float64
}

func printArea(r areaer) {
	fmt.Printf("area is %v\n", r.area())
}

func printAreaMain() {
	rec := rectangle{w: 4, h: 2}
	printArea(rec)
	tri := triangle{b: 4, h: 5}
	printArea(tri)
}

func switchcase() {
	var n int = 2

	switch {
	case n%2 == 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}
}

func emptyInterface() {
	var i interface{}

	fmt.Println(i == nil)

	i = 10
	fmt.Printf("%T %v\n", i, i)

	i = "ten"
	fmt.Printf("%T %v\n", i, i)

	if s, ok := i.(int); ok {
		fmt.Printf("%v is int\n", s)
	}
	if s, ok := i.(string); ok {
		fmt.Printf("%v is string\n", s)
	}
}

func deferDemo() {
	defer func() {
		fmt.Println("in defer func")
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var err error
	fmt.Println(err.Error())
}
