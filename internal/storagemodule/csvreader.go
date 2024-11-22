package storagemodule

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Reset  = "\033[0m"
)

type Task struct {
	Id           int
	Name         string
	Status       bool
	Creationdate time.Time
}

func (t Task) String() string {
	return fmt.Sprintf("%d\t%s\t%s\t%s",
		t.Id,
		t.Name,
		t.StatusToString(),
		t.Creationdate.Format(time.DateOnly))
}

func (t Task) StatusToString() string {
	if t.Status {
		return "Completed"
	} else {
		return "Not completed"
	}
}

func readCSV(filepath string) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Unable to read input file"+filepath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse CSV file")
	}

	return records

}

func LoadStorage() ([]Task, error) {

	csv := readCSV("todolite.csv")

	var tasks []Task
	for elem := range csv {
		id, err := strconv.Atoi(csv[elem][0])

		if err != nil {
			log.Fatal("Unable to get task id")
			return []Task{}, err

		}

		name := csv[elem][1]
		status, err := strconv.ParseBool(csv[elem][2])

		if err != nil {
			log.Fatal("Failed to get task status")
			return []Task{}, err
		}

		creationdate, err := time.Parse(time.DateOnly, csv[elem][3])
		if err != nil {
			log.Fatal("Unable to get task creation date")
			return []Task{}, err
		}

		tasks = append(tasks, Task{
			Id:           id,
			Name:         name,
			Status:       status,
			Creationdate: creationdate,
		})
	}
	return tasks, nil
}

func AppendTaskToCSV(task Task) error {
	file, err := os.OpenFile("todolite.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	row := []string{
		strconv.Itoa(task.Id),
		task.Name,
		strconv.FormatBool(task.Status),
		task.Creationdate.Format(time.DateOnly),
	}

	if err := writer.Write(row); err != nil {
		return err
	}

	writer.Flush()

	return writer.Error()
}

func SaveStorage(filepath string, tasks []Task) error {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Unable to read input file"+filepath, err)
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	for _, task := range tasks {
		record := []string{strconv.Itoa(task.Id), task.Name, strconv.FormatBool(task.Status), task.Creationdate.Format(time.DateOnly)}
		if err := csvWriter.Write(record); err != nil {
			log.Fatal("Unable to write to a file")
			return err
		}
	}

	return nil
}
