package main

import (
	"github.com/gin-gonic/gin"
	"github.com/voje/delve/internal/agent"

    log "github.com/sirupsen/logrus"
)

const server = "http://localhost:13000"

func main() {
    log.SetLevel(log.DebugLevel)

    a := agent.NewAgent(server)
    go a.Run()

    r := gin.Default()

    r.POST("/configure", func(c *gin.Context) {
        conf := &agent.AgentConf{} 
        c.BindJSON(&conf)
        a.UpdateConf(conf)
    })

    r.Run()
}
