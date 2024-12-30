package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

// Student struct embedding Person
type Student struct {
	Person
	studentID uint64
}

func (s *Student) GetStudentID() uint64 {
	return s.studentID
}

func (s *Student) SetStudentID(id uint64) {
	s.studentID = id
}

// Teacher struct to demonstrate Polymorphism
type Teacher struct {
	Person
	subject string
}

func (t *Teacher) GetSubject() string {
	return t.subject
}

func (t *Teacher) SetSubject(subject string) {
	t.subject = subject
}

// Human interface to demonstrate Polymorphism
type Human interface {
	GetName() string
}

// PrintName demonstrates polymorphism by treating both Student and Teacher as Human
func PrintName(h Human) {
	fmt.Println("Name:", h.GetName())
}

func main() {
	// Create a Student object
	student := &Student{}
	student.SetName("Alice")
	student.SetStudentID(12345)

	// Create a Teacher object
	teacher := &Teacher{}
	teacher.SetName("Mr. Bob")
	teacher.SetSubject("Mathematics")

	// Demonstrate Encapsulation and Abstraction
	fmt.Printf("Student Name: %s, ID: %d\n", student.GetName(), student.GetStudentID())
	fmt.Printf("Teacher Name: %s, Subject: %s\n", teacher.GetName(), teacher.GetSubject())

	// Demonstrate Polymorphism
	PrintName(student)
	PrintName(teacher)
}
