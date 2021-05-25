package custom_type

import "fmt"

type user struct {
	name  string
	phone string
}

func Create() {
	u := user{name: "deniz", phone: "555"}
	u2 := new(user)
	u2.name = "deniz"
	u2.phone = "554"
	fmt.Println(u, u2)
}