#include <iostream>
#include <string>
#include <vector>
#include <cstdlib>
#include <ctime>
#include <algorithm>

using namespace std;

const int NMAXrider = 2000;
const int NMAXevent = 21;
const int NMAXcomplexity = 400; // 2000 / 5 = 400

struct Rider {
    string name;
    int timeRAC, idEvent;
};

struct Event {
    int idEvent;
    string name, circuit, date;
};

struct Complexity {
    int awal, tengah, akhir;
};

void clear_screen() {
    system("cls"); // For Windows
}

void inputFrasa(string &str) {
    getline(cin, str);
}

int randomInt(int min, int max) {
    return min + rand() % (max - min);
}

void printEvent(const vector<Event> &dataEvent) {
    for (const auto &event : dataEvent) {
        cout << "idEvent : " << event.idEvent << endl;
        cout << "name    : " << event.name << endl;
        cout << "circuit : " << event.circuit << endl;
        cout << "date    : " << event.date << endl;
        cout << "---------------------------" << endl;
    }
}

void printAllRiderByIDEvent(const vector<Rider> &dataRider, int ID) {
    for (const auto &rider : dataRider) {
        if (rider.idEvent == ID) {
            cout << "name    : " << rider.name << endl;
            cout << "timeRAC : " << rider.timeRAC << endl;
            cout << "---------------------------" << endl;
        }
    }
}

void addTimeComplexity(vector<Complexity> &dataTime, int awal, int tengah, int akhir) {
    Complexity comp = {awal, tengah, akhir};
    dataTime.push_back(comp);
}

void printComplexity(const vector<Complexity> &dataTime) {
    for (const auto &comp : dataTime) {
        cout << "!" << comp.awal << " @" << comp.tengah << " #" << comp.akhir << endl;
    }
}

