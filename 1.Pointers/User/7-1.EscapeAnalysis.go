package main

type UserProfile struct {
	Name   string
	Status string
}

func createUser() *UserProfile {
	user := UserProfile{Name: "Albert", Status: "Active"}

	updateStatus(&user, "Inactive")
	return &user
}

func updateStatus(u *UserProfile, newStatus string) {
	u.Status = newStatus
}

//func main() {

//	createUser()
//}
