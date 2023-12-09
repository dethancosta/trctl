package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/muesli/go-app-paths"
)

var (
	DefaultServerUrl = "http://localhost:6576"
	configPath       = GetConfigPath()
)

func RemovePid() error {
	config := GetConfig()

	if config == nil {
		return errors.New("Could not open config file")
	}
	pidStr, ok := config["pid"]
	if !ok {
		return errors.New("No pid found")
	}
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return fmt.Errorf("Couldn't get pid from file: %w", err)
	}
	err = syscall.Kill(pid, syscall.SIGINT)
	if err != nil {
		return fmt.Errorf("Couldn't send kill signal: %w", err)
	}
	delete(config, "pid")
	fileContents, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("Error while marshalling new contents: %w", err)
	}
	configFile, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("Error while opening config file: %w", err)
	}
	_, err = configFile.Write(fileContents)

	return err
}

// TODO refactor to return (map[string]string, error)
func GetConfig() map[string]string {
	config := make(map[string]string)

	if configPath == "" {
		return config
	}

	file, err := os.Open(configPath)
	if err != nil {
		return config
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return config
	}

	_ = json.Unmarshal(fileContents, &config)
	return config
}

func GetConfigPath() string {
	configPath, err := gap.NewScope(gap.User, "timeruler").ConfigPath("config.json")
	if err != nil {
		return ""
	}

	return configPath
}

func GetConfigDir() string {
	configDir, err := gap.NewScope(gap.User, "timeruler").ConfigPath("")
	if err != nil {
		return ""
	}

	return configDir
}

type Task struct {
	Description string    `json:"Description"`
	StartTime   time.Time `json:"Start"`
	EndTime     time.Time `json:"End"`
	Tag         string    `json:"Tag"`
}

func TasksFromCsv(filename string) ([]Task, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Can't open file: %w", err)
	}
	r := csv.NewReader(f)
	taskList := []Task{}
	lc := 0
	var desc string
	var start time.Time
	var end time.Time
	var tag string

	line, err := r.Read()
	for err != io.EOF {
		if err != nil {
			return nil, fmt.Errorf("Error parsing file: %w", err)
		}
		lc++
		if len(line) < 3 {
			return nil,
				errors.New("Error reading file: Field missing from line " + strconv.Itoa(lc))
		}
		desc = line[0]

		// Set task times to the current day (for now)
		now := time.Now()
		start, err = time.Parse(time.TimeOnly, line[1])
		if err != nil {
			return nil, errors.New("Error reading file: time value improperly formatted on line " + strconv.Itoa(lc))
		}
		start = time.Date(now.Year(), now.Month(), now.Day(), start.Hour(), start.Minute(), 0, 0, time.Local)
		end, err = time.Parse(time.TimeOnly, line[2])
		if err != nil {
			return nil, errors.New("Error reading file: time value improperly formatted on line " + strconv.Itoa(lc))
		}

		end = time.Date(now.Year(), now.Month(), now.Day(), end.Hour(), end.Minute(), 0, 0, time.Local)
		tag = strings.TrimSpace(line[3])
		var task Task
		if len(tag) > 0 {
			task = Task{
				Description: desc,
				StartTime:   start,
				EndTime:     end,
				Tag:         tag,
			}
		}
		if task.Description == "" {
			return nil, errors.New("Error creating tasks: Task could not be created on line " + strconv.Itoa(lc))
		}

		taskList = append(taskList, task)
		line, err = r.Read()
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("Error parsing file: %w", err)
		}
	}
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("Error parsing file: %w", err)
	}

	return taskList, nil
}
