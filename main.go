package main

import (
	"context"
	"fmt"
	"koda-b8-db5/lib"
	"koda-b8-db5/models"
	"log"

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

}

func addContact(conn *pgx.Conn) {
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
}

func main() {
	conn := lib.Conn()
	defer conn.Close(context.Background())
	var opt string
	for {

		fmt.Println("====== CONTACT LIST ======")
		fmt.Println("1. Lihat Semua Kontak")
		fmt.Println("2. Tambah Kontak")
		fmt.Print("pilih: ")
		fmt.Scan(&opt)

		switch opt {
		case "1":
			showlistContacts(conn)
		case "2":
			addContact(conn)
		}

	}

}
