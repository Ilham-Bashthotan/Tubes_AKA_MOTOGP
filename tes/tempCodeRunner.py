import os
import random
import time

# ----------------------------
# Konstanta
# ----------------------------
NMAXRIDER = 2000
NMAXEVENT = 21
NMAXCOMPLEXITY = 400  # 2000 / 5 = 400

# ----------------------------
#  Definisi Kelas
# ----------------------------
class Rider:
    def __init__(self, name="", timeRAC=0, idEvent=0):
        self.name = name
        self.timeRAC = timeRAC
        self.idEvent = idEvent

class Event:
    def __init__(self, idEvent=0, name="", circuit="", date=""):
        self.idEvent = idEvent
        self.name = name
        self.circuit = circuit
        self.date = date

class Complexity:
    def __init__(self, awal=0, tengah=0, akhir=0):
        self.awal = awal
        self.tengah = tengah
        self.akhir = akhir

# ----------------------------
#  Fungsi Bantu
# ----------------------------
def clear_screen():
    if os.name == "nt":
        os.system("cls")
    else:
        os.system("clear")

def inputFrasa(prompt=""):
    return input(prompt)

def randomInt(min_val, max_val):
    return random.randint(min_val, max_val)

def printEvent(dataEvent, nDataEvent):
    for i in range(nDataEvent):
        print("idEvent : ", dataEvent[i].idEvent)
        print("name    : ", dataEvent[i].name)
        print("circuit : ", dataEvent[i].circuit)
        print("date    : ", dataEvent[i].date)
        print("---------------------------")

def printAllRiderByIDEvent(dataRider, nDataRider, ID):
    for i in range(nDataRider):
        if dataRider[i].idEvent == ID:
            print("name    : ", dataRider[i].name)
            print("timeRAC : ", dataRider[i].timeRAC)
            print("---------------------------")

def printComplexity(dataTime, nDataTime, lap):
    for i in range(nDataTime):
        awal   = dataTime[i].awal   // lap
        tengah = dataTime[i].tengah // lap
        akhir  = dataTime[i].akhir  // lap
        print(f"!{awal} @{tengah} #{akhir}")

def addTimeComplexity(dataTime, nDataTime_ptr, awal, tengah, akhir):
    # Pastikan dataTime adalah list of Complexity
    if not all(isinstance(item, Complexity) for item in dataTime):
        raise TypeError("dataTime must be a list of Complexity objects")
    idx = nDataTime_ptr[0]  # posisi index "terakhir"
    dataTime[idx].awal   += awal
    dataTime[idx].tengah += tengah
    dataTime[idx].akhir  += akhir
    nDataTime_ptr[0] += 1

# ----------------------------
#  Algoritma Binary Search
# ----------------------------
def binarySearchIterative(dataRider, nDataRider, timeRAC):
    left = 0
    right = nDataRider - 1
    found = -1

    while left <= right and found == -1:
        # time.sleep(0.001) # Bisa dikomentari jika ingin cepat
        mid = (right + left) // 2
        if dataRider[mid].timeRAC == timeRAC:
            found = mid
        elif dataRider[mid].timeRAC < timeRAC:
            left = mid + 1
        else:
            right = mid - 1

    return found

def binarySearchRecursive(dataRider, timeRAC, left, right):
    if left > right:
        return -1
    # time.sleep(0.001) # Bisa dikomentari jika ingin cepat

    mid = (left + right) // 2
    if dataRider[mid].timeRAC == timeRAC:
        return mid
    elif dataRider[mid].timeRAC < timeRAC:
        return binarySearchRecursive(dataRider, timeRAC, mid+1, right)
    else:
        return binarySearchRecursive(dataRider, timeRAC, left, mid-1)

def binarySearchRecursiveMaster(dataRider, nDataRider, timeRAC):
    return binarySearchRecursive(dataRider, timeRAC, 0, nDataRider-1)

