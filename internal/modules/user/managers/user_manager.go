package managers

import (
	"fmt"
	"user/data"
	"user/objects"
)

var manager *UserManager

func GetUserManager() UserManager {
	if manager == nil {
		manager = &UserManager{
			userData: data.GetUserData(),
		}
	}
	return *manager
}

type UserManager struct {
	userData data.UserData
}

func (manager *UserManager) Create(firstName string, lastName string) (*objects.User, error) {
	result := manager.userData.Create(firstName, lastName)
	if result.Status == false {
		return nil, fmt.Errorf(fmt.Sprintf("could not create user: %s", result.Message))
	}
	return &objects.User{
		Id:        result.InsertId,
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func (manager *UserManager) GetAll() ([]*objects.User, error) {
	result := manager.userData.LoadAll()
	if result.Status == false {
		return nil, fmt.Errorf(fmt.Sprintf("could not fetch users: %s", result.Message))
	}

	defer func() {
		err := result.Data.Close()
		if err != nil {
			panic(err)
		}
	}()

	var users []*objects.User
	for result.Data.Next() {
		var user objects.User
		err := result.Data.Scan(&user.Id, &user.FirstName, &user.LastName)
		if err != nil {
			panic(err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (manager *UserManager) Update(id int64, firstName string, lastName string) error {
	result := manager.userData.Update(id, firstName, lastName)
	if result.Status == false || result.AffectedRows == 0 {
		return fmt.Errorf("could not update user")
	}

	return nil
}

func (manager *UserManager) Delete(id int64) error {
	result := manager.userData.Delete(id)
	if result.Status == false || result.AffectedRows == 0 {
		return fmt.Errorf("could not delete user")
	}

	return nil
}
