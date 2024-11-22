package storagemodule

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Id           int
	Name         string
	Status       bool
	Creationdate time.Time
}

func (t Task) String() string {
	return fmt.Sprintf("%s\t%s\t%s",
		t.StatusToString(),
		t.Name,
		t.Creationdate.Format(time.DateOnly))
}

func (t Task) StringWithId() string {
	return fmt.Sprintf("%d\t%s\t%s\t%s",
		t.Id,
		t.StatusToString(),
		t.Name,
		t.Creationdate.Format(time.DateOnly))
}

func (t Task) StatusToString() string {
	if t.Status {
		return "[✓]"
	} else {
		return "[X]"
	}
}

var DB *sql.DB

func InitDB(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Create tasks table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		status BOOLEAN NOT NULL DEFAULT 0,
		creation_date TEXT NOT NULL
	);`
	_, err = DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func LoadStorage() ([]Task, error) {
	rows, err := DB.Query("SELECT id, name, status, creation_date FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var creationDate string

		if err := rows.Scan(&task.Id, &task.Name, &task.Status, &creationDate); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		task.Creationdate, err = time.Parse(time.DateOnly, creationDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse creation date: %w", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func AppendTaskToDB(task Task) error {
	_, err := DB.Exec(
		"INSERT INTO tasks (name, status, creation_date) VALUES (?, ?, ?)",
		task.Name, task.Status, task.Creationdate.Format("2006-01-02"),
	)
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}
	return nil
}

func UpdateTaskStatusById(id int, status bool) error {
	_, err := DB.Exec("UPDATE tasks SET status = ? WHERE id = ?", status, id)
	if err != nil {
		return fmt.Errorf("failed to update task status: %w", err)
	}
	return nil
}

func UpdateTaskStatusByName(namePart string, status bool) error {
	query := `
	SELECT id, name, status, creation_date
	FROM tasks
	WHERE name LIKE ?;
	`
	rows, err := DB.Query(query, "%"+namePart+"%")
	if err != nil {
		return fmt.Errorf("failed to search tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var creationDate string
		if err := rows.Scan(&task.Id, &task.Name, &task.Status, &creationDate); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}
		task.Creationdate, err = time.Parse("2006-01-02", creationDate)
		if err != nil {
			return fmt.Errorf("failed to parse creation date: %w", err)
		}
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks matched the given name.")
		return nil
	}

	if len(tasks) == 1 {

		_, err := DB.Exec("UPDATE tasks SET status = ? WHERE id = ?", status, tasks[0].Id)
		if err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}
		fmt.Printf("Task '%s' updated successfully.\n", tasks[0].Name)
		return nil
	}

	fmt.Println("Multiple tasks matched:")
	for i, task := range tasks {
		fmt.Printf("[%d] %s (Status: %s, Created: %s)\n", i+1, task.Name, task.StatusToString(), task.Creationdate.Format(time.DateOnly))
	}

	fmt.Print("Enter the number of the task to update: ")
	var choice int
	_, err = fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(tasks) {
		return fmt.Errorf("invalid selection")
	}

	selectedTask := tasks[choice-1]
	_, err = DB.Exec("UPDATE tasks SET status = ? WHERE id = ?", status, selectedTask.Id)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	fmt.Printf("Task '%s' updated successfully.\n", selectedTask.Name)
	return nil
}
