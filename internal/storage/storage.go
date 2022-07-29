package storage

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
	"time"
)

var data map[uint]*Task

var TaskNotExists = errors.New("Task does not exist")

func init() {
	data = make(map[uint]*Task)
	if _, err := Add("Please, add some tasks :)"); err != nil {
		log.Panic(err)
	}
}

func List() []*Task {
	res := make([]*Task, 0, len(data))
	for _, v := range data {
		res = append(res, v)
	}
	return res
}

func Add(title string) (*Task, error) {
	lastId++
	var task = &Task{
		id:          lastId,
		title:       title,
		isCompleted: false,
		createdAt:   time.Now(),
	}
	data[task.GetId()] = task
	return task, nil
}

func Complete(id uint) error {
	if _, ok := data[id]; !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(id), 10))
	}
	data[id].isCompleted = true
	data[id].SetCompletedAt(time.Now())
	return nil
}

func Get(id uint) (*Task, error) {
	if _, ok := data[id]; !ok {
		return nil, errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(id), 10))
	}
	return data[id], nil
}

func Update(id uint, title string) error {
	if _, ok := data[id]; !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(id), 10))
	}
	data[id].title = title
	return nil
}

func Delete(id uint) error {
	if _, ok := data[id]; !ok {
		return errors.Wrap(TaskNotExists, strconv.FormatUint(uint64(id), 10))
	}
	delete(data, id)
	return nil
}
