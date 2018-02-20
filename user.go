package main

type User struct {
	ID        int64  `form:"id" db:"id"`
	Username  string `form:"username" db:"username"`
	Password  string `form:"password" db:"password"`
	Password2 string `form:"password2" db:"-"`
	IsAdmin   bool   `form:"-" db:"is_admin"`
}

func (u *User) Login() {

}

func (u *User) Logout() {

}

// func (u *User) Auth() {

// }
