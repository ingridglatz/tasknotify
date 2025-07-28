package task

import "time"

func CreateTask(t *Task) error {
	t.Status = "pendente"
	t.CreatedAt = time.Now()
	return Create(t)
}

func FinishTask(id uint) error {
	task, err := FindByID(id)
	if err != nil {
		return err
	}
	now := time.Now()
	task.Status = "concluida"
	task.FinishedAt = &now
	return Update(&task)
}
