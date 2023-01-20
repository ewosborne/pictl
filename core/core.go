package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Toggle struct {
	Command string
	Delay   string
}

// http://pi.hole/admin/api.php?disable=300&auth=PWHASH 470

func Picmd(cmd Toggle) {
	fmt.Println("cmd", cmd.Command, cmd.Delay)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://pi.hole/admin/api.php", nil)
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()
	q.Add("auth", os.Getenv("PIPWHASH"))
	q.Add(cmd.Command, cmd.Delay)

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	fmt.Println("query", req.URL.Host, req.URL.Path, req.URL.RawQuery)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))

}
