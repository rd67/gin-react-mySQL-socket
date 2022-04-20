package middlewares

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/helpers"
	"github.com/sirupsen/logrus"
)

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var log = logrus.New()

		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&logrus.JSONFormatter{})
		log.Out = os.Stdout

		//	Start Timer
		start := time.Now()

		//	Process Request
		c.Next()

		//	Stop Timer
		duration := helpers.GetDurationInMillseconds(start)

		currentDate := start.Format("2006-01-02")

		// You could set this to any `io.Writer` such as a file
		file, err := os.OpenFile("logs/"+currentDate+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Info("Failed to log to file, using default stderr")
		}

		entry := log.WithFields(logrus.Fields{
			"client_id":  helpers.GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    helpers.GetUserId(c),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}

	}
}
