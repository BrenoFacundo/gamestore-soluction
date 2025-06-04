package repository

import (
	"database/sql"
	"fmt"

	"github.com/brenofacundo/gamestore-soluction/model"
)

type GamesRepository struct {
	connection *sql.DB
}

func NewGamesRepository(connection *sql.DB) GamesRepository{
	return GamesRepository{
		connection: connection,
	}
}

func (gamesRepo *GamesRepository) CreateGame(game model.Game) (int, error){
	var id int

	query, err := gamesRepo.connection.Prepare("INSERT INTO games" + "(nome, preco, plataforma, descricao)" + " VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(game.Name, game.Price, game.Platform, game.Description).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (gamesRepo *GamesRepository) GetAllGames() ([]model.Game, error){
	var allGames []model.Game
	var game model.Game

	query := "select * from games"
	row, err := gamesRepo.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Game{}, err
	}

	for row.Next(){
		err = row.Scan(
			&game.ID,
			&game.Name,
			&game.Price,
			&game.Platform,
			&game.Description)
			if err != nil {
				fmt.Println(err)
				return []model.Game{}, err
			}
			allGames = append(allGames, game)
	}
	row.Close()

	return allGames, nil
}

func (gamesRepo *GamesRepository) GetGamebyId(game_id int) (*model.Game, error){
	var game model.Game
	query, err := gamesRepo.connection.Prepare("SELECT * FROM games WHERE id = $1")
	if err != nil {
		return nil, err
	}
	err = query.QueryRow(game_id).Scan(
		&game.ID,
		&game.Name,
		&game.Price,
		&game.Platform,
		&game.Description)

	if err != nil {
		if err == sql.ErrNoRows{
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &game, err
}

func (gamesRepo *GamesRepository) UpdateGame(game *model.Game) (*model.Game, error) {
	var updateGame model.Game
	query := `
		UPDATE games
		SET nome = $1, preco = $2, plataforma = $3, descricao = $4
		WHERE id = $5
		RETURNING id, nome, preco, plataforma, descricao;
	`

	row := gamesRepo.connection.QueryRow(query,
		game.Name, 
		game.Price, 
		game.Platform, 
		game.Description, 
		game.ID)
	
	err := row.Scan(
		&updateGame.ID,
		&updateGame.Name,
		&updateGame.Price,
		&updateGame.Platform,
		&updateGame.Description)

	if err != nil {
		return nil, err
	}

	return &updateGame, nil
}

func (gamesRepo *GamesRepository) DeleteGame(id int) error {
	query := `DELETE FROM games WHERE id = $1`

	result, err := gamesRepo.connection.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}