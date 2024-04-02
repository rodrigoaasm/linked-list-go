package structs

import (
	"fmt"
	"math/rand"
	"time"
)

func Benchmark(len int) {
	lessFn := func(contentA int, contentB int) bool {
		return contentA < contentB
	}

	listInsertionSort := NewLinkedList[int]()
	listMergeSerial := NewLinkedList[int]()
	listMergeSerialWithInsertion := NewLinkedList[int]()
	listMergeParallel := NewLinkedList[int]()
	listMergeParallelWithInsertion := NewLinkedList[int]()

	randgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		n := randgen.Intn(len * 10)
		listInsertionSort.Push(n)
		listMergeSerial.Push(n)
		listMergeSerialWithInsertion.Push(n)
		listMergeParallel.Push(n)
		listMergeParallelWithInsertion.Push(n)
	}

	start := time.Now()
	insertionSort(listInsertionSort.head, lessFn)
	fmt.Println("Insertion Sort:", time.Since(start))

	start = time.Now()
	mergeSort(listMergeSerial.head, uint32(len), lessFn, false)
	fmt.Println("Merge Sort Serial:", time.Since(start))

	start = time.Now()
	mergeSort(listMergeSerialWithInsertion.head, uint32(len), lessFn, true)
	fmt.Println("Merge Sort Serial with Insertion sort:", time.Since(start))

	start = time.Now()
	mergeSortParallel(listMergeParallel.head, uint32(len), lessFn, false)
	fmt.Println("Merge Sort Parallel:", time.Since(start))

	start = time.Now()
	mergeSortParallel(listMergeParallelWithInsertion.head, uint32(len), lessFn, true)
	fmt.Println("Merge Sort Parallel with Insertion Sort:", time.Since(start))
}
