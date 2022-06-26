package middlewares

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//  Log to file
func LoggerToFile() gin.HandlerFunc {	
    logFilePath := "/usr/src/app/storage/logs/"
    logFileName := "access.log"
    // Log files 	
    fileName := path.Join(logFilePath, logFileName)	
    // write file 	
    // src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)	
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {	
        fmt.Println("err", err)	
    }	
    // Instantiation 	
    logger := logrus.New()	
    // Set output 	
    logger.Out = src	
    // Set the log level 	
    logger.SetLevel(logrus.DebugLevel)	
    // Format log 	
	logger.SetFormatter(&logrus.TextFormatter{	
		TimestampFormat:"2006-01-02 15:04:05",	
	})

    return func(c *gin.Context) {	
        //  Starting time 	
        startTime := time.Now()	
        //  Processing requests 	
        c.Next()	
        //  End time 	
        endTime := time.Now()	
        //  execution time 	
        latencyTime := endTime.Sub(startTime)	
        //  Request mode 	
        reqMethod := c.Request.Method	
        //  Request routing 	
        reqUri := c.Request.RequestURI	
        //  Status code 	
        statusCode := c.Writer.Status()	
        //  request IP	
        clientIP := c.ClientIP()	
        //  Log format 	
        logger.Infof("| %3d | %13v | %15s | %s | %s |",	
            statusCode,	
            latencyTime,	
            clientIP,	
            reqMethod,	
            reqUri,	
        )	
    }	
}