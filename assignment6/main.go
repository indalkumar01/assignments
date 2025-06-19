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

func main() {
	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /create-employee", func(w http.ResponseWriter, r *http.Request) {
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("insert into employees(employee_id, employee_name, employee_designation, employee_salery) values(?, ?, ?, ?)", emp.EmployeeID, emp.EmployeeName, emp.EmployeeDesignation, emp.EmployeeSalary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte("successfully created employee"))
	})

	mux.HandleFunc("PUT /update-employee", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("update employees set employee_name = ?, employee_designation=?, employee_salery=? where employee_id=?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = stmt.Exec(emp.EmployeeName, emp.EmployeeDesignation, emp.EmployeeSalary, emp.EmployeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer stmt.Close()

		w.Write([]byte("updated"))

	})

	mux.HandleFunc("GET /get-employee/{employeeID}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		paths := strings.Split(r.URL.Path, "/")
		employeeID := paths[len(paths)-1]

		row := db.QueryRow("select * from employees where employee_id = $1", employeeID)

		emp := Employee{}

		if err := row.Scan(&emp.EmployeeID, &emp.EmployeeName, &emp.EmployeeDesignation, &emp.EmployeeSalary); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(emp)
	})

	mux.HandleFunc("GET /get-all-employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		employees := []Employee{}

		rows, err := db.Query("select * from employees")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for rows.Next() {
			var emp Employee
			rows.Scan(&emp.EmployeeID, &emp.EmployeeName, &emp.EmployeeDesignation, &emp.EmployeeSalary)
			employees = append(employees, emp)
		}

		json.NewEncoder(w).Encode(employees)
	})

	mux.HandleFunc("DELETE /delete-employee/{employeeID}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		paths := strings.Split(r.URL.Path, "/")
		employeeID := paths[len(paths)-1]

		_, err := db.Exec("DELETE from employees where employee_id = ?", employeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte("deleted!"))
	})

	fmt.Println("server runing on port :8080")
	http.ListenAndServe(":8080", mux)

}
