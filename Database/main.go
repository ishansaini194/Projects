package main

import (
	"encoding/json"
	"fmt"

	"github.com/ishansaini194/Projects/db"
	"github.com/ishansaini194/Projects/models"
)

func main() {
	dir := "./data"

	db, err := db.New(dir, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	employees := []models.User{
		{"Ishan", "23", "8968678456", "notGood", models.Address{"Bhatoya", "Punjab", "India", "143534"}},
		{"Tushar", "20", "8968678456", "Gov Hospital", models.Address{"Faridkot", "Punjab", "India", "146001"}},
		{"Kunal", "17", "8968678456", "Gov School", models.Address{"Dinanagar", "Punjab", "India", "143531"}},
		{"Anubhav", "23", "8968678456", "Pvt School", models.Address{"Dinanagar", "Punjab", "India", "143531"}},
		{"Uday", "20", "8968678456", "Pvt Uni", models.Address{"Patiala", "Punjab", "India", "147001"}},
	}

	// write users to DB
	for _, e := range employees {
		if err := db.Write("user", e.Name, e); err != nil {
			fmt.Println("Write error:", err)
		}
	}

	// read all users
	records, err := db.ReadAll("user")
	if err != nil {
		fmt.Println("ReadAll error:", err)
		return
	}

	allUsers := []models.User{}
	for _, r := range records {
		u := models.User{}
		if err := json.Unmarshal([]byte(r), &u); err != nil {
			fmt.Println("Unmarshal error:", err)
		}
		allUsers = append(allUsers, u)
	}

	// pretty print output
	b, err := json.MarshalIndent(allUsers, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(b))

	// delete example
	if err := db.Delete("user", "Uday"); err != nil {
		fmt.Println("Delete error:", err)
	}
}
