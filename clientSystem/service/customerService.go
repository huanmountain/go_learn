package service

import (
	"fmt"
	_ "strconv"

	"jim/clientSystem/model"
)

//方法,customer的增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段,表示切片含有客户数量(可以作为客户id)
	customerNum int

}
//一个方法返回customerservice
func NewcustomerService() *CustomerService  {

	customerService := &CustomerService{}
	customerService.customerNum =1
	customer := model.Customer{1,"芹芹","男",20,"1806923","out@look"}
	customerService.customers = append(customerService.customers,customer)
	return customerService
}

//返回客户切片
func(this *CustomerService) List() []model.Customer{
	return this.customers
}


//添加客户
func (this *CustomerService)  Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers =append(this.customers,customer)
	return true
}
//delete
func (this *CustomerService) Delete(id int)bool {
	index := this.FindById(id)
	if index ==-1 {
		return false
	}
	//this.customers[index+1:]是切片,在末尾加上...
	this.customers =append(this.customers[0:index],this.customers[index+1:]...)
	return true
}


func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customer
	for i:=0;i<len(this.customers);i++{
		if this.customers[i].Id == id {
			index =i
		}

	}
	return index
}

//update
func  (this *CustomerService) Update(index int) {
	//var updateName string;
	//var updateGander string
	//var updateAge int
	//var updatePhone string
	//var updateEmil string
	var (
		updateName  string
		updateGander string
		updateAge int
		updatePhone string
		updateEmil string

	)
	fmt.Print("姓名:"+"("+this.customers[index].GetName()+")")
	fmt.Scanln(&updateName)

	fmt.Print("性别:"+"(",this.customers[index].GetGander(),")")
	fmt.Scanln(&updateGander)
	fmt.Print("年龄:"+"(",this.customers[index].GetAge(),")")
	fmt.Scanln(&updateAge)
	fmt.Print("电话号码:"+"(",this.customers[index].GetPhone(),")")
	fmt.Scanln(&updatePhone)
	fmt.Print("邮箱:"+"(",this.customers[index].GetEmile(),")")
	fmt.Scanln(&updateEmil)

	if updateAge == 0 {
		updateAge =this.customers[index].GetAge()
	}
	if updateName == "" {
		updateName =this.customers[index].GetName()
	}
	if updateGander == "" {
		updateGander =this.customers[index].GetGander()
	}
	if updatePhone == "" {
		updatePhone =this.customers[index].GetPhone()
	}
	if updateEmil == "" {
		updateEmil =this.customers[index].GetEmile()
	}


	customer := model.NewCustomer(index+1,
		updateName,updateGander,updateAge,updatePhone,
		updateEmil)
	this.customers[index] = customer
	fmt.Println("-----------修改成功------------")

}