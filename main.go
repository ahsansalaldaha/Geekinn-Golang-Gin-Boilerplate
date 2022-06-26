package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/Geekinn/go-micro/app/middlewares"
	"github.com/Geekinn/go-micro/app/models"
	"github.com/Geekinn/go-micro/db"
	"github.com/Geekinn/go-micro/routes"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func getLogWriter() (io.Writer)  {
	logFilePath := "/usr/src/app/storage/logs/"
    logFileName := "micro.log"
    // Log files 	
    fileName := path.Join(logFilePath, logFileName)	
    // write file 	
    // src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)	
	src, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return io.MultiWriter(src)
}

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

	gin.DefaultWriter = getLogWriter()

	// get global Monitor object
	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// set middleware for gin
	m.Use(r)

	// r.Use(middlewares.LoggerToFile())
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
