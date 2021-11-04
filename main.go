package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var count int
	fmt.Print("Init :")
	fmt.Scan(&count)

	arr := make([]string, count)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(arr); i++ {
		fmt.Print("Input ")
		t, _ := reader.ReadString('\n')
		t = strings.TrimSuffix(t, "\n")
		arr[i] = t
		fmt.Println("Kartu identitas tersimpan di loker nomor ", i+1)
	}

	pilih := true
	for pilih {
		var pilihan string
		fmt.Print("pilih pilihan : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case "status":
			fmt.Printf("%0s %10s %10s", "No Loker", "Tipe Identitas", "No Identitas \n")
			for i := 0; i < len(arr); i++ {
				if arr[i] == "" {
					continue
				}
				fmt.Printf("%0d %10s %10s \n", i+1, strings.Split(arr[i], " ")[0], strings.Split(arr[i], " ")[1])
			}
		case "tambah":
			fmt.Print("Input ")
			t, _ := reader.ReadString('\n')
			ada := true
			var index int
			for i := 0; i < len(arr); i++ {
				if arr[i] == "" {
					arr[i] = t
					ada = false
					index = i
					break
				}
			}
			if ada {
				fmt.Println("Loker masih penuh")
			} else {
				fmt.Println("Identitas anda terletak pada loker ", index+1)
			}
		case "leave":
			fmt.Print("Leave ")
			t, _ := reader.ReadString('\n')
			index, _ := strconv.Atoi(t)
			arr[index+2] = ""
			fmt.Println("Loker sudah dikosongkan")
		case "findbyID":
			var search string
			fmt.Print("find ")
			fmt.Scan(&search)
			index := -1
			for i := 0; i < len(arr); i++ {
				if strings.Contains(arr[i], search) {

					index = i + 1
					break
				}
			}
			if index != -1 {
				fmt.Println("Kartu IDentitas berada di loker ", index)
			} else {
				fmt.Println("Kartu tidak ditemukan")
			}
		case "findByIdentitas":
			var identitas string
			fmt.Print("IDentitas ")
			fmt.Scan(&identitas)
			for i := 0; i < len(arr); i++ {
				if strings.ToLower(strings.Split(arr[i], " ")[0]) == strings.ToLower(identitas) {
					fmt.Print(strings.Split(arr[i], " ")[1], ",")
				}

			}
			fmt.Println()
		case "keluar":
			pilih = false
		}

	}
}
