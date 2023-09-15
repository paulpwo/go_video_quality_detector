package functions

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"net/http"

	"vquality/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DATA_SERVICES_URI string
var DATA_BASE_URI_AUTH string
var DATA_USERNAME string
var DATA_PASSWORD string
var DATA_PREFFIX_AUTH string
var DATA_PREFIX_RECUA string
var API_URL string
var CATEGORIES_API_URL string
var CTX context.Context
var MEDIA_URL string
var DB *gorm.DB
var DB_DATABASE_BI string

func GetFavicon(c *gin.Context) {
	c.File("./assets/favicon.png")
}

func Ping(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
}

type VideoInfo struct {
	Duration      string
	BlackStart    float64
	BlackEnd      float64
	BlackDuration float64
	QualityVideo  float64
}

func Test(c *gin.Context) {
	var req models.RequestTest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": http.StatusBadRequest})
		return
	}

	cmd := "ffmpeg -i %s -vf blackdetect=d=0.1:pix_th=.1 -f null - 2>&1 | grep \"black_start\\|black_end\\|Duration\""
	cmd = fmt.Sprintf(cmd, req.Key)
	rs := exec.Command("bash", "-c", cmd)
	combinedOut, err := rs.CombinedOutput()
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "error", "status": http.StatusOK, "error": err.Error()})
		return
	}

	// Convierte la salida combinada a una cadena de texto
	// Llamar a la función para analizar la salida
	videoInfo, err := ParseOutput(string(combinedOut))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK, "data": videoInfo})
}

func ParseOutput(outputText string) (*VideoInfo, error) {
	// Definir patrones de expresión regular para buscar los valores de interés
	durationPattern := regexp.MustCompile(`Duration: ([0-9:.]+)`)
	blackStartPattern := regexp.MustCompile(`black_start:([0-9.]+)`)
	blackEndPattern := regexp.MustCompile(`black_end:([0-9.]+)`)
	blackDurationPattern := regexp.MustCompile(`black_duration:([0-9.]+)`)

	// Buscar coincidencias en la cadena de salida
	durationMatch := durationPattern.FindStringSubmatch(outputText)
	blackStartMatch := blackStartPattern.FindStringSubmatch(outputText)
	blackEndMatch := blackEndPattern.FindStringSubmatch(outputText)
	blackDurationMatch := blackDurationPattern.FindStringSubmatch(outputText)

	// Verificar si se encontraron todas las coincidencias
	if len(durationMatch) != 2 || len(blackStartMatch) != 2 || len(blackEndMatch) != 2 || len(blackDurationMatch) != 2 {
		fmt.Println("No se pudieron encontrar todos los valores en la salida de ffmpeg")
		fmt.Println("**************************************************************************")
		fmt.Println(outputText)
		fmt.Println(durationMatch)
		fmt.Println(blackStartMatch)
		fmt.Println(blackEndMatch)
		fmt.Println(blackDurationMatch)
		fmt.Println("**************************************************************************")

		return nil, fmt.Errorf("No se pudieron encontrar todos los valores en la salida de ffmpeg")
	}

	// Parsear los valores de las coincidencias y crear un struct VideoInfo
	duration := durationMatch[1]
	blackStart, _ := strconv.ParseFloat(blackStartMatch[1], 64)
	blackEnd, _ := strconv.ParseFloat(blackEndMatch[1], 64)
	blackDuration, _ := strconv.ParseFloat(blackDurationMatch[1], 64)

	videoInfo := &VideoInfo{
		Duration:      duration,
		BlackStart:    blackStart,
		BlackEnd:      blackEnd,
		BlackDuration: blackDuration,
	}
	videoInfo.QualityVideo = calculateVideoQuality(videoInfo)

	return videoInfo, nil
}

func calculateVideoQuality(videoInfo *VideoInfo) float64 {
	// Parsear la duración total del video en segundos
	parts := strings.Split(videoInfo.Duration, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])

	// Separar los segundos y los milisegundos
	secondsAndMillis := strings.Split(parts[2], ".")
	seconds, _ := strconv.Atoi(secondsAndMillis[0])
	millis, _ := strconv.Atoi(secondsAndMillis[1])

	// Calcular la duración total en segundos, incluyendo milisegundos
	totalDurationSeconds := float64(hours*3600+minutes*60+seconds) + float64(millis)/100.0

	// Definir duración mínima y máxima en segundos
	duracionMinima := 30.0

	// Calcular la calidad en base a la cantidad de pantalla negra
	quality := 1.0 - (videoInfo.BlackDuration / totalDurationSeconds)

	// Verificar la duración y ajustar la calidad según los criterios
	if totalDurationSeconds < duracionMinima {
		// Si la duración es menor que la mínima, reducir la calidad en un 30%
		quality *= 0.7
	}

	return quality
}
