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
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, id := range mac() {
		fmt.Printf("Mac %v\n", id)

		var roteador map[string]interface{}

		req, err := http.NewRequest(http.MethodGet, "https://app.ccomtelecom.com.br/api/v2/device/update/"+id+"/", nil)
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

		if err := json.Unmarshal([]byte(body), &roteador); err != nil {
			log.Printf("Could not unmarshal json response byte -%v\n", err)
		}
		if roteador["online_status"] == true {
			fmt.Print(roteador["pppoe_user"])
			fmt.Printf(" - Cliente online. Não pode excluir.\n")
		} else {
			fmt.Print(roteador["pppoe_user"])
			fmt.Printf(" - Cliente offline. Pode excluir.\n")
		}
	}

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func mac() []string {
	macsList, err := os.ReadFile("macs.txt")
	checkNilError(err)
	macs := string(macsList[:])

	mac := strings.Fields(macs)
	return mac
}
