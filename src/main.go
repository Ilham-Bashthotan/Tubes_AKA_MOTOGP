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

func binarySearchRecursive(arr []int, target int, left int, right int) int {
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
	size := 1000000000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1 // Mengisi array dengan angka 1 sampai size
	}

	target := rand.Intn(size) + 1 // Mengambil target acak dari 1 sampai size

	// Lakukan pengulangan untuk mendapatkan rata-rata waktu eksekusi
	numIterations := 100
	totalDuration := time.Duration(0)

	for i := 0; i < numIterations; i++ {
		// Catat waktu sebelum eksekusi
		start := time.Now()

		// Panggil fungsi binary search
		_ = binarySearch(arr, target)

		// Catat waktu setelah eksekusi
		elapsed := time.Since(start)
		totalDuration += elapsed
	}

	// Hitung rata-rata waktu eksekusi
	averageDuration := totalDuration / time.Duration(numIterations)

	// Tampilkan hasil
	fmt.Printf("Target %d dicari dalam array berukuran %d\n", target, size)
	fmt.Printf("Rata-rata waktu eksekusi: %d ms\n", averageDuration.Milliseconds())
}
