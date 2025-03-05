package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on request: %s\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on read body: %s\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on unmarshal: %s\n", err)
		}

		fmt.Println(data)
		file, err := os.Create(data.Cep + ".json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on create file: %s\n", err)
		}

		defer file.Close()

		jsonFile, err := json.Marshal(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on marshal: %s\n", err)
		}

		_, err = file.WriteString(string(jsonFile))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on write file: %s\n", err)
		}
	}
}
