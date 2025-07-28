package user

import "time"

func CreateUser(u *User) error {
	u.CreatedAt = time.Now()
	return CreateUserInRepo(u)
}

func FinishUser(id uint) error {
	user, err := FindByID(id)
	if err != nil {
		return err
	}
	now := time.Now()
	user.UpdatedAt = now
	return UpdateUser(&user)
}
