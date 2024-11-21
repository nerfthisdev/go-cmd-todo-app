package csvmodule

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id           int
	Name         string
	Status       string
	Creationdate time.Time
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

func LoadStorage(filename string) ([]Task, error) {

	const layout = "2006-Jan-02"

	csv := readCSV(filename)

	var tasks []Task
	for elem := range csv {
		id, errc := strconv.Atoi(csv[elem][0])

		if errc != nil {
			log.Fatal("Unable to get task id")
			return []Task{}, errc

		}

		name := csv[elem][1]
		status := csv[elem][2]

		creationdate, err := time.Parse(layout, csv[elem][3])
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
