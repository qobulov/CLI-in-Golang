package postgres

import (
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func TestGetAllTasks(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a repository instance
	gemRepo := Gem{db: db}

	// Call the GetAllTasks method
	tasks, err := gemRepo.GetAllTasks()
	if err != nil {
		t.Fatalf("unexpected error during GetAllTasks: %v", err)
	}

	if len(tasks) == 0 {
		t.Fatalf("expected some tasks, got 0")
	}
}

func TestGetDoneTasks(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a repository instance
	gemRepo := Gem{db: db}

	// Call the GetDoneTasks method
	tasks, err := gemRepo.GetDoneTasks()
	if err != nil {
		t.Fatalf("unexpected error during GetDoneTasks: %v", err)
	}

	if len(tasks) == 0 {
		t.Fatalf("expected some done tasks, got 0")
	}

}

func TestGetNotDoneTasks(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a repository instance
	gemRepo := Gem{db: db}

	// Call the GetNotDoneTasks method
	tasks, err := gemRepo.GetNotDoneTasks()
	if err != nil {
		t.Fatalf("unexpected error during GetNotDoneTasks: %v", err)
	}

	if len(tasks) == 0 {
		t.Fatalf("expected some not done tasks, got 0")
	}

}

func TestGetNextTasks(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a repository instance
	gemRepo := Gem{db: db}

	// Call the GetNextTasks method
	_, err = gemRepo.GetNextTasks()
	if err != nil {
		t.Fatalf("unexpected error during GetNextTasks: %v", err)
	}

}

func TestGetByDay(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a repository instance
	gemRepo := Gem{db: db}

	// Call the GetByDay method
	task, err := gemRepo.GetByDay(4)
	if err != nil {
		t.Fatalf("unexpected error during GetByDay: %v", err)
	}

	if task.Day != 4 {
		t.Fatalf("expected task day to be 4, got %v", task.Day)
	}
}

func TestDoDone(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a repository instance
	gemRepo := Gem{db: db}

	// Call the DoDone method
	err = gemRepo.DoDone(5)
	if err != nil {
		t.Fatalf("unexpected error during DoDone: %v", err)
	}

}
