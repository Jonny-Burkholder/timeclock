package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type shift struct {
	Start       time.Time     `json:"start_time"`
	End         time.Time     `json:"end_time"`
	ShiftLength time.Duration `json:"shift_length"`
}

//user is where we'll store useful information
type User struct {
	Username string
	Shift    *shift
}

//userMap stores the hashed version of each user's pin as a key and their userID as the value
type userMap struct {
	Users map[string]*User `json:"users"`
}

type AnyPage interface {
	Save() error
}

type Page struct{}

func (p *Page) Save() error {
	return nil
}

func (user *User) Save() error {
	filepath := "../internal/storage/users/" + user.Username + ".json"
	file, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

//NewUserMap is exactly what it says on the tin
func NewUserMap() *userMap {
	return &userMap{
		Users: make(map[string]*User),
	}
}

//Load allows us to retrieve the usermap from disk into memory
func (u *userMap) Load() error {
	filepath := "../internal/storage/users/usermap.json"

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("Error! No user map file found")
	}

	err = json.Unmarshal(file, u)

	if err != nil {
		return err
	}

	return nil
}

func (u *userMap) Save() error {
	filepath := "../internal/storage/users/usermap.json"

	file, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, file, 0644)
}

//Hash takes a string input and returns a hash of that string
func (u *userMap) Hash(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), 5)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(hash), nil
}

//AddUser allows us to put a new user into the map
//I realize that it's terrible, but I'm going for expedience on this one
func (u *userMap) AddUser(username, pin string) error {
	/*
		hash, err := u.Hash(pin)
		if err != nil {
			return err
		}
	*/
	_, ok := u.Users[pin]
	if ok != false {
		return fmt.Errorf("Choose a new pin")
	}

	user := User{
		Username: username,
		Shift:    &shift{},
	}

	u.Users[pin] = &user

	err := user.Save()
	if err != nil {
		return err
	}
	return nil
}

//LoadUser takes a user json and unmarshals it into a user struct
func (u *userMap) LoadUser(name string) (*User, error) {
	filepath := "../internal/storage/users/" + name + ".json"
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return &User{}, err
	}
	user := &User{}
	err = json.Unmarshal(file, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

//CheckPin is a super non-secure way to retrieve a user based on a pin. Like, don't do this for real
func (u *userMap) CheckPin(pin string) (*User, error) {
	/*hash, err := u.Hash(pin)
	if err != nil {
		return &User{}, err
	}*/

	us, ok := u.Users[pin]
	if ok != true {
		return &User{}, fmt.Errorf("User not found")
	}

	return us, nil

}

func (u *User) StartShift() {
	u.Shift.Start = time.Now()
	fmt.Println(u.Shift.Start)
	u.Save()
}

func (u *User) EndShift() {
	u.Shift.End = time.Now()
	fmt.Println(u.Shift.End)
	u.Shift.ShiftLength = u.Shift.End.Sub(u.Shift.Start)
	fmt.Println(u.Shift.ShiftLength)
	u.Save()
}

func DisplayTime(time time.Time) string {
	fmt.Println(time)
	a, b, c := time.Clock()
	return strconv.Itoa(a) + ":" + strconv.Itoa(b) + ":" + strconv.Itoa(c)
}

func DisplayShift(d time.Duration) string {
	fmt.Println(d)
	return strconv.Itoa(int(d.Hours())) + ":" + strconv.Itoa(int(d.Minutes())) + ":" + strconv.Itoa(int(d.Seconds()))
}
