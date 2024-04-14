package main

import (
	"github.com/gin-gonic/gin"
	"github.com/voje/delve/internal/agent"
)

const server = "http://localhost:13000"

func main() {
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
