package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
	// "math/rand"
)

const NMAXrider int = 2000
const NMAXevent int = 21
const NMAXcomplexity int = 400 // 2000 / 5 = 400

type rider struct {
	name string
	timeRAC, idEvent int
}
type tabRider [NMAXrider]rider
type event struct {
	idEvent int
	name, circuit, date string
}
type tabEvent [NMAXevent]event
type complexity struct {
	awal, tengah, akhir int
}
type tabComplexity [NMAXcomplexity]complexity

/////////////////////////// Main Function ///////////////////////////
func main() {
	var bin string
	var dataEvent tabEvent
	var dataRider tabRider
	var nDataEvent int
	var nDataRider int
	var pilih string

	nDataRider = 0
	nDataEvent = 0
	for {
		clear_screen()
		fmt.Println("--------------------------")
		fmt.Println("          M E N U         ")
		fmt.Println("--------------------------")
		fmt.Println("1. Input Data Rider")
		fmt.Println("2. Tampilkan Data Rider")
		fmt.Println("3. Pencarian Data Random")
		fmt.Println("4. Kompleksitas Pencarian Data")
		fmt.Println("5. Keluar")
		fmt.Println("--------------------------")
		fmt.Print("Pilih (1-4): ")
		inputFrasa(&pilih)
		switch pilih {
			case "1": 
				clear_screen()	
				menuInputDataRider(&dataEvent, &nDataEvent, &dataRider, &nDataRider)
			case "2": 
				clear_screen()
				menuTampilkanDataRider(dataEvent, nDataEvent, dataRider, nDataRider)
			case "3": 
				clear_screen()
				menuPencarianDataRandom(dataRider, nDataRider)
			case "4": 
				clear_screen()
				menuKompleksitasPencarianData(dataRider, nDataRider)
		} 
		fmt.Print("Tekan apapun untuk melanjutkan")
		inputFrasa(&bin)
		clear_screen()
		if pilih == "5" {
			break
		}
	}
}

/////////////////////////// Main Menu ///////////////////////////
// menu 1
func menuInputDataRider(dataEvent *tabEvent, nDataEvent *int, dataRider *tabRider, nDataRider *int) {
	var i, n int
	
	fmt.Println("Input Data Rider")
	fmt.Println("Masukkan data ajang sebanyak 20 + 1 dummy")
	for i = 0; i < 21; i++ {
		fmt.Print("idEvent: ")
		fmt.Scan(&dataEvent[i].idEvent)
		fmt.Print("name: ")
		inputFrasa(&dataEvent[i].name)
		fmt.Print("circuit: ")
		inputFrasa(&dataEvent[i].circuit)
		fmt.Print("date: ")
		inputFrasa(&dataEvent[i].date)
		fmt.Println("---------------------------")
		*nDataEvent++;
	}
	fmt.Print("Masukkan banyak pembalap: ")
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Print("idEvent: ")
		fmt.Scan(&dataRider[i].idEvent)
		fmt.Print("name: ")
		inputFrasa(&dataRider[i].name)
		fmt.Print("date: ")
		fmt.Scan(&dataRider[i].timeRAC)
		fmt.Println("---------------------------")
		*nDataRider = n
	}
}
// menu 2
func menuTampilkanDataRider(dataEvent tabEvent, nDataEvent int, dataRider tabRider, nDataRider int) {
	var ID int

	fmt.Println("Tampilkan Data Rider")
	printEvent(dataEvent, nDataEvent);
	fmt.Print("Pilih idx yang anda pilih: ")
	fmt.Scan(&ID)
	clear_screen()
	fmt.Println("name    : ", dataEvent[ID].name)
	fmt.Println("circuit : ", dataEvent[ID].circuit)
	fmt.Println("date    : ", dataEvent[ID].date)
	fmt.Println("===========================")
	printAllRiderByIDEvent(dataRider, nDataRider, ID + 1)
}
// menu 3
func menuPencarianDataRandom(dataRider tabRider, nDataRider int) {
	var timeRAC, random int

	fmt.Println("Kompleksitas Pencarian Data")
	random = randomInt(0, nDataRider - 1)
	timeRAC = dataRider[random].timeRAC
	
	fmt.Println("Data yang dicari: ", timeRAC, " pada index ", random)
	fmt.Println("name    : ", dataRider[random].name)
	fmt.Println("timeRAC : ", dataRider[random].timeRAC)
	fmt.Println("Kompleksitas Pencarian Data Iterative")
	fmt.Println("Data yang dicari: ", binarySearchIterative(dataRider, nDataRider, timeRAC))
	fmt.Println("Kompleksitas Pencarian Data Recursive")
	fmt.Println("Data yang dicari: ", binarySearchRecursiveMaster(dataRider, nDataRider, timeRAC))
}
// menu 4
func menuKompleksitasPencarianData(dataRider tabRider, nDataRider int) {
	var dataTimeIterative, dataTimeRecursive tabComplexity
	var nDataTimeIterative, nDataTimeRecursive int
	var i, n int
	var start time.Time
	var awal, tengah, akhir int
	var bin string

	fmt.Println("Kompleksitas Pencarian Data")
	nDataTimeIterative = 0
	nDataTimeRecursive = 0
	i = 4
	for i < nDataRider {
		n = i + 1
		// fmt.Println("n =", n, "| dataRider[0].timeRAC =", dataRider[0].timeRAC, "| dataRider[i/2].timeRAC =", dataRider[i/2].timeRAC, "| dataRider[i].timeRAC =", dataRider[i].timeRAC)
		// Iterative //
		// awal
		start = time.Now()
		binarySearchIterative(dataRider, n, dataRider[0].timeRAC)
		awal = int(time.Duration(time.Since(start).Nanoseconds()))
		// tengah
		start = time.Now()
		binarySearchIterative(dataRider, n, dataRider[i/2].timeRAC)
		tengah = int(time.Duration(time.Since(start).Nanoseconds()))
		// akhir
		start = time.Now()
		binarySearchIterative(dataRider, n, dataRider[i].timeRAC)
		akhir = int(time.Duration(time.Since(start).Nanoseconds()))
		addTimeComplexity(&dataTimeIterative, &nDataTimeIterative, awal, tengah, akhir)
		// Recursive //
		// awal
		start = time.Now()
		binarySearchRecursiveMaster(dataRider, n, dataRider[0].timeRAC)
		awal = int(time.Duration(time.Since(start).Nanoseconds()))
		// tengah
		start = time.Now()
		binarySearchRecursiveMaster(dataRider, n, dataRider[i/2].timeRAC)
		tengah = int(time.Duration(time.Since(start).Nanoseconds()))
		// akhir
		start = time.Now()
		binarySearchRecursiveMaster(dataRider, n, dataRider[i].timeRAC)
		akhir = int(time.Duration(time.Since(start).Nanoseconds()))
		addTimeComplexity(&dataTimeRecursive, &nDataTimeRecursive, awal, tengah, akhir)
		i += 5
	}
	fmt.Scan(&bin)
	fmt.Println("Kompleksitas Pencarian Data Iterative")
	printComplexity(dataTimeIterative, nDataTimeIterative)
	fmt.Scan(&bin)
	fmt.Println("Kompleksitas Pencarian Data Recursive")
	printComplexity(dataTimeRecursive, nDataTimeRecursive)
}

