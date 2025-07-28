package task

import "tasknotify/config"

func Create(task *Task) error {
	return config.DB.Create(task).Error
}

func FindAll() ([]Task, error) {
	var tasks []Task
	err := config.DB.Find(&tasks).Error
	return tasks, err
}

func FindByID(id uint) (Task, error) {
	var task Task
	err := config.DB.First(&task, id).Error
	return task, err
}

func Update(task *Task) error {
	return config.DB.Save(task).Error
}

func Delete(id uint) error {
	return config.DB.Delete(&Task{}, id).Error
}
