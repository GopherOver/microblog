package main

type User struct {
	ID                 string
	Username, Password string
	IsAdmin            bool
}

func (u *User) Login() {

}

func (u *User) Logout() {

}

// func (u *User) Auth() {

// }
