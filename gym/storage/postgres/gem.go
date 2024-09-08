package postgres

import (
	"cli/gym/models"
	"cli/gym/storage"
	"github.com/jmoiron/sqlx"
)

type Gem struct {
	db *sqlx.DB
}

func NewGem(db *sqlx.DB) storage.Storage {
	return &Gem{db: db}
}

func (g *Gem) GetAllTasks() ([]models.Tasks, error) {

	result := []models.Tasks{}

	err := g.db.Select(&result, "SELECT * FROM gem")
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gem) GetDoneTasks() ([]models.Tasks, error) {

	var result []models.Tasks

	err := g.db.Select(&result, "SELECT * FROM gem where done=$1", "t")
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gem) GetNotDoneTasks() ([]models.Tasks, error) {

	var result []models.Tasks

	err := g.db.Select(&result, "SELECT * FROM gem where done=$1", "f")
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gem) GetNextTasks() (models.Tasks, error) {

	var result models.Tasks

	err := g.db.Get(&result, "SELECT * FROM gem where done=$1 order by day limit 1", "f")
	if err != nil {
		return models.Tasks{}, err
	}

	return result, nil
}

func (g *Gem) GetByDay(day int) (models.Tasks, error) {

	var result models.Tasks

	err := g.db.Get(&result, "SELECT * FROM gem where day=$1", day)
	if err != nil {
		return models.Tasks{}, err
	}

	return result, nil
}

func (g *Gem) DoDone(day int) error {

	_, err := g.db.Exec("UPDATE gem SET done=$1 where day=$2", "t", day)
	if err != nil {
		return err
	}

	return nil
}
