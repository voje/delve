package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
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
    server string
}

func NewAgent(server string) *Agent {
    return &Agent {
        conf: AgentConf{},
        health: time.Duration(time.Second * 10),
        // Health ticks down each second
        interval: time.Duration(time.Second),
        server: server,
    }
}

func (a *Agent) UpdateConf(conf *AgentConf) {
    a.mu.Lock() 
    a.conf = *conf
    a.health = a.health + (time.Duration(conf.Heal) * time.Second)
    log.Infof("%+v", a.conf)
    a.mu.Unlock()
}

func (a *Agent) pingServer() error {
    agentsURL := fmt.Sprintf("%s/agents", a.server)
    // TODO is json.Marshal thread safe?
    b, err := json.Marshal(a.conf)
    if err != nil {
        log.Panic(err)
    }
    req, _ := http.NewRequest("POST", agentsURL, bytes.NewBuffer(b))
    req.Header.Set("Content-Type", "application/json")
    c := &http.Client{}
    res, err := c.Do(req)
    if err != nil {
        log.Panic(err)
    }
    log.Debugf("%+v", res)
    return nil
}

func (a *Agent) Run() {
    for {
        a.pingServer()
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

