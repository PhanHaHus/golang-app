package db


//CreateUser will create a new user, take as input the parameters and
//insert it into database
func CreateUser(username, password, email string) string {
	err := "err"
	return err
}

//ValidUser will check if the user exists in db and if exists if the username password
//combination is valid
func ValidUser(username, password string) bool {

	//by default return false
	return false
}
