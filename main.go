package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	dbuser, dbpassword, dbname, dbhost string
	configFile                         = flag.String("configFile", "config.yml", "The confgiguration file.")
	e                                  = Engine{}
	router                             *gin.Engine
)

func init() {
	flag.Parse()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Printf("Using config: %s\n", viper.ConfigFileUsed())

	dbuser := viper.GetString("db.user")
	dbpassword := viper.GetString("db.password")
	dbname := viper.GetString("db.dbname")
	dbhost := viper.GetString("db.host")

	e.Initialize(dbuser, dbpassword, dbname, dbhost)
	//ensureTableExists()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	initializeRoutes()
	router.Run()
}