# ----------------------------
#  Menu 1: Input Data
# ----------------------------
def menuInputDataRider(dataEvent, nDataEvent_ptr, dataRider, nDataRider_ptr):
    print("Input Data Rider")
    print("Masukkan data ajang sebanyak 20 + 1 dummy (total 21)")
    for i in range(NMAXEVENT):
        try:
            dataEvent[i].idEvent = int(input("idEvent: "))
        except ValueError:
            dataEvent[i].idEvent = 0

        dataEvent[i].name    = inputFrasa("name: ")
        dataEvent[i].circuit = inputFrasa("circuit: ")
        dataEvent[i].date    = inputFrasa("date: ")
        print("---------------------------")

        nDataEvent_ptr[0] += 1

        nDataEvent_ptr[0] += 1

        if dataEvent[i].idEvent != 0 or dataEvent[i].name or dataEvent[i].circuit or dataEvent[i].date:
            nDataEvent_ptr[0] += 1
    try:
        n = int(input("Masukkan banyak pembalap: "))
    except ValueError:
        n = 0

    for i in range(n):
        try:
            dataRider[i].idEvent = int(input("idEvent: "))
        except ValueError:
            dataRider[i].idEvent = 0

        dataRider[i].name = inputFrasa("name: ")
        try:
            dataRider[i].timeRAC = int(input("timeRAC: "))
        except ValueError:
    print("---------------------------")
    nDataRider_ptr[0] += 1
    nDataRider_ptr[0] = n

# ----------------------------
#  Menu 2: Tampilkan Data Rider
# ----------------------------
def menuTampilkanDataRider(dataEvent, nDataEvent, dataRider, nDataRider):
    print("Tampilkan Data Rider")
    printEvent(dataEvent, nDataEvent)
    try:
    while True:
        try:
            ID = int(input("Pilih idx event yang Anda pilih (0..nDataEvent-1): "))
            if 0 <= ID < nDataEvent:
                break
            else:
                print(f"ID harus dalam rentang 0 hingga {nDataEvent-1}.")
        except ValueError:
            print("Input tidak valid. Harap masukkan angka.")
    clear_screen()
    if 0 <= ID < nDataEvent:
        print("name    : ", dataEvent[ID].name)
        print("circuit : ", dataEvent[ID].circuit)
        print("date    : ", dataEvent[ID].date)
        print("===========================")
        # Di Go, pemanggilan ID+1. Tergantung definisi idEvent
        # Jika event di input sama persis dengan ID, pakai ID.
        # Jika event di input = ID+1, silakan sesuaikan.
        eventID = dataEvent[ID].idEvent
        printAllRiderByIDEvent(dataRider, nDataRider, eventID)
    else:
        print("ID event tidak valid.")

# ----------------------------
#  Menu 3: Pencarian Data Random
# ----------------------------
def menuPencarianDataRandom(dataRider, nDataRider):
    if nDataRider <= 0:
        print("Data Rider masih kosong!")
        return

    print("Pencarian Data Random")
    random_idx = randomInt(0, nDataRider - 1)
    timeRAC = dataRider[random_idx].timeRAC

    print(f"Data yang dicari: {timeRAC}, pada index {random_idx}")
    print("name    : ", dataRider[random_idx].name)
    print("timeRAC : ", dataRider[random_idx].timeRAC)

    print("Pencarian Data (Iterative)")
    idx_iter = binarySearchIterative(dataRider, nDataRider, timeRAC)
    print("Index ditemukan: ", idx_iter if idx_iter != -1 else "Tidak ditemukan")

    print("Pencarian Data (Recursive)")
    idx_recur = binarySearchRecursiveMaster(dataRider, nDataRider, timeRAC)
    print("Index ditemukan: ", idx_recur if idx_recur != -1 else "Tidak ditemukan")

