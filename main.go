package main

import (
	
	"log"
	"net/http"
	"os"
	"runtime"

	
	"github.com/Geekinn/go-micro/app/models"
	"github.com/Geekinn/go-micro/routes"
	"github.com/Geekinn/go-micro/db"
	"github.com/Geekinn/go-micro/app/middlewares"
	"github.com/gin-contrib/gzip"
	"github.com/joho/godotenv"
	

	"github.com/gin-gonic/gin"
)







func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Start the default gin server
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.APIRoutes(r)

	//Start PostgreSQL database
	db.Init()
	// Migrate all the models
	new(models.ArticleModel).Migrate()
	new(models.UserModel).Migrate()

	//Start Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	//Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
	db.InitRedis(1)

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"frameworkVersion": "v1",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	if os.Getenv("SSL") == "TRUE" {

		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		r.Run(":" + port)
	}

}
