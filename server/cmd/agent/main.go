package main

import (
	"log"
	agent "zpi/server/shared/kitex_gen/agent/agentservice"
)

func main() {
	svr := agent.NewServer(new(AgentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
