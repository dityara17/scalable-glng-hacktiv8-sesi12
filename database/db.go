package database

import (
	"fmt"
	"go-jwt-glng-aditya/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

//Host
//ec2-54-161-189-150.compute-1.amazonaws.com
//Database
//d78nlv2nbbpt0h
//User
//rjltjsigtoinsh
//Port
//5432
//Password
//068a3136cf8397fc6c08b18960ae1dec256f14268555ed46532b7e10c0875816
//URI
//postgres://rjltjsigtoinsh:068a3136cf8397fc6c08b18960ae1dec256f14268555ed46532b7e10c0875816@ec2-54-161-189-150.compute-1.amazonaws.com:5432/d78nlv2nbbpt0h
//Heroku CLI
//heroku pg:psql postgresql-rigid-03338 --app learn-golang-aditya

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbPort   = os.Getenv("DB_PORT")
	dbName   = os.Getenv("DB_NAME")
	db       *gorm.DB
	err      error
)

// StartDB starting open connection
func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbName, dbPort)
	//config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, dbPort, dbName)
	dsn := config

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("error connection to database...", err)
	}

	fmt.Println("success open connection")
	err = db.Debug().AutoMigrate(models.User{}, models.Product{})
	if err != nil {
		log.Println(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
