package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/json-iterator/go"
)

func help() {
	fmt.Println("Usage: teargas <URL> [output file]")
	fmt.Println("Example: teargas http://example.com/ output.json")
}

func main() {
	var (
		URL string // URL to be tested
		OutputFile string // Output file
		SaveToFile bool // Save output to file
		JSONData jsoniter.Any // JSON data
		httpClient http.Client // HTTP client
	)


	// Leer parámetros de la línea de comandos.
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	URL = os.Args[1]

	// Argumentos con dependencia. (OutputFile depende de SaveToFile)
	if len(os.Args) > 2 {
		OutputFile = os.Args[2]
		SaveToFile = true
	} else {
		SaveToFile = false
	}

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
	if SaveToFile {
		if err := ioutil.WriteFile(OutputFile, body, 0644); err != nil {
			fmt.Println("Error al guardar el JSON en el archivo:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(JSONData)
	}
}
