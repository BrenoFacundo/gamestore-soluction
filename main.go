package main

import (
	"github.com/brenofacundo/gamestore-soluction/controller"
	"github.com/brenofacundo/gamestore-soluction/db"
	"github.com/brenofacundo/gamestore-soluction/repository"
	"github.com/brenofacundo/gamestore-soluction/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbconn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	noteRepository := repository.NewGamesRepository(dbconn)
	noteUsecase := usecase.NewNoteUsecase(noteRepository)
	noteController := controller.NewNoteController(noteUsecase)

	server.GET("/ola", noteController.Teste)
	server.POST("/game", noteController.CreateGame)
	server.GET("/game", noteController.GetAllGames)
	server.GET("/game/:gameId", noteController.GetGamebyId)
	server.PUT("/game/:gameId", noteController.UpdateGame)
	server.DELETE("/game/:gameId", noteController.DeleteGame)

	server.Run(":8000")
}