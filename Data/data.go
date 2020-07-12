package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Redirect struct {
	Id 		int
	Slug 	string 	`db:"slug" form:"slug"`
	Url  	string	`db:"url" form:"url"`
}

var db, err = sql.Open("mysql", "root:nineleaps@tcp(127.0.0.1:3306)/golang")
var url_Orig = "https://www.geeksforgeeks.org/golang-tutorial-learn-go-programming-language/?ref=lbp"

func getData(Url string) string {
	response, err := http.Get(Url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	return string(contents)
}

func Url_Shortener(urlOrig string) (string, string) {
	eurl := url.QueryEscape(urlOrig)
	gd := fmt.Sprintf("http://is.gd/create.php?url=%s&format=simple", eurl)
	return getData(gd), urlOrig
}

func (u *Redirect) short(urlOrig string) *Redirect {
	shortUrl, originalUrl := Url_Shortener(urlOrig)
	u.Slug = shortUrl
	u.Url = originalUrl
	return u
}

func main() {
	var redirect Redirect
	row := db.QueryRow("select id, slug, url from redirect where url = ?;", url_Orig)
	err = row.Scan(&redirect.Id, &redirect.Slug, &redirect.Url)
	print(redirect.Slug)

	if err != nil {

		urlOrig := Redirect{}
		urlOrig.short(url_Orig)
		fmt.Println(urlOrig.Slug)
		fmt.Println(urlOrig.Url)

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