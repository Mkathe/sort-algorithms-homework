package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	unsortedNumbers := GenerateArray(1000)

	//We need a data structure for graph
	results := make(map[string]float64)

	// BubbleSort
	numsForBubble := append([]int(nil), unsortedNumbers...)
	startBubbleSortTimer := time.Now()
	BubbleSort(numsForBubble)
	elapsedBubble := time.Since(startBubbleSortTimer).Seconds() * 1000 // в мс
	results["Bubble Sort O(n^2)"] = elapsedBubble
	fmt.Printf("Bubble Sort took: %.4f ms\n", elapsedBubble)

	// QuickSort
	numsForQuick := append([]int(nil), unsortedNumbers...)
	startQuickSortTimer := time.Now()
	QuickSort(numsForQuick)
	elapsedQuick := time.Since(startQuickSortTimer).Seconds() * 1000
	results["Quick Sort O(n log n)"] = elapsedQuick
	fmt.Printf("Quick Sort took: %.4f ms\n", elapsedQuick)

	// InsertionSort
	numsForIns := append([]int(nil), unsortedNumbers...)
	startInsertionSortTimer := time.Now()
	InsertionSort(numsForIns)
	elapsedInsertion := time.Since(startInsertionSortTimer).Seconds() * 1000
	results["Insertion Sort O(n^2)"] = elapsedInsertion
	fmt.Printf("Insertion Sort took: %.4f ms\n", elapsedInsertion)

	BuildChart(results)
}

func GenerateArray(size int) []int {
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(size)
	}

	return numbers
}

func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	} else {
		pivot := nums[0]
		less := []int{}
		greater := []int{}

		for _, num := range nums[1:] {
			if num <= pivot {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		return append(append(QuickSort(less), pivot), QuickSort(greater)...)
	}
}

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func InsertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && key < nums[j] {
			nums[j+1] = nums[j]
			j = j - 1
		}
		nums[j+1] = key
	}

	return nums
}

func BuildChart(results map[string]float64) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Sorting Algorithm Performance",
			Subtitle: "ms",
		}),
	)

	names := []string{}
	values := []opts.BarData{}
	for name, val := range results {
		names = append(names, name)
		values = append(values, opts.BarData{Value: val})
	}

	bar.SetXAxis(names).AddSeries("Time (ms)", values)

	f, _ := os.Create("sorting_results.html")
	defer f.Close()
	err := bar.Render(f)
	if err != nil {
		panic(err)
	}
}
