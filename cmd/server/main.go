package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voje/delve/internal/agent"

	log "github.com/sirupsen/logrus"
)

const serverAddr = "localhost:13000"

func main() {
    r := gin.Default()

    r.POST("/agents", func(c *gin.Context) {
        conf := &agent.AgentConf{} 
        c.BindJSON(&conf)
        log.Infof("%+v", conf)
        c.Status(http.StatusOK)
    })

    r.Run(serverAddr)
}
