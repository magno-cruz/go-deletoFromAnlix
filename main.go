package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var device map[string]interface{}

func main() {
	quantExcluidos := 0
	for _, id := range mac() {
		fmt.Printf("Mac %v\n", id)

		req, err := http.NewRequest("GET", "https://app.ccomtelecom.com.br/api/v2/device/update/"+id+"/", nil)
		checkNilError(err)
		req.SetBasicAuth("magno", "10203040")
		req.Header.Add("Content-Type", "application/json")
		req.Close = true

		client := http.Client{}
		response, err := client.Do(req)
		checkNilError(err)

		if response.StatusCode != http.StatusOK {
			panic("Non 2xx response from server, request" + response.Status + "\n")
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		checkNilError(err)

		if err := json.Unmarshal([]byte(body), &device); err != nil {
			log.Printf("Could not unmarshal json response byte -%v\n", err)
		}
		if checkStatus() {
			fmt.Printf("%v está online. Não foi excluído.\n", id)

		} else {
			req, err := http.NewRequest("DELETE", "https://app.ccomtelecom.com.br/api/v2/device/delete/"+id+"/", nil)
			checkNilError(err)
			req.SetBasicAuth("magno", "10203040")
			req.Header.Add("Content-Type", "application/json")
			req.Close = true

			client := http.Client{}
			response, err := client.Do(req)
			checkNilError(err)

			if response.StatusCode != http.StatusOK {
				panic("Non 2xx response from server, request" + response.Status + "\n")
			}
			fmt.Printf("%v foi excluído.\n", id)
		}
		quantExcluidos++
	}
	fmt.Printf("Foram excluídos %v aparelhos.\n", quantExcluidos)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func mac() []string {
	macsList, err := os.ReadFile("excluir do anlix.txt")
	checkNilError(err)
	macs := string(macsList[:])

	mac := strings.Fields(macs)
	return mac
}

func checkStatus() bool {
	if device["online_status"] == true {
		return true
	} else {
		return false
	}
}
