package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	// "math/rand"
)

const NMAX int = 400

type rider struct {
	name, nat, team string
	timeRAC int
	idEvent int
}
type tabRider [NMAX]rider
type event struct {
	idEvent int
	name, circuit, date string
}
type tabEvent [NMAX]event
type tabTime [NMAX]int

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
		fmt.Println("3. Kompleksitas Pencarian Data")
		fmt.Println("4. Keluar")
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
				menuKompleksitasPencarianData(dataRider, nDataRider)
		} 
		fmt.Print("Tekan apapun untuk melanjutkan")
		inputFrasa(&bin)
		clear_screen()
		if pilih == "4" {
			break
		}
	}
}

/////////////////////////// Main Menu ///////////////////////////
func menuInputDataRider(dataEvent *tabEvent, nDataEvent *int, dataRider *tabRider, nDataRider *int) {
	var i, n int
	
	fmt.Println("Input Data Rider")
	fmt.Println("Masukkan data ajang sebanyak 20")
	for i = 0; i < 20; i++ {
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
		fmt.Print("nat: ")
		inputFrasa(&dataRider[i].nat)
		fmt.Print("team: ")
		inputFrasa(&dataRider[i].team)
		fmt.Print("date: ")
		fmt.Scan(&dataRider[i].timeRAC)
		fmt.Println("---------------------------")
		*nDataRider = n
	}
}
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
	printAllRiderByIDEvent(dataRider, nDataRider, ID)
}
func menuKompleksitasPencarianData(dataRider tabRider, nDataRider int) {
	var dataTimeIterative, dataTimeRecursive tabTime
	var nDataTimeIterative, nDataTimeRecursive int
	var i, timeRAC int
	// var bestCase, worstCase, averageCase int
	var start time.Time
	var elapsed time.Duration

	fmt.Println("Kompleksitas Pencarian Data")
	nDataTimeIterative = 0
	nDataTimeRecursive = 0
	for i = 0; i < nDataRider; i++ {
		timeRAC = dataRider[i].timeRAC
		// Iterative
		start = time.Now()
		binarySearchIterative(dataRider, nDataRider, timeRAC)
		elapsed = time.Since(start)
		addTimeComplexity(&dataTimeIterative, &nDataTimeIterative, int(elapsed.Nanoseconds()))
		// Recursive
		start = time.Now()
		binarySearchRecursiveMaster(dataRider, nDataRider, timeRAC)
		elapsed = time.Since(start)
		addTimeComplexity(&dataTimeRecursive, &nDataTimeRecursive, int(elapsed.Nanoseconds()))
	}
	fmt.Println("Kompleksitas Pencarian Data Iterative")
	printComplexity(dataTimeIterative, nDataTimeIterative)
	fmt.Println("Kompleksitas Pencarian Data Recursive")
	printComplexity(dataTimeRecursive, nDataTimeRecursive)

}

/////////////////////////// func ///////////////////////////
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
			fmt.Println("nat     : ", dataRider[i].nat)
			fmt.Println("team    : ", dataRider[i].team)
			fmt.Println("timeRAC : ", dataRider[i].timeRAC)
			fmt.Println("---------------------------")
		}
	}
}

func printComplexity(dataTime tabTime, nDataTime int) {
	var i int

	for i = 0; i < nDataTime; i++ {
		fmt.Println(dataTime[i])
	}
}

func binarySearchIterative(dataRider tabRider, nDataRider, timeRAC int) int {
// mengaembalikan index pencarian apabila x ada di dalam array yang berisi n
// bilangan atau -1 apabila tidak ditemukan, tab terurut membesar atau ascending,
// algotitma dengan Iterative
	var mid, left, right, found int

	found = -1
	left = 0
	right = nDataRider - 1
	for left <= right && found == -1 {
		time.Sleep(10)
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
	time.Sleep(10)
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

func addTimeComplexity(dataTime *tabTime, nDataTime *int, elapsed int) {
	dataTime[*nDataTime] = elapsed
	*nDataTime++
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