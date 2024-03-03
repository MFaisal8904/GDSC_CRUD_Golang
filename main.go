package main

import (
	"GOLANG_DATABASE/connection"
	"GOLANG_DATABASE/controller"
	"fmt"
)

func main() {
	db := connection.GetConnection()
	if db == nil {
		fmt.Print("ERROR CONECTION")
	}
	var choice int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Insert Mahasiswa")
		fmt.Println("2. List Mahasiswa")
		fmt.Println("3. Select Mahasiswa by NIM")
		fmt.Println("4. Delete Mahasiswa by NIM")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			insertMahasiswa()
		case 2:
			listMahasiswa()
		case 3:
			selectMahasiswaByNIM()
		case 4:
			deleteMahasiswaByNIM()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func insertMahasiswa() {
	var newMahasiswa controller.Mahasiswa
	fmt.Println("\nInsert Mahasiswa:")
	fmt.Print("NIM: ")
	fmt.Scanln(&newMahasiswa.NIM)
	fmt.Print("Nama: ")
	fmt.Scanln(&newMahasiswa.Nama)
	fmt.Print("Kelas: ")
	fmt.Scanln(&newMahasiswa.Kelas)
	fmt.Print("Email")
	fmt.Scanln(&newMahasiswa.Email)

	err := controller.InsertUser(newMahasiswa)
	if err != nil {
		fmt.Println("Error inserting user: ", err)
	} else {
		fmt.Println("Success adding new user.")
	}
}

func listMahasiswa() {
	fmt.Println("\nList Mahasiswa:")
	err := controller.GetUser()
	if err != nil {
		fmt.Println("Error selecting Mahasiswa:", err)
	}
}

func selectMahasiswaByNIM() {
	var nim int
	fmt.Print("\nInput NIM to select Mahasiswa: ")
	fmt.Scanln(&nim)
	err := controller.SelectByNIM(nim)
	if err != nil {
		fmt.Println("Error selecting Mahasiswa by NIM:", err)
	}
}

func deleteMahasiswaByNIM() {
	var deleteNIM int
	fmt.Print("\nInput NIM to delete Mahasiswa: ")
	fmt.Scanln(&deleteNIM)
	err := controller.DeleteUser(deleteNIM)
	if err != nil {
		fmt.Println("Error deleting Mahasiswa by NIM:", err)
	} else {
		fmt.Println("Mahasiswa deleted successfully.")
	}
}
