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

/*
go build -gcflags "-m -l" study
go build -gcflags "-m=2 -l" study
https://github.com/golang/go/blob/master/src/cmd/compile/internal/escape/escape.go


(1) pointers to stack objects cannot be stored in the heap
(2) pointers to a stack object cannot outlive that object (
   e.g., because the declaring function returned and destroyed the object's stack frame, or its space is reused across loop iterations for logically distinct variables).

*/
