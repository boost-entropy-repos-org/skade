package engine

import (

)

//CreateUser creates a user 
func (a *analysisService) CreateUser(username string, email string) error {
	a.repo.CreateUser(username, email)
}