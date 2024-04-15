package agent

type AgentConf struct {
    ID string `json:"id"`
    Targets []Target `json:"targets"`
    Interval int  `json:"interval"`
}

type Target struct {
    Host string `json:"host"`
    Port int `json:"port"`
    Status string `json:"res"`
}
