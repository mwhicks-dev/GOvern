package mysql

import (
	"fmt"
	"testing"
)

func TestAddAndQueryRecords(t *testing.T) {
	// Sample data
	username := "mwhicks2"
	record := Record{"NCSU", "mwhicks2@ncsu.edu", "password123"}

	// Initialize database
	db, err := InitializeSql()
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = InitDatabaseTable(db, username)
	if err != nil {
		t.Fatalf(err.Error())
	}
	_, err = db.Exec(fmt.Sprintf("DELETE FROM %v WHERE id >= 0;", username))
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Add record to database
	err = AddNewRecord(db, username, record)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Ensure correct behavior
	records, err := QueryAllRecords(db, username)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(records) != 1 {
		t.Fatalf("Expected 1 record, found %v", len(records))
	}
	if records[0].sid != record.sid {
		t.Fatalf("Expected SID %v, found %v", record.sid, records[0].sid)
	}
	if records[0].usr != record.usr {
		t.Fatalf("Expected usr %v, found %v", record.usr, records[0].usr)
	}
	if records[0].pwd != record.pwd {
		t.Fatalf("Expected pwd %v, found %v", record.pwd, records[0].pwd)
	}

	record2 := Record{"Email", "mason@hicksm.dev", "password234"}

	// Add record to database
	err = AddNewRecord(db, username, record2)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Ensure correct behavior
	records, err = QueryAllRecords(db, username)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(records) != 2 {
		t.Fatalf("Expected 2 records, found %v", len(records))
	}
	if records[1].sid != record2.sid {
		t.Fatalf("Expected SID %v, found %v", record2.sid, records[1].sid)
	}
	if records[1].usr != record2.usr {
		t.Fatalf("Expected usr %v, found %v", record2.usr, records[1].usr)
	}
	if records[1].pwd != record2.pwd {
		t.Fatalf("Expected pwd %v, found %v", record2.pwd, records[1].pwd)
	}

	record1, err := QueryExistingRecord(db, username, "NCSU")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if record.sid != record1.sid {
		t.Fatalf("Expected SID %v, found %v", record.sid, record1.sid)
	}
	if record.usr != record1.usr {
		t.Fatalf("Expected usr %v, found %v", record.usr, record1.usr)
	}
	if record.pwd != record1.pwd {
		t.Fatalf("Expected pwd %v, found %v", record.pwd, record1.pwd)
	}

	// Try to add duplicate record
	record3 := Record{"NCSU", "tajdari@ncsu.edu", "password1011"}
	err = AddNewRecord(db, username, record3)
	if err == nil {
		t.Fatalf("Able to add two records with identical name")
	}
}

func TestUpdateExistingRecord(t *testing.T) {
	// Sample data
	username := "mwhicks2"
	record := Record{"NCSU", "mwhicks2@ncsu.edu", "password123"}

	// Initialize database
	db, err := InitializeSql()
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = InitDatabaseTable(db, username)
	if err != nil {
		t.Fatalf(err.Error())
	}
	_, err = db.Exec(fmt.Sprintf("DELETE FROM %v WHERE id >= 0;", username))
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Add record to database
	err = AddNewRecord(db, username, record)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Update existing record by name
	record.pwd = "password789"
	err = UpdateExistingRecord(db, username, record)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Query all records
	records, err := QueryAllRecords(db, username)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Ensure correct behavior
	if len(records) != 1 {
		t.Fatalf("Expected 1 record, found %v", len(records))
	}
	if record.sid != records[0].sid {
		t.Fatalf("Expected SID %v, was actually %v", record.sid, records[0].sid)
	}
	if record.usr != records[0].usr {
		t.Fatalf("Expected usr %v, was actually %v", record.usr, records[0].usr)
	}
	if record.pwd != records[0].pwd {
		t.Fatalf("Expected pwd %v, was actually %v", record.pwd, records[0].pwd)
	}
}
