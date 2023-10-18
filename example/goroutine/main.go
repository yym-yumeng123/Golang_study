package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type User struct {
	Name     string
	Phone    string
	Position string
}

func GetBaseUserInfo(id string) (*User, error) {
	time.Sleep(2 * time.Second)
	return &User{
		Name:  "yym",
		Phone: "1221",
	}, nil
}

func GetUserPosition(id string) (string, error) {
	time.Sleep(2 * time.Second)
	return "developer", nil
}

func GetUser(id string) (*User, error) {
	var baseUserInfoErr error
	var positionErr error
	var user *User
	var position string
	var wg sync.WaitGroup

	wg.Add(1)
	// 启用 goroutine
	go func() {
		// goroutine body
		user, baseUserInfoErr = GetBaseUserInfo(id)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		position, positionErr = GetUserPosition(id)
		wg.Done()
	}()

	wg.Wait()

	if baseUserInfoErr != nil {
		return nil, baseUserInfoErr
	}

	if positionErr != nil {
		return nil, positionErr
	}

	user.Position = position
	return user, nil
}

func main() {
	now := time.Now()
	user, err := GetUser("1")
	if err != nil {
		log.Fatalf("GetUser failed, %v", err)
	}
	fmt.Println(time.Now().Sub(now).Seconds())
	fmt.Println(user)
}
