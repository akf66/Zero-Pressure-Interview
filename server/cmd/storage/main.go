package main

import (
	"log"
	storage "zpi/server/shared/kitex_gen/storage/storageservice"
)

func main() {
	svr := storage.NewServer(new(StorageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
