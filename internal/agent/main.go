package agent

import (
	_ "encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Agent struct {
    mu sync.Mutex
    conf AgentConf
    health time.Duration
    interval time.Duration
}

func NewAgent() *Agent {
    return &Agent {
        conf: AgentConf{},
        health: time.Duration(time.Second * 10),
        // Health ticks down each second
        interval: time.Duration(time.Second),
    }
}

func (a *Agent) UpdateConf(conf *AgentConf) {
    a.mu.Lock() 
    a.conf = *conf
    a.health = a.health + (time.Duration(conf.Heal) * time.Second)
    log.Infof("%+v", a.conf)
    a.mu.Unlock()
}

func (a *Agent) Run() {
    for {
        log.Infof("Health: %v", a.health)
        for _, t := range a.conf.Targets {
            e := t.DelveTCP()
            log.Info(e)
        }
        time.Sleep(a.interval)
        a.health = a.health - a.interval
        if (a.health <= 0) {
            log.Info("Out of health!")
            os.Exit(0)
        }
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

