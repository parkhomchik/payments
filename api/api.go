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

	//USER
	r.GET("/users", Clients)							// Получить список всех пользователей
	r.GET("/users/:id", Clients)						// Получить информацию по пользователю

	//CLIENT
	r.GET("/clients", Clients)							// Получить список всех клиентов
	r.GET("/clients/:id", Clients)						// Получить информацию по клиенту
	r.POST("/clients", ClientInsert)					// Новый клиент
	r.PUT("/clients", ClientUpdate)						// Изменить клиента
	r.DELETE("/client/:id", ClientDelete)				// Удалить клиента (Enable = false)

	//CLIENT TO SERVICE
	r.GET("/clients/:id/services", ClientServices)
	r.POST("/clients/:id/services", ClientServicesInsert)

	//SERVICES
	r.GET("/services", Services) 						// Получить все сервисы
	r.GET("/services/:id", Services)				// Получить информацию по короткому имени сервиса
	r.POST("/services", ServicesInsert)					// Новый сервис
	r.PUT("/services", ServicesUpdate)					// Изменить сервис
	r.DELETE("/services/:id", ServicesDelete)	// Удалить сервис (Enable = false)

	


	err := r.Run(":" + strconv.Itoa(config.Port))
	if err!= nil {
		panic(err)
	}
}

func Services(c *gin.Context) {
	services := []payment.Service{
		{ID: 1, Name:"KCell", ShortName:"KCELL", Description:""},
		{ID: 2, Name:"Altel", ShortName:"ALTEL", Description:""}}
	c.JSON(200, services)
}

func ServicesInsert(c *gin.Context) {
	c.JSON(200, nil)
}

func ServicesUpdate(c *gin.Context) {
	c.JSON(200, nil)
}

func ServicesDelete(c *gin.Context) {
	c.JSON(200, nil)
}

func Clients(c *gin.Context) {
	c.JSON(200, nil)
}

func ClientInsert(c *gin.Context) {
	c.JSON(200, nil)
}

func ClientUpdate(c *gin.Context) {
	c.JSON(200, nil)
}

func ClientDelete(c *gin.Context) {
	c.JSON(200, nil)
}