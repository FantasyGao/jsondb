package jsondb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

type base struct {
	table map[string]interface{}
}

var db base
var filePath string
var lock sync.RWMutex

func checkErr(errMsg error) {
	if errMsg != nil {
		panic(errMsg)
		fmt.Println("run error:", errMsg)
	}
}

func isExist(fileName string) bool {
	f, err := os.Stat(fileName)

	if err == nil {
		if !f.IsDir() {
			return true
		}
	}
	return false
}

// init json db
func initialize(fileName string) {

	if path.Ext(fileName) != ".json" {
		filePath = fileName + ".json"
	} else {
		filePath = fileName
	}

	exist := isExist(filePath)

	if !exist {
		f, err := os.Create(filePath)
		defer f.Close()

		checkErr(err)
	}

	db.syncData()
}

// sync file data to memo
func (this *base) syncData() {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
	defer f.Close()

	checkErr(err)

	contentByte, err := ioutil.ReadAll(f)

	checkErr(err)

	if len(contentByte) != 0 {
		err = json.Unmarshal(contentByte, &this.table)

		checkErr(err)
	} else {
		this.table = make(map[string]interface{})
	}
}

// save data to File
func (this *base) Save() *base {

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()

	checkErr(err)

	data, err := json.Marshal(this.table)

	checkErr(err)

	_, err = f.Write(data)

	checkErr(err)

	return this
}

// write data by key value
func (this *base) Write(key string, value interface{}) *base {
	lock.Lock()
	this.table[key] = value
	lock.Unlock()
	return this
}

// read data by key
func (this *base) Read(key string) interface{} {
	return this.table[key]
}

// read all data
func (this *base) ReadAll() map[string]interface{} {
	return this.table
}

// delete key
func (this *base) Del(key string) *base {
	delete(this.table, key)
	return this
}

// create db install
func Create(fileName string) *base {

	initialize(fileName)

	return &db
}
