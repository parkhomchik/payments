package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"github.com/parkhomchik/payments/models"
)

type Configuration struct {
	Port int
}

var config Configuration

func Init(){
	bytes, err := ioutil.ReadFile("configuration.json")

	fmt.Println("len",len(bytes))

	if err != nil {
		fmt.Println(err)
	}	
	if json.Unmarshal(bytes, &config) != nil {
		fmt.Println(err)
	}

	fmt.Println(config)
}

func main() {
	Init()

	r := gin.Default()
	r.GET("/services", Services) 					// Получить все сервисы
	r.GET("/services/:service", Services)			// Получить информацию по сервису
													// Выполнить команду сервиса
													// Получить информацию по операции
	r.Run(":" + strconv.Itoa(config.Port))
}

func Services(c *gin.Context) {
	serv := payment.Service{ID: 1, Name:"KCell", ShortName:"KCELL", Description:""}	
	c.JSON(200, serv)
}

func GetHolidays(c *gin.Context) {
	c.JSON(200, "")
}
