package usecase

import (
	"github.com/brenofacundo/gamestore-soluction/model"
	"github.com/brenofacundo/gamestore-soluction/repository"
)

type GameUsecase struct {
	repository repository.GamesRepository
}

func NewNoteUsecase(repository repository.GamesRepository) GameUsecase {
	return GameUsecase{
		repository: repository,
	}
}

func (gameUsecase *GameUsecase) CreateGame(game model.Game) (model.Game, error) {
	gameId, err := gameUsecase.repository.CreateGame(game)
	if  err != nil {
		return model.Game{}, err
	}

	game.ID = gameId

	return game, nil
}

func (gameUsecase *GameUsecase) GetAllGames() ([]model.Game, error) {
	return gameUsecase.repository.GetAllGames()
}

func (gameUsecase *GameUsecase) GetGamebyId(game_id int) (*model.Game, error) {
	game, err := gameUsecase.repository.GetGamebyId(game_id)
	if err != nil{
		return nil, err
	}

	return game, nil
}

func (gameUsecase *GameUsecase) UpdateGame(game *model.Game) (*model.Game, error) {
	return gameUsecase.repository.UpdateGame(game)
}

func (gameUsecase *GameUsecase) DeleteGame(id int) error {
	return gameUsecase.repository.DeleteGame(id)
}