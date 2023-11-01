package main

import (
    "fmt"
    "math/rand"
    "time"
)

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Swap arr[j] and arr[j+1]
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    // Initialize random seed
    rand.Seed(time.Now().UnixNano())

    // Create an array with 1000 elements
    arr := make([]int, 1000)

    // Fill the array with random values
    for i := 0; i < 1000; i++ {
        arr[i] = rand.Intn(1000) // Fill with random integers between 0 and 999
    }

    // Sort the array using bubble sort
    bubbleSort(arr)

    // Print the sorted array in a single line
    fmt.Print("Sorted Array: ")
    for i := 0; i < 1000; i++ {
        fmt.Printf("%d ", arr[i])
    }
    fmt.Println()
}

