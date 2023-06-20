package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/json-iterator/go"
)

func main() {
	var (
		URL string // URL to be tested (Flag)
		OutputFile string // Output file (Flag)
		JSONData jsoniter.Any // JSON data
		httpClient http.Client // HTTP client
	)

	flag.StringVar(&URL, "url", "https://localhost:8080/", "URL to be tested")
	flag.StringVar(&OutputFile, "output", "output.json", "Output file")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Printf("Example: %s -url http://example.com/ output.json\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("Please report bugs at: ventgrey@gmail.com")
	}

	flag.Parse()

	// Realizar petición HTTP y obtener respuesta.
	resp, err := httpClient.Get(URL)
	if err != nil {
		fmt.Println("Error al hacer la petición HTTP:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Leer cuerpo de la respuesta HTTP y convertirlo a JSON.
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error al leer el cuerpo de la respuesta HTTP:", err)
		os.Exit(1)
	}

	if err := jsoniter.Unmarshal(body, &JSONData); err != nil {
		fmt.Println("Error al convertir el cuerpo de la respuesta HTTP a JSON:", err)
		os.Exit(1)
	}

	// Guardar el JSON en un archivo si es necesario.
	if err := ioutil.WriteFile(OutputFile, body, 0644); err != nil {
		fmt.Println("Error al guardar el JSON en el archivo:", err)
		os.Exit(1)
	}

	jsonString, err := jsoniter.MarshalIndent(JSONData, "", "  ")
	if err != nil {
		fmt.Println("Error al imprimir el JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonString))
}
