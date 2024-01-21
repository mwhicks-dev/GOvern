package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

type Record struct {
	sid string
	usr string
	pwd string
}

func InitDatabaseTable(username string) (*sql.DB, error) {
	// Configure SQL access
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "govern",
	}

	// Get database handle
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create SQL query
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (id BIGINT AUTO_INCREMENT NOT NULL, record_name VARCHAR(256) NOT NULL, username VARCHAR(256), password VARCHAR(256), PRIMARY KEY(id));", username)

	// Execute SQL query
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddNewRecord(db *sql.DB, username string, record Record) error {
	// Ensure database access still OK
	err := db.Ping()
	if err != nil {
		return err
	}

	/* Check for existing record */
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM %v WHERE record_name = \"%v\";", username, record.sid))
	id := -1
	if err := row.Scan(&id, &record.sid, &record.usr, &record.pwd); err == nil {
		return errors.New(fmt.Sprintf("Cannot add new record %v; already exists", record.sid))
	}

	/* Add new record */
	// Create SQL query
	query := fmt.Sprintf("INSERT INTO %v (record_name, username, password) VALUES (\"%v\", \"%v\", \"%v\");", username, record.sid, record.usr, record.pwd)

	// Execute SQL query
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateExistingRecord(db *sql.DB, username string, record Record) error {
	// Ensure database access still OK
	err := db.Ping()
	if err != nil {
		return err
	}

	// Create SQL query
	query := fmt.Sprintf("UPDATE %v SET username = \"%v\", password = \"%v\" WHERE record_name = \"%v\";", username, record.usr, record.pwd, record.sid)

	// Execute SQL query
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func QueryExistingRecord(db *sql.DB, username string, record_id string) (Record, error) {
	// Initial record setup
	record := Record{"", "", ""}

	// Ensure database access still OK
	err := db.Ping()
	if err != nil {
		return record, err
	}

	// Query row
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM %v WHERE record_name = \"%v\";", username, record_id))
	id := -1
	if err := row.Scan(&id, &record.sid, &record.usr, &record.pwd); err != nil {
		return record, err
	}

	return record, nil
}

func QueryAllRecords(db *sql.DB, username string) ([]Record, error) {
	// Initial records setup
	records := make([]Record, 0)

	// Query rows
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %v;", username))
	if err != nil {
		return nil, err
	}

	// Loop through rows
	id := -1
	for rows.Next() {
		record := Record{"", "", ""}
		if err := rows.Scan(&id, &record.sid, &record.usr, &record.pwd); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}
