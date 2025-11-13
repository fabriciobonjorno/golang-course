package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Address struct {
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
	http.HandleFunc("/", findZipCodeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func findZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	zipCodeParam := r.URL.Query().Get("zip_code")
	if zipCodeParam == "" {
		http.Error(w, "Missing zip_code parameter", http.StatusBadRequest)
		return
	}

	addr, err := FindZipCode(zipCodeParam)
	if err != nil {
		http.Error(w, "Error fetching zip code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the address as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(addr)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	// Write the JSON response
	// json.NewEncoder(w).Encode(addr)
	w.Write(b)
}

func FindZipCode(zipCode string) (*Address, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var a Address
	if err := json.Unmarshal(body, &a); err != nil {
		return nil, err
	}
	return &a, nil
}
