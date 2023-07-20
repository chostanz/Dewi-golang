package main

import (
	"fmt"
	
	"database/sql"
	_ "github.com/lib/pq"
)


// User
//gaboleh declare di main
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
  }




func main() {

	//membuat koneksi database, mengambil objek di db untuk menjalankan berbagai method query
	 db, err :=  sql.Open("postgres", "user=postgres password=12345 dbname= inventory sslmode=disable" )
	
	if err  = db.Ping(); err != nil {
		fmt.Println(err)
	} 

	//menjalankan query di simpan di var rows
	//parameter dan argumen tu sama, ada di dalam function
	rows, err := db.Query("select * from produk")
		 
	for rows.Next(){
		fmt.Print("Data muncul")
	}

	//  e := echo.New()



	// u := &User{
	// 	Name:  "Jon",
	// 	Email: "jon@labstack.com",
	// }


	// e.GET("/", func(c echo.Context) error {
    //     return c.JSON(http.StatusOK, u)
	// 	//sqlStatement := "SELECT * FROM inventory"
    // })

	// e.Logger.Fatal(e.Start(":8080"))
}

