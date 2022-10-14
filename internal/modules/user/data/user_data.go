package data

import (
	"fmt"
	"util/managers"
	"util/objects"
)

var data *UserData

func GetUserData() UserData {
	if data == nil {
		data = &UserData{
			connectionManager: managers.GetConnectionManager(),
		}
	}
	return *data
}

type UserData struct {
	connectionManager managers.ConnectionManager
}

func (data *UserData) Create(firstName string, lastName string) objects.DatabaseResult {
	return data.connectionManager.Insert(fmt.Sprintf(
		"INSERT INTO users (first_name, last_name) VALUES ('%s', '%s')",
		firstName,
		lastName,
	))
}

func (data *UserData) LoadAll() objects.DatabaseResult {
	return data.connectionManager.Select("SELECT * FROM users")
}

func (data *UserData) Update(id int64, firstName string, lastName string) objects.DatabaseResult {
	return data.connectionManager.Query(fmt.Sprintf(
		"UPDATE users SET first_name = '%s', last_name = '%s' WHERE id = %d",
		firstName,
		lastName,
		id,
	))
}

func (data *UserData) Delete(id int64) objects.DatabaseResult {
	return data.connectionManager.Query(fmt.Sprintf(
		"DELETE FROM users WHERE id = %d",
		id,
	))
}
