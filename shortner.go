package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type UrlShortener struct {
	LongUrl  string
	ShortUrl string
}

func response_data(o_url string) string {
	response, err := http.Get(o_url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	return string(contents)
}

func url_shortener(o_url string) (string, string) {
	eurl := url.QueryEscape(o_url)
	gd := fmt.Sprintf("http://is.gd/create.php?url=%s&format=simple", eurl)
	return response_data(gd), o_url
}

func (u *UrlShortener) short(o_url string) *UrlShortener {
	shortUrl, longurl := url_shortener(o_url)
	u.ShortUrl = shortUrl
	u.LongUrl = longurl
	return u
}

func main() {
	l_url := UrlShortener{}
	l_url.short("https://www.geeksforgeeks.org/golang-tutorial-learn-go-programming-language/?ref=lbp")
	fmt.Println("Original url is- \n", l_url.LongUrl)
	fmt.Println("Shortened url is- \n", l_url.ShortUrl)
}
