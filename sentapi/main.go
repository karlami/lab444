package main

import (
	"fmt"
	"os"
)

type ResultAnalysis struct {
	Sentiment string `json:"sentiment"`
	Score     int    `json:"score"`
}

func main() {

	//Get the path of where is located the file
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd + "TEMP_FILES")

	sentResult, scoreResult := sentapi("./filesTest/pruebaPositivo.txt")
	fmt.Printf(sentResult, scoreResult)

	/*body := &ResultAnalysis{
		Sentiment: sentResult,
		Score:     scoreResult,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req, _ := http.NewRequest("POST", "https://httpbin.org/post", buf)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)

	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)*/
}
