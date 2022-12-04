package handlerUsers

type User struct {
	UserId      int64  `json:"user id"`
	UserName    string `json:"user name"`
	UserSurname string `json:"user surname"`
}
