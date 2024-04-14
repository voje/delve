package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Agent struct {
    mu sync.Mutex
    conf AgentConf
    server string
}

func NewAgent(server string) *Agent {
    return &Agent {
        conf: AgentConf{
            Health: 10,
        },
        // Health ticks down each second
        server: server,
    }
}

func (a *Agent) UpdateConf(conf *AgentConf) {
    a.mu.Lock() 
    a.conf = *conf
    a.mu.Unlock()
}

func (a *Agent) pingServer() error {
    agentsURL := fmt.Sprintf("%s/agents", a.server)
    // TODO is json.Marshal thread safe?
    b, err := json.Marshal(a.conf)
    if err != nil {
        return err
    }
    req, _ := http.NewRequest("POST", agentsURL, bytes.NewBuffer(b))
    req.Header.Set("Content-Type", "application/json")
    c := &http.Client{}
    res, err := c.Do(req)
    if err != nil {
        return err
    }
    conf := AgentConf{}
    json.NewDecoder(res.Body).Decode(&conf)
    a.UpdateConf(&conf)
    return nil
}

func (a *Agent) scanTargets() {
    for _, t := range a.conf.Targets {
        t.DelveTCP()
        log.Debugf("DD %+v", t)
    }
}

func (a *Agent) Run() {
    for {
        // Pinging the server pulls configuration (targets) and extends health
        a.pingServer()
        log.Infof("Health: %v", a.conf.Health)
        log.Debugf("Conf: %v", a.conf)
        a.mu.Lock()
        if (a.conf.Health > 0) {
            a.scanTargets()
            log.Debugf("%+v", a.conf.Targets)
            a.conf.Health = a.conf.Health - 1
        } else {
            a.conf.Targets = []Target{}
        }
        a.mu.Unlock()
        time.Sleep(time.Duration(time.Second))
    }
}

func (t *Target) DelveTCP() {
    timeout := time.Duration(time.Second)
    c, e := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", t.Host, t.Port), timeout) 
    if c != nil {
        defer c.Close()
    }

    if e != nil {
        t.Status = fmt.Sprintf("Unreachable: (%v)", e)
    } else {
        t.Status = "OK"
    }

    log.Debugf("DelveTCP: %v", t)
}

