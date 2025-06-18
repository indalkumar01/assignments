package main

import (
	"assignment6/db"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Employee struct {
	EmployeeID          int     `json:"employee_id"`
	EmployeeName        string  `json:"employee_name"`
	EmployeeDesignation string  `json:"employee_designation"`
	EmployeeSalary      float64 `json:"employee_salery"`
}

func (e *Employee) Create() {

}

func (e *Employee) Update() {

}

func (e *Employee) GetOne(employeeID int) {

}

func (e *Employee) GetAll() {

}

func (e *Employee) Delete(employeeID int) {

}

func main() {
	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("db:===>", db)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /create-employee", func(w http.ResponseWriter, r *http.Request) {
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(emp)

		w.Write([]byte("successfully created employee"))
	})

	mux.HandleFunc("PUT /update-employee", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		emp := new(Employee)
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte("updated"))

	})

	mux.HandleFunc("GET /get-employee/{employeeID}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		paths := strings.Split(r.URL.Path, "/")
		employeeID := paths[len(paths)-1]

		json.NewEncoder(w).Encode(map[string]interface{}{"msg": "success", "employeeID": employeeID})
	})

	mux.HandleFunc("GET /get-all-employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Employee{
			{EmployeeID: 1, EmployeeName: "test", EmployeeDesignation: "test", EmployeeSalary: 100},
			{EmployeeID: 1, EmployeeName: "test2", EmployeeDesignation: "test2", EmployeeSalary: 200},
		})
	})

	mux.HandleFunc("DELETE /delete-employee/{employeeID}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		w.Write([]byte("deleted!"))
	})

	fmt.Println("server runing on port :8080")
	http.ListenAndServe(":8080", mux)

}
