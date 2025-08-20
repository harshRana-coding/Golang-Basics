package student

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)
type student struct {
	firstName   string
	lastName    string
	dateOfBirth string
	age         int64
	createdAt   time.Time
}

type monitor struct {
	class int32
	marksPercentage float32
	student
}

func newMonitor(class int32, marksPercentage float32)*monitor {
	return &monitor{
		class: class,
		marksPercentage: marksPercentage,
		student: student{
			firstName: "ABC",
			lastName: "BVC",
			dateOfBirth: "---",
			createdAt: time.Now(),
		},
	}
}

func newStudent(firstName, lastName, dateOfBirth string , age int64) (*student,error) {

	if firstName == "" || lastName == "" || dateOfBirth == "" {
		return nil, errors.New("First name, last name and birthdate are required. ")
	}

	return &student{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dateOfBirth,
		age:         age,
		createdAt:   time.Now(), 
	}, nil
}

func (s *student) outputUserDetails() {
	fmt.Println(s.firstName, s.lastName, s.dateOfBirth, s.age)
}

func (s *student) clearUserName() {
	s.firstName = ""
	s.lastName = ""
}

func Init() {
	firstName := getUserData("Enter your first name : ")
	lastName := getUserData("Enter your last name : ")
	dateOfBirth := getUserData("Enter your DOB [DD/MM/YYYY] : ")
	age, _ := strconv.ParseInt(getUserData("Enter your age : "), 10, 64)

	var appUser *student

	appUser,err := newStudent(firstName, lastName, dateOfBirth, age);
	if err != nil {
		fmt.Println(err);
		return 
	}

	monitor := newMonitor(11,95.6)
	monitor.outputUserDetails()

	appUser.outputUserDetails();
	appUser.clearUserName();
	appUser.outputUserDetails();
}


func getUserData(textString string) string {
	fmt.Print(textString)
	var value string
	fmt.Scanln(&value)
	return value
}