/////////////////////////// func ///////////////////////////
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func printEvent(dataEvent tabEvent, nDataEvent int) {
	var i int

	for i = 0; i < nDataEvent; i++ {
		fmt.Println("idEvent : ", dataEvent[i].idEvent)
		fmt.Println("name    : ", dataEvent[i].name)
		fmt.Println("circuit : ", dataEvent[i].circuit)
		fmt.Println("date    : ", dataEvent[i].date)
		fmt.Println("---------------------------")
	}
}

func printAllRiderByIDEvent(dataRider tabRider, nDataRider, ID int) {
	var i int

	for i = 0; i < nDataRider; i++ {
		if dataRider[i].idEvent == ID {
			fmt.Println("name    : ", dataRider[i].name)
			fmt.Println("timeRAC : ", dataRider[i].timeRAC)
			fmt.Println("---------------------------")
		}
	}
}

func printComplexity(dataTime tabComplexity, nDataTime int) {
	var i int

	for i = 0; i < nDataTime; i++ {
		fmt.Print("!", dataTime[i].awal, " @", dataTime[i].tengah, " #", dataTime[i].akhir, "\n")
	}
}

func addTimeComplexity(dataTime *tabComplexity, nDataTime *int, awal, tengah, akhir int) {
	dataTime[*nDataTime].awal = awal
	dataTime[*nDataTime].tengah = tengah
	dataTime[*nDataTime].akhir = akhir
	*nDataTime++
}

/////////////////////////// Algo ///////////////////////////
func binarySearchIterative(dataRider tabRider, nDataRider, timeRAC int) int {
// mengaembalikan index pencarian apabila x ada di dalam array yang berisi n
// bilangan atau -1 apabila tidak ditemukan, tab terurut membesar atau ascending,
// algotitma dengan Iterative
	var mid, left, right, found int

	found = -1
	left = 0
	right = nDataRider - 1
	for left <= right && found == -1 {
		// time.Sleep(1)
		mid = (right + left) / 2
		if dataRider[mid].timeRAC == timeRAC {
			found = mid
		} else if dataRider[mid].timeRAC < timeRAC {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}

func binarySearchRecursive(dataRider tabRider, timeRAC, mid, left, right, found int) int {
// mengaembalikan index pencarian apabila x ada di dalam array yang berisi n
// bilangan atau -1 apabila tidak ditemukan, tab terurut membesar atau ascending,
// algotitma dengan Recursive
	// time.Sleep(1)
	if left <= right && found == -1 {
		mid = (right + left) / 2
		if dataRider[mid].timeRAC == timeRAC {
			found = mid
		} else if dataRider[mid].timeRAC < timeRAC {
			left = mid + 1
			return binarySearchRecursive(dataRider, timeRAC, mid, left, right, found)
		} else {
			right = mid - 1
			return binarySearchRecursive(dataRider, timeRAC, mid, left, right, found)
		}
	}
	return found
}
func binarySearchRecursiveMaster(dataRider tabRider, nDataRider, time int) int {
	return binarySearchRecursive(dataRider, time, 0, 0, nDataRider -1, -1)
}

/////////////////////////// Hiasan ///////////////////////////
func inputFrasa(str *string) {
	/*
		IS: terdefinisi sembarang str (string)
		FS: str terisi frasa dan berhanti ketika masukan berupa enter
	*/
	var ch, i byte
		
	*str = ""
	for i = 1; i <= 2; i++ {
		fmt.Scanf("%c", &ch)
		if ch != '\r'&& ch != '\n' {
			*str += string(ch)
		}
	}
	for {
		fmt.Scanf("%c", &ch)
		if ch != '\r' && ch != '\n' {
			*str += string(ch)
		}
		if ch == '\n' {
			break
		}
	}
}

func clear_screen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
} 