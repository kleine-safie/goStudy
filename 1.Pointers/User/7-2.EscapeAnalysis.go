package main

type UserProfile2 struct {
	Name   string
	Status string
}

func createUser2() UserProfile2 {
	user := UserProfile2{Name: "Albert", Status: "Active"}

	updateStatus2(&user, "Inactive")
	return user
}

func updateStatus2(u *UserProfile2, newStatus string) {
	u.Status = newStatus
}

func main() {

	createUser2()
}