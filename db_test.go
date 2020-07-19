package jsondb

import (
	"os"
	"testing"
)

func Test_Create(t *testing.T) {
	var fileName = "db.json"

	os.Remove(fileName)

	Create("db.json")

	f, err := os.Stat(fileName)

	if err != nil {
		t.Error("os.Stat return err should be <nil>")
	} else {
		if f.IsDir() {
			t.Error("file should not be dir")
		}
	}
}

func Test_Create2(t *testing.T) {
	var fileName = "/Users/gaotingrong/go/src/code.huayi.com/accountSoft/db.json"

	os.Remove(fileName)

	Create(fileName)

	f, err := os.Stat(fileName)

	if err != nil {
		t.Error("os.Stat return err should be <nil>")
	} else {
		if f.IsDir() {
			t.Error("file should not be dir")
		}
	}
}

func Test_Create_Write(t *testing.T) {
	var fileName = "db.json"

	os.Remove(fileName)

	DB := Create("db.json")

	DB.Write("x", "344")

	if DB.table["x"] != "344" {
		t.Error("DB.table['x'] should return 334")
	}
}

func Test_Create_Write_ReadALL(t *testing.T) {
	var fileName = "db.json"

	os.Remove(fileName)

	DB := Create("db.json")

	DB.Write("x", "344")
	DB.Write("y", "345")

	if DB.table["x"] != "344" {
		t.Error("DB.table['x'] should return 334")
	}

	if DB.table["y"] != "345" {
		t.Error("DB.table['y'] should return 335")
	}
}

func Test_Create_Write_Read(t *testing.T) {

	var x = Create("db.json").Write("x", "344").Save().Read("x")
	if x != "344" {
		t.Error("x should return 334")
	}
}
