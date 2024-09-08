package storage

import (
	"cli/gym/models"
)

type Storage interface {
	GetAllTasks() ([]models.Tasks, error)
	GetDoneTasks() ([]models.Tasks, error)
	GetNotDoneTasks() ([]models.Tasks, error)
	GetNextTasks() (models.Tasks, error)
	GetByDay(day int) (models.Tasks, error)
	DoDone(day int) error
}
