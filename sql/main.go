package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type People struct {
	id    int
	name  string
	age   int
	email string
}

type Company struct {
	id         int
	people_id  int
	name       string
	startedat  sql.NullString
	finishedat sql.NullString
}

//SavePeople saves the people list from db in a slice
func SavePeople(db *sql.DB) []People {

	var people People
	rows, err := db.Query(`SELECT id, name, age, email FROM people;`)
	if err != nil {
		panic(err)
	}
	var peopleSlice []People

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&people.id, &people.name, &people.age, &people.email)
		if err != nil {
			panic(err)
		}
		peopleSlice = append(peopleSlice, people)
	}
	return peopleSlice
}

//SaveCompany saves the company list from db in a slice
func SaveCompany(db *sql.DB) []Company {
	var company Company
	rows, err := db.Query(`SELECT id, people_id, name, startedat, finishedat FROM company;`)
	if err != nil {
		panic(err)
	}
	var companySlice []Company

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&company.id, &company.people_id, &company.name, &company.startedat, &company.finishedat)
		if err != nil {
			panic(err)
		}
		companySlice = append(companySlice, company)
	}
	return companySlice
}

func main() {
	connStr := "user=postgres dbname=task2 sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(SavePeople(db))
	fmt.Println(SaveCompany(db))
}
