package main

import (
	"log"
	interview "zpi/server/shared/kitex_gen/interview/interviewservice"
)

func main() {
	svr := interview.NewServer(new(InterviewServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
