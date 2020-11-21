package entities

type DbCreateUserParams struct {
	Username	string
	Email		string
}

type User struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
}