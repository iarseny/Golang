package main

import ("fmt"
	"net/http"
    "io/ioutil"
    "log")

func main() {
	var res, err = http.Get("https://www.binance.com/ru");
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println(res.Header, err);
}