package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var strToFind string = "go"

//----------------------getCount-------------------------------
func getCount(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	kolvo := strings.Count(string(body), strToFind)
	fmt.Println("Count for", url, "-", kolvo)
	return kolvo
}

//---------------------------main--------------------------
func main() {

	totalKolvo := 0

	file, err := os.Open("url.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := getCount(scanner.Text())
		totalKolvo = totalKolvo + v
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total: ", totalKolvo)

}

//---------------------------------------------------------------
