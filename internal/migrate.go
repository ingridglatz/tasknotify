package migrate

import (
	"tasknotify/config"
	"tasknotify/internal/task"
	"tasknotify/internal/user"
)

func AutoMigrate() {
	config.DB.AutoMigrate(
		&task.Task{},
		&user.User{},
	)
}
