package user

import "tasknotify/config"

func CreateUserInRepo(user *User) error {
	return config.DB.Create(user).Error
}

func FindAll() ([]User, error) {
	var users []User
	err := config.DB.Find(&users).Error
	return users, err
}

func FindByID(id uint) (User, error) {
	var user User
	err := config.DB.First(&user, id).Error
	return user, err
}

func UpdateUser(user *User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return config.DB.Delete(&User{}, id).Error
}
