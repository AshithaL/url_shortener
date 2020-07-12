package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Redirect struct {
	Id 		int
	Slug 	string 	`db:"slug" form:"slug"`
	Url  	string	`db:"url" form:"url"`
}

var db, err = sql.Open("mysql", "root:nineleaps@tcp(127.0.0.1:3306)/golang")
var url_Orig = "https://www.geeksforgeeks.org/golang-tutorial-learn-go-programming-language/?ref=lbp"

func main() {
	var redirect Redirect
	row := db.QueryRow("select id, slug, url from redirect where url = ?;", url_Orig)
	err = row.Scan(&redirect.Id, &redirect.Slug, &redirect.Url)
	print(redirect.Slug)

	if err != nil {

		urlOrig := Redirect{}

		Slug := urlOrig.Slug
		Url := urlOrig.Url

		stmt, err := db.Prepare("insert into url_shortener (slug, url) values(?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}

		_, err = stmt.Exec(Slug, Url)
		if err != nil {
			fmt.Print(err.Error())
		}

		defer stmt.Close()
	}
}