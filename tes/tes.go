package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fungsi binary search
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		time.Sleep(100)
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // Jika tidak ditemukan
}

// Fungsi Recursive
func binarySearchRecursive(arr []int, target int, left int, right int) int {
	time.Sleep(100)
	if left > right {
		return -1 // Jika tidak ditemukan
	}

	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, right)
	} else {
		return binarySearchRecursive(arr, target, left, mid-1)
	}
}

func main() {
	// Membuat array yang lebih besar
	size := 300
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1 // Mengisi array dengan angka 1 sampai size
	}

	target := rand.Intn(size) + 1 // Mengambil target acak dari 1 sampai size

	// Catat waktu sebelum eksekusi
	start := time.Now()

	// Panggil fungsi binary search
	result := binarySearch(arr, target)

	// Catat waktu setelah eksekusi
	elapsed := time.Since(start)

	// Tampilkan hasil
	if result != -1 {
		fmt.Printf("Target %d ditemukan di indeks %d\n", target, result)
	} else {
		fmt.Printf("Target %d tidak ditemukan\n", target)
	}

	// Tampilkan waktu eksekusi dalam milidetik
	fmt.Printf("Waktu eksekusi: %d ms\n", elapsed.Nanoseconds())
	// Catat waktu sebelum eksekusi
	start = time.Now()

	// Panggil fungsi binary search rekursif
	result = binarySearchRecursive(arr, target, 0, size-1)

	// Catat waktu setelah eksekusi
	elapsed = time.Since(start)

	// Tampilkan hasil
	if result != -1 {
		fmt.Printf("Target %d ditemukan di indeks %d\n", target, result)
	} else {
		fmt.Printf("Target %d tidak ditemukan\n", target)
	}

	// Tampilkan waktu eksekusi dalam milidetik
	fmt.Printf("Waktu eksekusi: %d ms\n", elapsed.Nanoseconds())
}
