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
	//client := payment.Client{ID:0, Name:"Test"} // Получаем клиента
	//user := "Test"// Получаем пользователя

	//Получаем список сервисов, доступных клиенту

	
	services := []payment.Service{
		{ID: 1, Name:"KCell", ShortName:"KCELL", Description:""},
		{ID: 2, Name:"Altel", ShortName:"ALTEL", Description:""}}
	c.JSON(200, services)
}

func GetHolidays(c *gin.Context) {
	c.JSON(200, "")
}
