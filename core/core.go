package core

import (
	"fmt"
	"golang.org/x/exp/slog"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Toggle struct {
	Command string
	Delay   string
	Host    string
}

func NewToggle() Toggle {
	return Toggle{Host: "pi.hole"}
}

func Picmd(cmd Toggle) {
	slog.Info("in picmd")
	slog.Info(cmd.Host)
	reqstr := fmt.Sprintf("http://%s/admin/api.php", cmd.Host)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, reqstr, nil)
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()
	q.Add("auth", os.Getenv("PIPWHASH"))
	q.Add(cmd.Command, cmd.Delay)

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

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

	//fmt.Println(resp.Status)
	fmt.Println(string(responseBody))

}
