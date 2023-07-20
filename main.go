package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// buat satu route mthod get
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// buat satu route mthod get alggiii
	e.GET("/hallo", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, dunia!")
	})


	func getUser(c echo.Context) error {
		// User ID from path `users/:id`
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	}
	e.GET("/luas", func(c echo.Context) error {
		// var (
		// 	sisi int
		// 	//luas int
		// )
		// fmt.Printf("Masukkan sisi : ")
		// fmt.Scan(&sisi)
		//luas = sisi * sisi
		//return c.String(http.StatusOK, "Jadi luasnya adalah :")
		// Get team and member from the query string
		e.GET("/users/:id", getUser)
		team := c.QueryParam("team")
		member := c.QueryParam("member")
		return c.String(http.StatusOK, "team:"+team+", member:"+member)

	})

	// njalanin web server
	e.Logger.Fatal(e.Start(":1323"))
}
