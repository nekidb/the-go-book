package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buffer := bytes.Buffer{}

	for item, price := range db {
		buffer.WriteString(fmt.Sprintf("%s: %s\n", item, price))
	}

	w.Write(buffer.Bytes())
}

func main() {
	db := database{"tea": 1, "coffee": 1.5}

	log.Println("Starting server on localhost:8081")
	err := http.ListenAndServe("localhost:8081", db)
	log.Fatal(err)
}