int binarySearchIterative(const vector<Rider> &dataRider, int nDataRider, int timeRAC) {
    int left = 0, right = nDataRider - 1;
    while (left <= right) {
        int mid = (left + right) / 2;
        if (dataRider[mid].timeRAC == timeRAC) {
            return mid;
        } else if (dataRider[mid].timeRAC < timeRAC) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return -1;
}

int binarySearchRecursive(const vector<Rider> &dataRider, int timeRAC, int left, int right) {
    if (left <= right) {
        int mid = (left + right) / 2;
        if (dataRider[mid].timeRAC == timeRAC) {
            return mid;
        } else if (dataRider[mid].timeRAC < timeRAC) {
            return binarySearchRecursive(dataRider, timeRAC, mid + 1, right);
        } else {
            return binarySearchRecursive(dataRider, timeRAC, left, mid - 1);
        }
    }
    return -1;
}

int main() {
    srand(static_cast<unsigned int>(time(0))); // Seed for random number generation
    vector<Event> dataEvent(NMAXevent);
    vector<Rider> dataRider(NMAXrider);
    int nDataEvent = 0, nDataRider = 0;
    string pilih;

    while (true) {
        clear_screen();
        cout << "--------------------------" << endl;
        cout << "          M E N U         " << endl;
        cout << "--------------------------" << endl;
        cout << "1. Input Data Rider" << endl;
        cout << "2. Tampilkan Data Rider" << endl;
        cout << "3. Pencarian Data Random" << endl;
        cout << "4. Kompleksitas Pencarian Data" << endl;
        cout << "5. Keluar" << endl;
        cout << "--------------------------" << endl;
        cout << "Pilih (1-5): ";
        inputFrasa(pilih);

        if (pilih == "1") {
            clear_screen();
            cout << "Input Data Rider" << endl;
            cout << "Masukkan data ajang sebanyak 20 + 1 dummy" << endl;
            for (int i = 0; i < 21; i++) {
                cout << "idEvent: ";
                cin >> dataEvent[i].idEvent;
                cin.ignore(); // Clear newline character from the input buffer
                cout << "name: ";
                inputFrasa(dataEvent[i].name);
                cout << "circuit: ";
                inputFrasa(dataEvent[i].circuit);
                cout << "date: ";
                inputFrasa(dataEvent[i].date);
                cout << "---------------------------" << endl;
                nDataEvent++;
            }
            cout << "Masukkan banyak pembalap: ";
            cin >> nDataRider;
            cin.ignore(); // Clear newline character from the input buffer
            for (int i = 0; i < nDataRider; i++) {
                cout << "idEvent: ";
                cin >> dataRider[i].idEvent;
                cout << "name: ";
                cin.ignore(); // Clear newline character from the input buffer
                inputFrasa(dataRider[i].name);
                cout << "date: ";
                cin >> dataRider[i].timeRAC;
                cout << "---------------------------" << endl;
            }
        } else if (pilih == "2") {
            clear_screen();
            cout << "Tampilkan Data Rider" << endl;
            printEvent(dataEvent);
            int ID;
            cout << "Pilih idx yang anda pilih: ";
            cin >> ID;
            clear_screen();
            cout << "name    : " << dataEvent[ID].name << endl;
            cout << "circuit : " << dataEvent[ID].circuit << endl;
            cout << "date    : " << dataEvent[ID].date << endl;
            cout << "===========================" << endl;
            printAllRiderByIDEvent(dataRider, ID + 1);
        } else if (pilih == "3") {
            clear_screen();
            cout << "Kompleksitas Pencarian Data" << endl;
            int random = randomInt(0, nDataRider);
            int timeRAC = dataRider[random].timeRAC;

            cout << "Data yang dicari: " << timeRAC << " pada index " << random << endl;
            cout << "name    : " << dataRider[random].name << endl;
            cout << "timeRAC : " << dataRider[random].timeRAC << endl;
            cout << "Kompleksitas Pencarian Data Iterative" << endl;
            cout << "Data yang dicari: " << binarySearchIterative(dataRider, nDataRider, timeRAC) << endl;
            cout << "Kompleksitas Pencarian Data Recursive" << endl;
            cout << "Data yang dicari: " << binarySearchRecursive(dataRider, timeRAC, 0, nDataRider - 1) << endl;
        } else if (pilih == "4") {
            clear_screen();
            cout << "Kompleksitas Pencarian Data" << endl;
            vector<Complexity> dataTimeIterative, dataTimeRecursive;
            int i = 4;

            while (i < nDataRider) {
                int n = i + 1;
                // Iterative
                auto start = clock();
                binarySearchIterative(dataRider, n, dataRider[0].timeRAC);
                int awal = (clock() - start) * 1000 / CLOCKS_PER_SEC;

                start = clock();
                binarySearchIterative(dataRider, n, dataRider[i / 2].timeRAC);
                int tengah = (clock() - start) * 1000 / CLOCKS_PER_SEC;

                start = clock();
                binarySearchIterative(dataRider, n, dataRider[i].timeRAC);
                int akhir = (clock() - start) * 1000 / CLOCKS_PER_SEC;

                addTimeComplexity(dataTimeIterative, awal, tengah, akhir);

                // Recursive
                start = clock();
                binarySearchRecursive(dataRider, dataRider[0].timeRAC, 0, n - 1);
                awal = (clock() - start) * 1000 / CLOCKS_PER_SEC;

                start = clock();
                binarySearchRecursive(dataRider, dataRider[i / 2].timeRAC, 0, n - 1);
                tengah = (clock() - start) * 1000 / CLOCKS_PER_SEC;

                start = clock();
                binarySearchRecursive(dataRider, dataRider[i].timeRAC, 0, n - 1);
                akhir = (clock() - start) * 1000 / CLOCKS_PER_SEC;

                addTimeComplexity(dataTimeRecursive, awal, tengah, akhir);
                i += 5;
            }

            cout << "Kompleksitas Pencarian Data Iterative" << endl;
            printComplexity(dataTimeIterative);
            cout << "Kompleksitas Pencarian Data Recursive" << endl;
            printComplexity(dataTimeRecursive);
        } else if (pilih == "5") {
            break;
        }
        cout << "Tekan apapun untuk melanjutkan" << endl;
        string bin;
        inputFrasa(bin);
    }
    return 0;
}
