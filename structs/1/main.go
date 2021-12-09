package main

import "fmt"

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

type employee struct {
	name   string
	sex    string
	age    int
	salary int
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник: %s\nПол: %s\nВозраст: %d\nЗарплата: %d\n", e.name, e.sex, e.age, e.salary)
}

func (e *employee) increaseSalary(amount int) {
	e.salary += amount
}

func main() {
	e1 := newEmployee("Ирка", "Ж", 31, 30)
	e2 := newEmployee("Валерка", "М", 31, 5000)

	e1.increaseSalary(30)

	fmt.Println(e1.getInfo())
	fmt.Println(e2.getInfo())

}
