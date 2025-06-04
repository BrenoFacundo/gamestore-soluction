package controller

import (
	"net/http"
	"strconv"

	"github.com/brenofacundo/gamestore-soluction/model"
	"github.com/brenofacundo/gamestore-soluction/usecase"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	usecase usecase.GameUsecase
}

func NewNoteController(usecase usecase.GameUsecase) GameController {
	return GameController{
		usecase: usecase,
	}
}

func (gameController *GameController) Teste(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ola mundo",
	})
}

func (gameController *GameController) CreateGame(ctx *gin.Context){
	var game model.Game

	err := ctx.BindJSON(&game)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedGame, err := gameController.usecase.CreateGame(game)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedGame)
}

func (gameController *GameController) GetAllGames(ctx *gin.Context) {
	games, err := gameController.usecase.GetAllGames()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, games)
}

func (gameController *GameController) GetGamebyId(ctx *gin.Context) {
	id := ctx.Param("gameId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Massage": "id do jogo não pode ser nulo",
		})
		return
	}

	gameId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Massage": "id do jogo precisa ser um numero",
		}) 
		return
	}

	game, err := gameController.usecase.GetGamebyId(gameId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if game == nil {
		
		ctx.JSON(http.StatusNotFound, gin.H{
			"Massage": "esse título não foi encontrado",
		}) 
		return
	}
	ctx.JSON(http.StatusOK, game)
}

func (gameController *GameController) UpdateGame(ctx *gin.Context) {
	id := ctx.Param("gameId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Massage": "id não pode ser nulo",
		})
		return
	}

	gameId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Massage": "id precisa ser um numero",
		}) 
		return
	}

	var game model.Game
	err = ctx.ShouldBindJSON(&game)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Massage": "JSON inválido"})
		return
	}

	game.ID = gameId
	updatedGame, err := gameController.usecase.UpdateGame(&game)
	if err != nil {
		
		ctx.JSON(http.StatusInternalServerError, gin.H{"Massage": "Erro ao atualizar"})
		return
	}

	ctx.JSON(http.StatusOK, updatedGame)
	
}


func (gameController *GameController) DeleteGame(ctx *gin.Context) {
	id := ctx.Param("gameId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Massage": "id não pode ser nulo",
		})
		return
	}
	gameId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Massage": "id precisa ser um numero",
		}) 
		return
	}

	err = gameController.usecase.DeleteGame(gameId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Jogo deletado com sucesso"})
}