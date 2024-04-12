package main

import (
	"fmt"
	"net"
	"time"
    _ "encoding/json"
    "github.com/gin-gonic/gin"
    "os"

	log "github.com/sirupsen/logrus"
)

type Target struct {
    Host string `json:"host"`
    Port int `json:"port"`
}

type ObserverState struct {
    Targets []Target `json:"targets"`
    Health time.Duration `json:"health"`
}

type Observer struct {
    Interval time.Duration
    State *ObserverState
}

func NewObserver(intervalSec int) *Observer {
    return &Observer{
        Interval: time.Duration(time.Second * time.Duration(intervalSec)),
        State: &ObserverState {
            Health: time.Duration(5 * time.Second),
        },
    }
}

func (o *Observer) Run() {
    for {
        log.Infof("Health: %v", o.State.Health)
        for _, t := range o.State.Targets {
            e := t.DelveTCP()
            log.Info(e)
        }
        time.Sleep(o.Interval)
        o.State.Health = o.State.Health - o.Interval
        if (o.State.Health <= 0) {
            log.Info("Out of health!")
            os.Exit(0)
        }
    }
}

func (o *Observer) SetState(state *ObserverState) {
    o.State = state 
}

func NewTarget(host string, port int) *Target {
    return &Target {
        Host: host,
        Port: port,
    }
}

func (t *Target) DelveTCP() error {
    timeout := time.Duration(time.Second)
    c, e := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", t.Host, t.Port), timeout) 
    if e == nil {
        defer c.Close()
    }
    return e
}

func main() {
    ob := NewObserver(1)
    go ob.Run()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


    r.POST("/set-state", func(c *gin.Context) {
        state := &ObserverState{} 
        c.BindJSON(&state)
        ob.SetState(state)
    })

    r.Run()
}
