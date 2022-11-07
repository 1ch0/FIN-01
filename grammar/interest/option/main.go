package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
}

type option func(*User)

func WithId(id int) option {
	return func(u *User) { u.Id = id }
}

func WithName(name string) option {
	return func(u *User) { u.Name = name }
}

func WithEmail(email string) option {
	return func(u *User) { u.Email = email }
}

func NewUser(opt ...option) *User {
	const (
		defaultId    = -1
		defaultName  = "guest"
		defaultEmail = "guest@example.com"
	)
	u := &User{
		Id:    defaultId,
		Name:  defaultName,
		Email: defaultEmail,
	}

	for _, opt := range opt {
		opt(u)
	}
	return u
}

func main() {
	u1 := NewUser(WithName("test"), WithId(3155))
	fmt.Printf("%+v\n", u1)
	u2 := NewUser(WithName("email"), WithEmail("a@asdgg.com"))
	fmt.Printf("%+v\n", u2)
	u3 := NewUser()
	fmt.Printf("%+v\n", u3)
}
