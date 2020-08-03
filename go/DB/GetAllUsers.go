package GetAllUsers

import (
	"database/sql"
	"fmt"
)

type Login struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
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
		id    int
		email string
		name  string
		pass  string
		login []Login
	)
	i := 0

	rows, err := db.Query("select * from users")
	for rows.Next() {
		err := rows.Scan(&id, &email, &name, &pass)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}
		login = append(login, Login{Email: email, Name: name, Pass: pass})

	}
	defer rows.Close()
	return login
}
