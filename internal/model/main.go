package model

import (
    "time"
    _ "encoding/json"
)

type Target struct {
    Host string `json:"host"`
    Port int `json:"port"`
    // Scan result
    Res string `json:"res,omitempty"`
}

type AgentConf struct {
    Targets []Target `json:"targets"`
    Health time.Duration `json:"health"`
}

