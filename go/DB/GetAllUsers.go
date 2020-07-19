package GetAllUsers

import (
	"database/sql"
	"fmt"
)

type Login struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

func GetAllUsers() []Login {
	//opening database
	fmt.Println("db begin")
	db, err := sql.Open("mysql", "root:@/user")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("open ")
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping ")
	}

	var (
		id       int
		email    string
		username string
		password string
		login    []Login
	)
	i := 0

	rows, err := db.Query("select * from users")
	for rows.Next() {
		err := rows.Scan(&id, &username, &email, &password)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}
		login = append(login, Login{Username: username, Email: email, Password: password})

	}
	defer rows.Close()
	return login
}
