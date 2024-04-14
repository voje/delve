package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voje/delve/internal/agent"

	log "github.com/sirupsen/logrus"
)

const serverAddr = "localhost:13000"

var testConf = agent.AgentConf {
    Health: 13,
    Targets: []agent.Target {
        { Host: "192.168.1.113", Port: 22 },
        { Host: "192.158.1.113", Port: 55 },
    },
}

func main() {
    r := gin.Default()

    r.POST("/agents", func(c *gin.Context) {
        currentConf := agent.AgentConf{}
        c.BindJSON(&currentConf)
        log.Infof("currentConf: %+v", currentConf)

        // Need to send conf based on agent ID/Hostname
        c.JSON(http.StatusOK, testConf)
    })

    r.Run(serverAddr)
}
