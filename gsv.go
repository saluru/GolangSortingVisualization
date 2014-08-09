package main

import (
	"bytes"
	cryptoRand "crypto/rand"
	"flag"
	"fmt"
	"time"
)

var max int
var fps int
var count int
var mode int

func randomArray(n int, max int) []int {
	var i int
	var number float64
	arr := make([]int, n)

	for i = 0; i < n; i++ {
		b := make([]byte, 1)
		cryptoRand.Read(b)
		number = float64(b[0])
		arr[i] = int(number / 255 * float64(max))
	}
	return arr
}

func visualize(arr []int) {
	var buffer bytes.Buffer
	var x int
	var y int

	for y = 0; y < max; y++ {
		for x = 0; x < count; x++ {
			if arr[x] == y {
				buffer.WriteByte(byte('#'))
			} else if arr[x] < y && mode == 1 {
				buffer.WriteByte(byte('#'))
			} else if arr[x] > y && mode == 2 {
				buffer.WriteByte(byte('#'))
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}
	time.Sleep(time.Second / time.Duration(fps))
	fmt.Print("\033[2J")
	fmt.Print(buffer.String())
}

func bubbleSort(arr []int) {
	var i int
	var j int

	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			visualize(arr)
		}
		visualize(arr)
	}
}

func combSort(arr []int) {
	var gap int = len(arr)
	var swapped bool = false
	var i int

	for gap > 1 || swapped == true {
		swapped = false
		if gap > 1 {
			gap = int(float64(gap) / 1.3)
		}
		for i = 0; i < len(arr)-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
			visualize(arr)
		}
		visualize(arr)
	}
}

func gnomeSort(arr []int) {
	var i int = 1

	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 {
				i--
			}
		}
		visualize(arr)
	}
}

func insertionSort(arr []int) {
	var i int
	var j int

	for i = 0; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j = j - 1
			visualize(arr)
		}
		visualize(arr)
	}
}

func oddEvenSort(arr []int) {
	var sorted bool = false
	var i int

	for !sorted {
		sorted = true
		for i = 1; i < len(arr) - 1; i += 2 {
			if arr[i] > arr[i + 1] {
				arr[i], arr[i + 1] = arr[i + 1], arr[i]
				sorted = false
			}
			visualize(arr)
		}
		for i = 0; i < len(arr) - 1; i += 2 {
			if arr[i] > arr[i + 1] {
				arr[i], arr[i + 1] = arr[i + 1], arr[i]
				sorted = false
			}
			visualize(arr)
		}
		visualize(arr)
	}
}

func selectionSort(arr []int) {
	var min int = 0
	var i int
	var j int

	for i = 0; i < len(arr); i++ {
		min = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
				visualize(arr)
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		visualize(arr)
	}
}

func main() {
	var algo string
	flag.StringVar(&algo, "algo", "bubble", "Select sorting algorithm all/[bubble]/comb/gnome/insertion/oddEven/selection")
	flag.IntVar(&fps, "fps", 10, "frames per second")
	flag.IntVar(&max, "max", 9, "highest value")
	flag.IntVar(&count, "count", 30, "number of values")
	flag.IntVar(&mode, "mode", 1, "visualization mode")
	flag.Parse()
	arr := randomArray(count, max)
	fmt.Printf("sorting via %v-sort\nhighest value: %v\nnumber of values: %v\n\n", algo, max, count)
	time.Sleep(time.Second * 1)
	switch algo {
	case "bubble":
		bubbleSort(arr)
	case "comb":
		combSort(arr)
	case "gnome":
		gnomeSort(arr)
	case "insertion":
		insertionSort(arr)
	case "oddEven":
		oddEvenSort(arr)
	case "selection":
		selectionSort(arr)
	case "all":
		arr = randomArray(count, max)
		bubbleSort(arr)
		arr = randomArray(count, max)
		combSort(arr)
		arr = randomArray(count, max)
		gnomeSort(arr)
		arr = randomArray(count, max)
		insertionSort(arr)
		arr = randomArray(count, max)
		oddEvenSort(arr)
		arr = randomArray(count, max)
		selectionSort(arr)
	}
}
