package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"govern.hicksm.dev/mysql"
	"govern.hicksm.dev/security"
)

func SignupLoop(reader *bufio.Reader) {

	fmt.Println("Welcome to GOvern.")
	fmt.Print("Enter a username: ")

	text, _ := reader.ReadString('\n')
	username := Clean(text)

	if len(username) == 0 {
		fmt.Println("Username may not be blank.")
		return
	}
	if len(username) > 64 {
		fmt.Println("Username must be 64 characters or less long.")
		return
	}
	if !Alphanumeric(username) {
		fmt.Println("Username must be alphanumeric.")
		return
	}

	db, err := mysql.InitializeSql()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if mysql.CheckUserTableExists(db, username) {
		fmt.Printf("Username %v already exists.\n", username)
		return
	}

	fmt.Print("Enter a password: ")

	text, _ = reader.ReadString('\n')
	password := Clean(text)

	if len(password) == 0 {
		fmt.Println("Password may not be blank.")
		return
	}
	if len(password) > 64 {
		fmt.Println("Password must be 64 characters or less long.")
		return
	}

	fmt.Print("Enter again: ")

	text, _ = reader.ReadString('\n')
	passwordRepeat := Clean(text)

	if password != passwordRepeat {
		fmt.Println("Passwords don't match.")
		return
	}
	passwordEncoded := hex.EncodeToString(security.PasswordToKey(password, "master"))

	err = mysql.InitDatabaseTable(db, username)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	initialEntry := mysql.CreateRecord("root", username, passwordEncoded)

	err = mysql.AddNewRecord(db, username, initialEntry)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("New account set up.")
}
