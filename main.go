package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	//"os"
	"strings"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

var device map[string]interface{}

func main() {
	a := app.New()
	w := a.NewWindow("Exportador de Macs")
	macsInput := widget.NewMultiLineEntry()
	content := container.NewVBox(macsInput, widget.NewButton("Deletar", func() {
		macs := strings.Split(macsInput.Text, "\n")
		for i := 0; i < len(macs); i++ {
			excluir(macs[i])
		}
	}))
	w.SetContent(widget.NewLabel("Excluir macs do Anlix"))
	w.SetContent(content)
	w.ShowAndRun()

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkStatus() bool {
	if device["online_status"] == true {
		return true
	} else {
		return false
	}
}

func excluir(macs string) {
	req, err := http.NewRequest("GET", "https://app.ccomtelecom.com.br/api/v2/device/update/"+macs+"/", nil)
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
		fmt.Printf("%v está online. Não foi excluído.\n", macs)

	} else {
		req, err := http.NewRequest("DELETE", "https://app.ccomtelecom.com.br/api/v2/device/delete/"+macs+"/", nil)
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
	}
}
