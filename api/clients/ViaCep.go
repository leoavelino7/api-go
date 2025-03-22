package clients

import (
	"encoding/json"
	"net/http"
	"strings"
)

type viaCEPResponse struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf string `json:"uf"`
	Ibge string `json:"ibge"`
}

type DTOCep struct {
	Cep string `json:"cep"`
	Street string `json:"street"`
	Complement string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City string `json:"city"`
	State string `json:"state"`
	Ibge string `json:"ibge"`
}

func fromViaCEPResponse(response *viaCEPResponse) *DTOCep {
	return &DTOCep{
		Cep: response.Cep,
		Street: response.Logradouro,
		Complement: response.Complemento,
		Neighborhood: response.Bairro,
		City: response.Localidade,
		State: response.Uf,
		Ibge: response.Ibge,
	}
}

func Get(cep string) (*DTOCep, error) {
	cep = strings.ReplaceAll(cep, "-", "")
	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var viaCEPResponse viaCEPResponse

	if err := json.NewDecoder(response.Body).Decode(&viaCEPResponse); err != nil {
		return nil, err
	}

	return fromViaCEPResponse(&viaCEPResponse), nil
}