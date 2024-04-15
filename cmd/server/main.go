package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voje/delve/internal/agent"

	log "github.com/sirupsen/logrus"
)

const serverAddr = "localhost:13000"

var testConf = agent.AgentConf {
    Targets: []agent.Target {
        { Host: "192.168.1.113", Port: 22 },
        { Host: "192.158.1.113", Port: 55 },
    },
}

type Server struct {
	Agents map[string]agent.AgentConf
}

func main() {
    r := gin.Default()

    s := Server{
	Agents: make(map[string]agent.AgentConf),
    }

    r.Static("/assets/", "./internal/server/assets/")

    r.POST("/agent", func(c *gin.Context) {
	currentConf := agent.AgentConf{}
	c.BindJSON(&currentConf)
	log.Infof("currentConf: %+v", currentConf)

	// TODO merge
	s.Agents[c.Param("id")] = currentConf

        // Need to send conf based on agent ID/Hostname
        c.JSON(http.StatusOK, testConf)
    })

    r.LoadHTMLGlob("./internal/server/templates/*.tmpl")

    r.GET("/", func(c *gin.Context) {
	log.Infof("%v", s)
        c.HTML(http.StatusOK, "hello.tmpl", gin.H{
		"server": s,
        })
    })

    r.Run(serverAddr)
}
