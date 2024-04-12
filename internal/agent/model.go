package agent

type AgentConf struct {
    Heal    int  `json:"heal"`
    Targets []Target `json:"targets"`
}

type Target struct {
    Host string `json:"host"`
    Port int `json:"port"`
}
