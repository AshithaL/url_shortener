package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main() {

	dab, err := sql.Open("mysql", "root:nineleaps@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}

	_,err = dab.Exec("CREATE DATABASE IF NOT EXISTS golang")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}

	_,err = dab.Exec("USE golang")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	stmt, err := dab.Prepare("CREATE TABLE IF NOT EXISTS url_shortener ( id int NOT NULL AUTO_INCREMENT, slug varchar(500) collate utf8mb4_unicode_ci NOT NULL, url varchar(620) collate utf8mb4_unicode_ci NOT NULL, PRIMARY KEY (id) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='URL shortener Table';")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}

	defer dab.Close()
}