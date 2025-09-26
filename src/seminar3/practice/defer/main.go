package main

import (
	"fmt"
	"os"
	"time"
)

func executionOrder() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("normal execution")
}

func trickyDefer() {
	i := 0
	defer fmt.Println("defer i:", i) // Тут внутренности defer выполняются сразу

	defer func() {
		fmt.Println("closure i:", i) // Тут внутренности defer откладываются на потом
	}()

	i = 42
	fmt.Println("after assignment:", i)
}

func namedReturns(x, y int) (result int, err error) {
	defer func() {
		if result > 100 {
			err = fmt.Errorf("result too large: %d", result)
		}
	}()

	result = x * y
	return
}

func processFiles(filenames []string) error {
	var files []*os.File
	defer func() {
		for _, f := range files {
			if f != nil {
				f.Close()
			}
		}
	}()

	for _, name := range filenames {
		f, err := os.Open(name)
		if err != nil {
			return fmt.Errorf("failed to open %s: %w", name, err)
		}
		files = append(files, f)
	}
	return nil
}

func timedOperation(name string) {
	defer func(start time.Time) {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}(time.Now())

	// Simulate work
	time.Sleep(100 * time.Millisecond)
}

func trace(name string) func() {
	fmt.Printf("Entering %s\n", name)
	return func() {
		fmt.Printf("Exiting %s\n", name)
	}
}

func complexFunction() {
	defer trace("complexFunction")()
	defer fmt.Println("Second defer")
	fmt.Println("Working...")
}

func main() {
	executionOrder()
	trickyDefer()
	result, err := namedReturns(101, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	err = processFiles(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}

	timedOperation("processFiles")

	complexFunction()

}
