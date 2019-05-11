package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

type Company struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func main() {
	path := "./profile.out"
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)

	}
	defer pprof.StopCPUProfile()

	usersRaw, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal("can't read file")
	}

	companiesRaw, err := ioutil.ReadFile("./companies.json")
	if err != nil {
		log.Fatal("can't read file")
	}

	users := make([]User, 0)
	if err := json.Unmarshal(usersRaw, &users); err != nil {
		log.Fatal("can't unmarshal data")
	}

	companies := make([]Company, 0)
	if err := json.Unmarshal(companiesRaw, &companies); err != nil {
		log.Fatal("can't unmarshal data")
	}

	users = merge(users, companies)
	fmt.Print(users)
}

func merge(users []User, companies []Company) []User {
	for userKey, user := range users {
		for _, company := range companies {
			if user.ID == company.UserID {
				users[userKey].Company = company.Name
				break
			}
		}
	}
	return users
}

func merge2(users []User, companies []Company) []User {
	companiesByUser := convertByUser(companies)
	for userKey, user := range users {
		company, ok := companiesByUser[user.ID]
		if !ok {
			continue
		}
		users[userKey].Company = company.Name
	}

	return users
}

func convertByUser(companies []Company) map[int]Company {
	companiesByUser := make(map[int]Company, len(companies))
	for _, company := range companies {
		companiesByUser[company.UserID] = company
	}
	return companiesByUser
}
