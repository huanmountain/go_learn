package model

import (
	"fmt"
)

//声明一个customer结构体,表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string

}

func NewCustomer(id int ,name string,cender string ,age int,phone string,email string) Customer {

	return Customer{
		Id:     id,
		Name:   name,
		Gender: cender,
		Age:    age,
		Phone:  phone,
		Email:  email,

	}

}
func NewCustoner2(name string,cender string ,age int,phone string,email string)  Customer {
	return Customer{
		Name:   name,
		Gender: cender,
		Age:    age,
		Phone:  phone,
		Email:  email,

	}

}

func (this Customer) GetInfo()string  {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email)
	return  info
}

func (this Customer)GetName()string{
	return this.Name
}
func (this Customer)GetAge()int{
	return this.Age
}
func (this Customer)GetPhone()string{
	return this.Phone
}
func (this Customer)GetEmile()string{
	return this.Email
}

func (this Customer)GetGander()string{
	return this.Gender
}