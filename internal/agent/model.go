package agent

type AgentConf struct {
    Targets []Target `json:"targets"`
    Interval int  `json:"interval"`
}

type Target struct {
    Host string `json:"host"`
    Port int `json:"port"`
    Status string `json:"res"`
}
