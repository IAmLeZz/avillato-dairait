package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"
)

type OperationRequest struct {
	Expression string `json:"expression"`
}

type OperationResponse struct {
	Result    float64 `json:"result"`
	Error     string  `json:"error"`
	Time      string  `json:"time"`
	Operation string  `json:"operation"`
}

var history = make([]OperationResponse, 0)

// Verificar y calcular la expresión recibida
func calculate(expression string) (float64, error) {
	if strings.HasPrefix(expression, "+") {
		expression = "0" + expression
	}

	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	return result.(float64), nil
}

// Recibir y verificar operación. Usar la función calculate para obtener un resultado que será devuelto al usuario y añadir al historial.
func postOperation(c *gin.Context) {
	var request OperationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := calculate(request.Expression)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	response := OperationResponse{
		Result:    result,
		Error:     "",
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		Operation: request.Expression,
	}

	history = append(history, response)

	c.JSON(http.StatusOK, response)
}

func getHistory(c *gin.Context) {
	c.JSON(http.StatusOK, history)
}

// Utilizar CORS para compatibilidad con navegadores.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}

func main() {
	router := gin.Default()
	router.Use(CORS())
	router.POST("/operation", postOperation)
	router.GET("/history", getHistory)
	router.Run(":8080")
}
