package main

import (
	"encoding/json"
	"fmt"
)

type City struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	InEurope bool   `json:"in_europe"`
}

func main() {
	citiesJson := `
[
	{
		"name": "Amsterdam",
		"country": "NL",
		"in_europe": true
	},
	{
		"name": "Buenos Aires",
		"country": "AR",
		"in_europe": false
	}
]`
	var unMarshalled []City
	err := json.Unmarshal([]byte(citiesJson), &unMarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
	}
	fmt.Println("Un-marshalled:", unMarshalled)

	warsaw := City{
		Name:     "Warsaw",
		Country:  "PL",
		InEurope: true,
	}
	vitoria := City{"Vitoria", "BR", false}
	var newCities []City
	newCities = append(newCities, warsaw, vitoria)
	newCitiesJson, err := json.MarshalIndent(newCities, "", "  ")
	if err != nil {
		fmt.Println("error marshalling", err)
	}
	fmt.Println("Marshalled:", string(newCitiesJson))
}
