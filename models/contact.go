package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Contact struct {
	Id        int
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AddContact(data Contact, conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `INSERT INTO contacts (email, phone) VALUES ($1, $2)`, data.Email, data.Phone)
	if err != nil {
		return err
	}
	return nil
}

func GetAllContacts(conn *pgx.Conn) ([]Contact, error) {
	rows, err := conn.Query(context.Background(), `
		SELECT id, email, phone, created_at, updated_at
		FROM contacts
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[Contact])
	if err != nil {
		return nil, err
	}

	return users, nil
}
