package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

//ID,Email Address,Name
type user struct {
	Email string
	Name  string
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5000; i++ {
		go read()
	}

	if _, err := fmt.Fprintf(w, "Done!"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/read/", readHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func read() {
	path := "./user.csv"
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println(err)
		return
	}

	for _, row := range rows {
		user := user{
			Name:  row[2],
			Email: row[1],
		}
		save(user)
	}
}

func save(u user) {
	// It is fake)
}

func read2() {
	path := "./user.csv"
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(bufio.NewReader(f))
	reader.ReuseRecord = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		user := user{
			Name:  record[2],
			Email: record[1],
		}
		save(user)
	}
}
