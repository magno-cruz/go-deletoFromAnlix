//package main

/*
import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//url := "https://app.ccomtelecom.com.br/api/v2/device/update"
	//response, err := http.Get(url)

	//req, _ := http.NewRequest("GET", "https://app.ccomtelecom.com.br/api/v2/device/update/", nil)
	//req.SetBasicAuth("magno", "102030")

	request, err := http.NewRequest(http.MethodGet, "https://app.ccomtelecom.com.br/api/v2/device/update/FHTT94087C20", nil)

	request.SetBasicAuth("magno", "10203040")

	if err != nil {
		fmt.Printf("There was an error from the API request %s", err.Error())
	} else {
		responseData, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Printf("There was an error from parsing the request body %s", err.Error())
		} else {
			fmt.Sprint(string(responseData))
		}
	}
}
*/
/*
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Roteador struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	//Type    string `json:"type"`
}

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://app.ccomtelecom.com.br/api/v2/device/update/FHTT94087C20", nil)

	//request.SetBasicAuth("magno", "10203040")

	if err != nil {
		log.Printf("Could not prepare a new request %v", err)
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("Authorization", "Basic bWFnbm86MTAyMDMwNDA=")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Could not prepare a new request %v", err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body %v", err)
	}

	var roteador Roteador
	fmt.Printf(string(responseBytes))
	if err := json.Unmarshal([]byte(responseBytes), &roteador); err != nil {
		log.Printf("Could not unmarshal json response byte -%v", err)
	}
	fmt.Printf(roteador.Message)
}
*/

package main

import (
	"fmt"
	"io/ioutil"

	//"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://app.ccomtelecom.com.br/api/v2/device/update/FHTT94087C20", nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("magno", "10203040")
	req.Header.Add("Content-Type", "application/json")
	req.Close = true

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		panic("Non 2xx response from server, request" + response.Status)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}
