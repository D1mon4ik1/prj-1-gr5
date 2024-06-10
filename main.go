package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"project-go/domain"
	"sort"
	"strconv"
	"time"
)

var id uint64 = 1

const (
	totalPoints      = 5
	pointPerQuestion = 5
)

func main() {
	fmt.Println("Вітаю у грі HARGCORE-MATH!")
	time.Sleep(2 * time.Second)

	users := getUsers()
	for _, user := range users {
		if user.Id >= id {
			id = user.Id + 1
		}
	}

	for {
		menu()
		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			users = getUsers()
			users = append(users, u)
			sortAndSave(users)
		case "2":
			users = getUsers()
			for _, user := range users {
				fmt.Printf("Id: %v, Name: %s, Time: %v\n",
					user.Id, user.NickName, user.Time)
			}
		case "3":
			return
		}
	}
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	for i := 5; i > 0; i-- {
		fmt.Printf("До початку: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	points := totalPoints
	myPoints := 0
	start := time.Now()
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y

		fmt.Println(x, "+", y, "=")

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if res == ansInt {
				myPoints += pointPerQuestion
				points -= pointPerQuestion
				fmt.Printf(
					"Чудово! Ви набрали %v очок, залишилось %v\n",
					myPoints, points,
				)
			} else {
				fmt.Println("Спробуй ще!")
			}
		}
	}

	finish := time.Now()
	timeSpent := finish.Sub(start)

	fmt.Printf("Вітаю, ти впорався за %v!", timeSpent.Seconds())

	fmt.Println("Введіть свій нікнейм:")
	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Id:       id,
		NickName: name,
		Time:     timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile(
		"users.json",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}

func getUsers() []domain.User {
	info, err := os.Stat("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create("users.json")
			if err != nil {
				fmt.Printf("Error: %s", err)
				return nil
			}
			return nil
		}
	}

	var users []domain.User
	if info.Size() != 0 {
		file, err := os.Open("users.json")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}

		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}(file)

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&users)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}
	}
	return users
}
