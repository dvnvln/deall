package main

import (
	"log"
	"net/http"
	"os"

	userController "github.com/dvnvln/deallscrud/controller"
	userHandler "github.com/dvnvln/deallscrud/handler"
	userRepository "github.com/dvnvln/deallscrud/repository"
)

func main() {

	var (
		userRepo, err = userRepository.New(os.Getenv("DBURI"), os.Getenv("DBNAME"))
		userCnt       = userController.New(userRepo)
		userHnd       = userHandler.New(userCnt)
	)
	if err != nil {
		log.Fatal("Unable to start app")
	}
	defer userRepo.Disconnect()
	// port := "8000"

	// if fromEnv := os.Getenv("PORT"); fromEnv != "" {
	// 	port = fromEnv
	// }
	port := os.Getenv("PORT")
	log.Printf("Starting up on http://localhost:%s", port)

	r := initRoute(userHnd)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
