package main

import (
	"ciucurdaniel.com/pet-adoption-server/controllers"
	"ciucurdaniel.com/pet-adoption-server/database"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("loading server configuration")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	LoadAppConfig()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	log.Println("Starting server ...")

	router := mux.NewRouter()

	// TODO: on root give the angular frontend
	// http.Handle("/", ...angular-frontend)

	RegisterAllRoutes(router)

	log.Println(fmt.Sprintf("Starting server on port: %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port),router))

}

func RegisterAllRoutes(router *mux.Router) {
	RegisterUserRoutes(router)
	log.Println("User routes initialized")
	RegisterPostRoutes(router)
	log.Println("Post routes initialized")
}
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
}

func RegisterPostRoutes(router *mux.Router){
	log.Println("Not yet implemented")
}