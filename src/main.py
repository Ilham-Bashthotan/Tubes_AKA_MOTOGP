import random
import os
import time

NMAXrider = 2000
NMAXevent = 21
NMAXcomplexity = 400  # 2000 / 5 = 400

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

def main():
    dataEvent = [Event() for _ in range(NMAXevent)]
    dataRider = [Rider() for _ in range(NMAXrider)]
    dataTimeIterative = [Complexity() for _ in range(NMAXcomplexity)]
    dataTimeRecursive = [Complexity() for _ in range(NMAXcomplexity)]
    nDataEvent = 0
    nDataRider = 0
    nDataTimeIterative = 0
    nDataTimeRecursive = 0
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
        pilih = input("Pilih (1-4): ")
        if pilih == "1":
            clear_screen()
            nDataEvent, nDataRider = menuInputDataRider(dataEvent, nDataEvent, dataRider, nDataRider)
        elif pilih == "2":
            clear_screen()
            menuTampilkanDataRider(dataEvent, nDataEvent, dataRider, nDataRider)
        elif pilih == "3":
            clear_screen()
            menuPencarianDataRandom(dataRider, nDataRider)
        elif pilih == "4":
            for _ in range(100):  # Change 1 to 100
                clear_screen()
                nDataTimeIterative, nDataTimeRecursive = menuKompleksitasPencarianData(dataRider, nDataRider, dataTimeIterative, nDataTimeIterative, dataTimeRecursive, nDataTimeRecursive, lap)
                lap += 1
        elif pilih == "5":
            break
        input("Tekan apapun untuk melanjutkan")
        clear_screen()

def menuInputDataRider(dataEvent, nDataEvent, dataRider, nDataRider):
    print("Input Data Rider")
    print("Masukkan data ajang sebanyak 20 + 1 dummy")
    for i in range(21):
        dataEvent[i].idEvent = int(input("idEvent: "))
        dataEvent[i].name = input("name: ")
        dataEvent[i].circuit = input("circuit: ")
        dataEvent[i].date = input("date: ")
        print("---------------------------")
        nDataEvent += 1
    n = int(input("Masukkan banyak pembalap: "))
    for i in range(n):
        dataRider[i].idEvent = int(input("idEvent: "))
        dataRider[i].name = input("name: ")
        dataRider[i].timeRAC = int(input("date: "))
        print("---------------------------")
    nDataRider = n
    return nDataEvent, nDataRider

def menuTampilkanDataRider(dataEvent, nDataEvent, dataRider, nDataRider):
    print("Tampilkan Data Rider")
    printEvent(dataEvent, nDataEvent)
    ID = int(input("Pilih idx yang anda pilih: "))
    clear_screen()
    print("name    : ", dataEvent[ID].name)
    print("circuit : ", dataEvent[ID].circuit)
    print("date    : ", dataEvent[ID].date)
    print("===========================")
    printAllRiderByIDEvent(dataRider, nDataRider, ID + 1)

def menuPencarianDataRandom(dataRider, nDataRider):
    print("Kompleksitas Pencarian Data")
    random_idx = random.randint(0, nDataRider - 1)
    timeRAC = dataRider[random_idx].timeRAC
    print("Data yang dicari: ", timeRAC, " pada index ", random_idx)
    print("name    : ", dataRider[random_idx].name)
    print("timeRAC : ", dataRider[random_idx].timeRAC)
    print("Kompleksitas Pencarian Data Iterative")
    print("Data yang dicari: ", binarySearchIterative(dataRider, nDataRider, timeRAC))
    print("Kompleksitas Pencarian Data Recursive")
    print("Data yang dicari: ", binarySearchRecursiveMaster(dataRider, nDataRider, timeRAC))

def menuKompleksitasPencarianData(dataRider, nDataRider, dataTimeIterative, nDataTimeIterative, dataTimeRecursive, nDataTimeRecursive, lap):
    print("Kompleksitas Pencarian Data")
    nDataTimeIterative = 0
    nDataTimeRecursive = 0
    i = 4
    while i < nDataRider:
        n = i + 1
        # Iterative
        start = time.time_ns()
        binarySearchIterative(dataRider, n, dataRider[0].timeRAC)
        awal = time.time_ns() - start
        start = time.time_ns()
        binarySearchIterative(dataRider, n, dataRider[i // 2].timeRAC)
        tengah = time.time_ns() - start
        start = time.time_ns()
        binarySearchIterative(dataRider, n, dataRider[i].timeRAC)
        akhir = time.time_ns() - start
        nDataTimeIterative = addTimeComplexity(dataTimeIterative, nDataTimeIterative, awal, tengah, akhir)
        # Recursive
        start = time.time_ns()
        binarySearchRecursiveMaster(dataRider, n, dataRider[0].timeRAC)
        awal = time.time_ns() - start
        start = time.time_ns()
        binarySearchRecursiveMaster(dataRider, n, dataRider[i // 2].timeRAC)
        tengah = time.time_ns() - start
        start = time.time_ns()
        binarySearchRecursiveMaster(dataRider, n, dataRider[i].timeRAC)
        akhir = time.time_ns() - start
        nDataTimeRecursive = addTimeComplexity(dataTimeRecursive, nDataTimeRecursive, awal, tengah, akhir)
        i += 5
    print("Kompleksitas Pencarian Data Iterative")
    printComplexity(dataTimeIterative, nDataTimeIterative, lap)
    print("Kompleksitas Pencarian Data Recursive")
    printComplexity(dataTimeRecursive, nDataTimeRecursive, lap)
    return nDataTimeIterative, nDataTimeRecursive

def randomInt(min, max):
    return random.randint(min, max)

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
        print("!", dataTime[i].awal // lap, " @", dataTime[i].tengah // lap, " #", dataTime[i].akhir // lap)

def addTimeComplexity(dataTime, nDataTime, awal, tengah, akhir):
    dataTime[nDataTime].awal += awal
    dataTime[nDataTime].tengah += tengah
    dataTime[nDataTime].akhir += akhir
    nDataTime += 1
    return nDataTime  # Ensure nDataTime is updated

def binarySearchIterative(dataRider, nDataRider, timeRAC):
    found = -1
    left = 0
    right = nDataRider - 1
    while left <= right and found == -1:
        # time.sleep(0.001)
        mid = (right + left) // 2
        if dataRider[mid].timeRAC == timeRAC:
            found = mid
        elif dataRider[mid].timeRAC < timeRAC:
            left = mid + 1
        else:
            right = mid - 1
    return found

def binarySearchRecursive(dataRider, timeRAC, mid, left, right, found):
    # time.sleep(0.001)
    if left <= right and found == -1:
        mid = (right + left) // 2
        if dataRider[mid].timeRAC == timeRAC:
            found = mid
        elif dataRider[mid].timeRAC < timeRAC:
            left = mid + 1
            return binarySearchRecursive(dataRider, timeRAC, mid, left, right, found)
        else:
            right = mid - 1
            return binarySearchRecursive(dataRider, timeRAC, mid, left, right, found)
    return found

def binarySearchRecursiveMaster(dataRider, nDataRider, time):
    return binarySearchRecursive(dataRider, time, 0, 0, nDataRider - 1, -1)

def inputFrasa():
    return input()

def clear_screen():
    os.system('cls' if os.name == 'nt' else 'clear')

if __name__ == "__main__":
    main()
