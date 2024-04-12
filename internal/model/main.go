package model

import (
    "time"
    _ "encoding/json"
)

type Target struct {
    Host string `json:"host"`
    Port int `json:"port"`
}

type AgentConf struct {
    Targets []Target `json:"targets"`
    Health time.Duration `json:"health"`
}

