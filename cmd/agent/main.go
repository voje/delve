package main

import (
	"github.com/gin-gonic/gin"
	"github.com/voje/delve/internal/agent"
)

func main() {
    a := agent.NewAgent()
    go a.Run()

    r := gin.Default()

    r.POST("/configure", func(c *gin.Context) {
        conf := &agent.AgentConf{} 
        c.BindJSON(&conf)
        a.UpdateConf(conf)
    })

    r.Run()
}