# ----------------------------
#  Menu 4: Kompleksitas Pencarian Data
# ----------------------------
def menuKompleksitasPencarianData(dataRider, nDataRider,
                                  dataTimeIterative, nDataTimeIterative_ptr,
                                  dataTimeRecursive, nDataTimeRecursive_ptr,
                                  lap):
    print("Kompleksitas Pencarian Data")

    # Set ulang ke 0 setiap pemanggilan (sesuai logika di Go)
    nDataTimeIterative_ptr[0] = 0
    nDataTimeRecursive_ptr[0] = 0

    i = 4
    while i < nDataRider:
        n = i + 1  # jumlah data yang digunakan

        # =========== Iterative ===========
        # 1) awal
        start = time.time_ns()
        binarySearchIterative(dataRider, n, dataRider[0].timeRAC)
        awal = time.time_ns() - start

        # 2) tengah
        start = time.time_ns()
        binarySearchIterative(dataRider, n, dataRider[i//2].timeRAC)
        tengah = time.time_ns() - start

        # 3) akhir
        start = time.time_ns()
        binarySearchIterative(dataRider, n, dataRider[i].timeRAC)
        akhir = time.time_ns() - start

        addTimeComplexity(dataTimeIterative, nDataTimeIterative_ptr, awal, tengah, akhir)

        # =========== Recursive ===========
        # 1) awal
        start = time.time_ns()
        binarySearchRecursiveMaster(dataRider, n, dataRider[0].timeRAC)
        awal = time.time_ns() - start

        # 2) tengah
        start = time.time_ns()
        binarySearchRecursiveMaster(dataRider, n, dataRider[i//2].timeRAC)
        tengah = time.time_ns() - start

        # 3) akhir
        start = time.time_ns()
        binarySearchRecursiveMaster(dataRider, n, dataRider[i].timeRAC)
        akhir = time.time_ns() - start

        addTimeComplexity(dataTimeRecursive, nDataTimeRecursive_ptr, awal, tengah, akhir)

        i += 5

    # Tampilkan hasil
    print("Kompleksitas Pencarian Data (Iterative)")
    printComplexity(dataTimeIterative, nDataTimeIterative_ptr[0], lap)
    print("Kompleksitas Pencarian Data (Recursive)")
    printComplexity(dataTimeRecursive, nDataTimeRecursive_ptr[0], lap)

# ----------------------------
#  Fungsi Main
# ----------------------------
def main():
    dataEvent = [Event() for _ in range(NMAXEVENT)]
    dataRider = [Rider() for _ in range(NMAXRIDER)]

    dataTimeIterative = [Complexity() for _ in range(NMAXCOMPLEXITY)]
    dataTimeRecursive = [Complexity() for _ in range(NMAXCOMPLEXITY)]

    nDataEvent_ptr = [0]
    nDataRider_ptr = [0]

    nDataTimeIterative_ptr = [0]
    nDataTimeRecursive_ptr = [0]

    lap = 1

    while True:
        clear_screen()
        print("--------------------------")
        print("          M E N U         ")
        print("--------------------------")
        print("1. Input Data Rider")
        print("2. Tampilkan Data Rider")
        print("3. Pencarian Data Random")
        print("4. Kompleksitas Pencarian Data")
        print("5. Keluar")
        print("--------------------------")

        menu_choice = inputFrasa("Pilih (1-5): ")

        if menu_choice == "1":
            clear_screen()
            menuInputDataRider(dataEvent, nDataEvent_ptr, dataRider, nDataRider_ptr)

        elif menu_choice == "2":
            clear_screen()
            menuTampilkanDataRider(dataEvent, nDataEvent_ptr[0], dataRider, nDataRider_ptr[0])

        elif menu_choice == "3":
            clear_screen()
            menuPencarianDataRandom(dataRider, nDataRider_ptr[0])

        elif menu_choice == "4":
            # Seperti di Go, loop 100 kali (bisa diubah)
            for _ in range(100):
                clear_screen()
                menuKompleksitasPencarianData(
                    dataRider,
                    nDataRider_ptr[0],
                    dataTimeIterative,
                    nDataTimeIterative_ptr,
                    dataTimeRecursive,
                    nDataTimeRecursive_ptr,
                    lap
                )
                lap += 1

        elif menu_choice == "5":
            print("Program selesai. Terima kasih.")
            break

        else:
            print("Pilihan tidak valid.")
            break

        else:
            print("Pilihan tidak valid.")

        input("\nTekan ENTER untuk melanjutkan...")

if __name__ == "__main__":
    main()