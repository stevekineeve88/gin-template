package objects

type User struct {
	Id        int64
	FirstName string
	LastName  string
}

func (u *User) GetMap() map[string]any {
	userMap := make(map[string]any)
	userMap["id"] = u.Id
	userMap["first_name"] = u.FirstName
	userMap["last_name"] = u.LastName
	return userMap
}
