package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(taskID int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(taskID int, task *model.Task) error {
	// TODO: replace this
	err := t.filebased.UpdateTask(taskID, *task)
	return err
}

func (t *taskRepository) Delete(id int) error {
	// TODO: replace this
	err := t.filebased.DeleteTask(id)
	return err
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	 // TODO: replace this
	 task, err := t.filebased.GetTaskByID(id)
	 return task, err
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	// TODO: replace this
	tasks, err := t.filebased.GetTasks()
	return tasks, err
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	// TODO: replace this
	tc, err := t.filebased.GetTaskListByCategory(id)
	return tc, err
}
