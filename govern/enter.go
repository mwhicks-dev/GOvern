package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"govern.hicksm.dev/mysql"
	"govern.hicksm.dev/security"
)

func EnterLoop(reader *bufio.Reader) {

	fmt.Println("Welcome back to GOvern.")
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

	if !mysql.CheckUserTableExists(db, username) {
		fmt.Println("Could not authenticate")
		return
	}

	passwordEncoded := hex.EncodeToString(security.PasswordToKey(password, "master"))

	root, err := mysql.QueryExistingRecord(db, username, "root")
	if err != nil {
		fmt.Println("Could not authenticate.")
		return
	}

	if mysql.GetRecordPwd(root) != passwordEncoded {
		fmt.Println("Could not authenticate.")
		return
	}

}

func AuthenticatedLoop(reader *bufio.Reader, username string, password string) {

	fmt.Println("Authenticated successfully.")

}
