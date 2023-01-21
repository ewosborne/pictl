package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Toggle struct {
	Command string
	Delay   int
	Host    string
}

func NewToggle(cmd *cobra.Command, command string) Toggle {
	// TODO: check for errors and handle them
	host, _ := cmd.Flags().GetString("host")
	d, _ := cmd.Flags().GetInt("delay")
	slog.Info(fmt.Sprintf("delay is %d", d))
	return Toggle{Host: host, Delay: d, Command: command}
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
	q.Add(cmd.Command, strconv.Itoa(cmd.Delay))

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()
	slog.Info(req.URL.String())

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
