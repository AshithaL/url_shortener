package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type UrlShortener struct {
	OriginalUrl string
	ShortUrl    string
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
	shortUrl, originalUrl := url_shortener(o_url)
	u.ShortUrl = shortUrl
	u.OriginalUrl = originalUrl
	return u
}

func main() {
	urlOrig := UrlShortener{}
	urlOrig.short("https://www.geeksforgeeks.org/golang-tutorial-learn-go-programming-language/?ref=lbp")
	fmt.Println("Original url is- \n",urlOrig.OriginalUrl)
	fmt.Println("Shortened url is- \n",urlOrig.ShortUrl)
}
