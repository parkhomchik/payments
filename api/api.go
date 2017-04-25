package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
	"strconv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	//"github.com/parkhomchik/payments/models"
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
//https://github.com/benschw/go-todo/blob/master/service/service.go

func GetDb() (gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=sservice sslmode=disable password=parkhom4ik")
	return *db, err
}

func Migrate() error {
	db, err := GetDb()
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.AutoMigrate(&api.Todo{})
	return nil
}

func main() {
	Init()

	r := gin.Default()

	//USER
	//r.GET("/users", Clients)							// Получить список всех пользователей
	//r.GET("/users/:id", Clients)						// Получить информацию по пользователю

	//CLIENT
	//r.GET("/clients", Clients)							// Получить список всех клиентов
	//r.GET("/clients/:id", Clients)						// Получить информацию по клиенту
	//r.POST("/clients", ClientInsert)					// Новый клиент
	//r.PUT("/clients", ClientUpdate)						// Изменить клиента
	//r.DELETE("/client/:id", ClientDelete)				// Удалить клиента (Enable = false)

	//CLIENT TO SERVICE
	//r.GET("/clients/:id/services", ClientServices)
	//r.POST("/clients/:id/services", ClientServicesInsert)

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