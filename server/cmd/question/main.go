package main

import (
	"log"
	question "zpi/server/shared/kitex_gen/question/questionservice"
)

func main() {
	svr := question.NewServer(new(QuestionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
