package main

import (
	"flag"
	"fmt"
	"os/exec"
	"vquality/functions"
)

func main() {
	// Declarar variables para los argumentos que deseas obtener
	var inputURL string

	// Definir los argumentos y asignarlos a las variables
	flag.StringVar(&inputURL, "i", "", "URL del video")

	// Parsear los argumentos de la línea de comandos
	flag.Parse()

	// Verificar si se proporcionó la URL del video
	if inputURL == "" {
		fmt.Println("Debes proporcionar la URL del video con el argumento -i")
		return
	}

	// Aquí puedes usar inputURL en tu programa
	fmt.Println("URL del video:", inputURL)

	cmd := "ffmpeg -i %s -vf blackdetect=d=0.1:pix_th=.1 -f null - 2>&1 | grep \"black_start\\|black_end\\|Duration\""
	cmd = fmt.Sprintf(cmd, inputURL)
	fmt.Println(cmd)
	rs := exec.Command("bash", "-c", cmd)
	combinedOut, err := rs.CombinedOutput()
	if err != nil {
		// output erro command in terminal
		fmt.Println("Error:", err)
		//finish program
		return
	}

	// Convierte la salida combinada a una cadena de texto
	// Llamar a la función para analizar la salida
	videoInfo, err := functions.ParseOutput(string(combinedOut))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Duration:", videoInfo.Duration)
	fmt.Println("BlackStart:", videoInfo.BlackStart)
	fmt.Println("BlackEnd:", videoInfo.BlackEnd)
	fmt.Println("BlackDuration:", videoInfo.BlackDuration)
	fmt.Println("QualityVideo:", videoInfo.QualityVideo)

}
