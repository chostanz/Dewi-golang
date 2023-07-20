package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	db, err := sql.Open("postgres", "user=postgres password=12345 dbname=buku sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	e := echo.New()

	type Employee struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Salary string `json: "salary"`
		Age    string `json : "age"`
	}
	type Employees struct {
		Employee
	}
	e.POST("/employee", func(c echo.Context) error {
		u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "INSERT INTO karyawan (name, salry,age)VALUES ($1, 2, 3)"
		res, err := db.Query(sqlStatement, u.Name, u.Salary, u.Age)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/employee", func(c echo.Context) error {
		sqlStatement := "SELECT id, name, salary, age FROM karyawan order by id"
		rows, err := db.Query(sqlStatement)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		for rows.Next() {
			employee := Employee{}
			rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)

			fmt.Println(employee)
		}
		return c.String(http.StatusCreated, "Data diambil")

		//return c.String(http.StatusOK, "ok")
	})

}
