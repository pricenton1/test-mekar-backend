package model

type Account struct {
	AccountID string `json:"accountID"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     Token  `json:"token"`
}

type Token struct {
	Key string `json:"token"`
}
