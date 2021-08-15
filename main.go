package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
)

func main() {
	var (
		n = flag.Int("n", 100, "find the primes between 1, n")
		p = flag.Int("p", 0, "the number of P")
		s = flag.Bool("s", true, "show the prime number")
	)
	flag.Parse()
	if *n < 1 {
		fmt.Fprintln(os.Stderr, "n should greate than 0")
		os.Exit(1)
	}
	if *p > 0 {
		runtime.GOMAXPROCS(*p)
	}

	fmt.Printf("n: %d, NumCPU: %d, GOMAXPROCS: %d\n", *n, runtime.NumCPU(), runtime.GOMAXPROCS(0))
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(*n)
	for i := 1; i <= *n; i++ {
		go func(i int) {
			defer wg.Done()
			if isPrime(i) {
				ch <- i
			}
		}(i)
	}

	go func() {
		for i := range ch {
			if *s {
				fmt.Println(i)
			} else {
				_ = i
			}
		}
	}()
	wg.Wait()
	close(ch)
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
