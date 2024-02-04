package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OperationRequest struct {
	Operator string `json:"operator"`
	Num1     int    `json:"num1"`
	Num2     int    `json:"num2"`
}

type OperationResponse struct {
	Result    int    `json:"result"`
	Error     string `json:"error"`
	Time      string `json:"time"`
	Operation string `json:"operation"`
}

var history = make([]OperationResponse, 0)

func calculate(operator string, num1 int, num2 int) (int, error) {
	var result int
	var err error

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	case "%":
		result = num1 % num2
	default:
		err = errors.New("syntax error")
	}

	return result, err
}

func postOperation(c *gin.Context) {
	var request OperationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := calculate(request.Operator, request.Num1, request.Num2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := OperationResponse{
		Result:    result,
		Error:     "",
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		Operation: request.Operator,
	}

	history = append(history, response)

	c.JSON(http.StatusOK, response)
}

func getHistory(c *gin.Context) {
	c.JSON(http.StatusOK, history)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.POST("/operation", postOperation)
	router.GET("/history", getHistory)
	router.Run(":8080")
}
