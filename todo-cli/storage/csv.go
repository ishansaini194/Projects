package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"todo-cli/models"
)

const fileName = "tasks.csv"

func LoadTasks() ([]models.Task, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	tasks := []models.Task{}
	for _, row := range rows {
		if len(row) < 4 {
			continue // skip bad rows
		}

		id, _ := strconv.Atoi(row[0])
		completed, _ := strconv.ParseBool(row[2])
		createdAt, _ := time.Parse(time.RFC3339, row[3])

		tasks = append(tasks, models.Task{
			ID:        id,
			Title:     row[1],
			Completed: completed,
			CreatedAt: createdAt,
		})
	}
	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	file, err := os.Create(fileName) // overwrite file
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, t := range tasks {
		row := []string{
			strconv.Itoa(t.ID),
			t.Title,
			strconv.FormatBool(t.Completed),
			t.CreatedAt.Format(time.RFC3339),
		}
		writer.Write(row)
	}

	return nil
}
