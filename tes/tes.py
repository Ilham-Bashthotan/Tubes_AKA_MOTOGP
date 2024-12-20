import random
import time

# Fungsi binary search iteratif dengan delay
def binary_search(arr, target):
    left, right = 0, len(arr) - 1
    while left <= right:
        mid = left + (right - left) // 2
        #time.sleep(0.01)  # Delay 10ms untuk simulasi
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return -1  # Jika tidak ditemukan

# Fungsi binary search rekursif dengan delay
def binary_search_recursive(arr, target, left, right):
    if left > right:
        return -1  # Jika tidak ditemukan
    mid = left + (right - left) // 2
    #time.sleep(0.01)  # Delay 10ms untuk simulasi
    if arr[mid] == target:
        return mid
    elif arr[mid] < target:
        return binary_search_recursive(arr, target, mid + 1, right)
    else:
        return binary_search_recursive(arr, target, left, mid - 1)

def main():
    # Membuat array yang lebih besar
    size = 400
    arr = list(range(1, size + 1))  # Mengisi array dengan angka 1 sampai size

    target = 0
    while target < size:
        # Catat waktu sebelum eksekusi
        start = time.time()

        # Panggil fungsi binary search
        result = binary_search(arr, target)

        # Catat waktu setelah eksekusi
        elapsed = (time.time() - start) * 1000  # Waktu dalam milidetik

        target = target + 1

        # Tampilkan hasil
        if result != -1:
            print(f"Target {target} ditemukan di indeks {result}")
        else:
            print(f"Target {target} tidak ditemukan")

        # Tampilkan waktu eksekusi dalam milidetik
        print(f"Waktu eksekusi: {elapsed:.20f} ms")


    # Catat waktu sebelum eksekusi
    start = time.time()

    # Panggil fungsi binary search rekursif
    result = binary_search_recursive(arr, target, 0, size - 1)

    # Catat waktu setelah eksekusi
    elapsed = (time.time() - start) * 1000  # Waktu dalam milidetik

    # Tampilkan hasil
    if result != -1:
        print(f"Target {target} ditemukan di indeks {result}")
    else:
        print(f"Target {target} tidak ditemukan")

    # Tampilkan waktu eksekusi dalam milidetik
    print(f"Waktu eksekusi: {elapsed:.20f} ms")

if __name__ == "__main__":
    main()
