package main

import (
	"context"
	"fmt"
	"koda-b8-db5/lib"
	"koda-b8-db5/models"
	"koda-b8-db5/utils"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func showlistContacts(conn *pgx.Conn) {
	listContacts, err := models.GetAllContacts(conn)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range listContacts {
		fmt.Printf("ID    : %d\n", c.Id)
		fmt.Printf("Email : %s\n", c.Email)
		fmt.Printf("Phone : %s\n", c.Phone)
		fmt.Println("-----------------------")
	}
	fmt.Print("tekan enter untuk kembali....")
	fmt.Scanln()
}

func UpdateContact(conn *pgx.Conn) {
	fmt.Println("Update Contact")
	var contact models.Contact
	fmt.Print("masukkan Id: ")
	fmt.Scan(&contact.Id)
	fmt.Print("masukkan email baru: ")
	fmt.Scan(&contact.Email)
	fmt.Print("masukkan no Hp Baru: ")
	fmt.Scan(&contact.Phone)
	err := models.UpdateContact(contact, conn)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("contact berhasil di update")

}

func deleteContact(conn *pgx.Conn) {
	var id int
	fmt.Println("Delete Contact")
	fmt.Print("input Id yang ingin dihapus: ")
	fmt.Scan(&id)

	err := models.DeleteContact(id, conn)
	if err != nil {
		fmt.Println("gagal menghapus contact")
		return
	}
	fmt.Println("contact berhasil dihapus")
	fmt.Scanln()
}

func addContact(conn *pgx.Conn) {
	fmt.Println("Add Contact")
	var contact models.Contact

	fmt.Print("Email : ")
	fmt.Scan(&contact.Email)

	fmt.Print("Phone : ")
	fmt.Scan(&contact.Phone)

	err := models.AddContact(contact, conn)
	if err != nil {
		log.Println("Gagal menambahkan contact:", err)
		return
	}

	fmt.Println("Contact berhasil ditambahkan!")
	fmt.Println("tekan enter untuk kembali....")
	fmt.Scanln()

}

func main() {

	conn := lib.Conn()
	defer conn.Close(context.Background())
	var opt string
	for {
		utils.Clear()
		fmt.Println("====== CONTACT LIST ======")
		fmt.Println("1. List Contact")
		fmt.Println("2. Add Contact")
		fmt.Println("3. Edit Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("--------------------------")
		fmt.Println("0. Keluar")
		fmt.Print("pilih: ")
		fmt.Scan(&opt)

		switch opt {
		case "1":
			showlistContacts(conn)
		case "2":
			addContact(conn)
		case "3":
			UpdateContact(conn)
		case "4":
			deleteContact(conn)
		case "0":
			os.Exit(0)
		}

	}

}
