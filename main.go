package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/ssr0016/goHex/handler"
	"github.com/ssr0016/goHex/repository"
	"github.com/ssr0016/goHex/service"
)

func main() {

	initConfig()
	db := initDB()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryDB
	_ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepositoryDB)
	customHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customHandler.GetCustomer).Methods(http.MethodGet)

	log.Printf("Banking service started on port %s", viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("app.port")), router)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) //APP_PORT=5000 go run .

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}

func initDB() *sqlx.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"))

	db, err := sqlx.Connect(viper.GetString("db.driver"), dsn)
	if err != nil {
		log.Fatalln(err)
	}

	// Set the time zone
	_, err = db.Exec("SET TIME ZONE 'Asia/Manila'")
	if err != nil {
		log.Fatalln("Failed to set time zone:", err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
