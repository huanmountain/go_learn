package main

import (
	"fmt"

	"jim/clientSystem/model"
	"jim/clientSystem/service"
)

type customerView struct {
	key string
	loop bool
	//增加一个字段customerService
	customerService *service.CustomerService

}

//显示所有客户信息
func (this customerView) List()  {

	//获取所有客户信息
	customers := this.customerService.List()
	fmt.Println("-----------------------客户列表-----------------------")
	fmt.Println("编号\t姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i:=0;i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}

	fmt.Println("---------------------客户列表完成----------------------")

}
//得到输入,添加客户信息
func (this *customerView) Add()  {


	var (
		name = ""
		gender = ""
		age = 0
		phone = ""
		email = ""

	)
	fmt.Println("-----------------------客户列表-----------------------")
	fmt.Println("姓名:")
	fmt.Scanln(&name)
	fmt.Println("性别:")
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	fmt.Scanln(&age)
	fmt.Println("电话:")
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	fmt.Scanln(&email)
	customer := model.NewCustoner2(name,gender,age,phone,email)
	if this.customerService.Add(customer){
		fmt.Println("-----------------------添加成功-----------------------")
	}else {
		fmt.Println("-----------------------添加失败-----------------------")
	}

}
//得到输入,修改客户信息
func (this *customerView)Update(){
	index := 0
	fmt.Print("--------------------修 改 用 户--------------------\n")
	fmt.Print("请输入要修改客户的id(-1退出):")
	id := 0
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	for{
		choice := id
		index = this.customerService.FindById(choice)
		if id==-1{
			fmt.Print("没有该id号的用户,修改失败，请重新输入：")
			fmt.Scanln(&choice)
		}else{
			break
		}
	}
	this.customerService.Update(index)
}





func (this *customerView) Del()  {
	fmt.Println("-----------------------删除客户-----------------------")
	fmt.Println("请选择待删除客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	//flag := 0
	//for flag == 0 {
	//fmt.Println("确认是否删除(Y/N):")
	//choise := ""
	//fmt.Scanln(&choise)
	//
	//
	//	if choise == "Y" || choise== "y"{
	//		flag =1
	//		if this.customerService.Delete(id){
	//			fmt.Println("-----------------------删除完成-----------------------")
	//		}else {
	//			fmt.Println("-----------------------删除失败-----------------------")
	//		}
	//	}else if choise == "N" || choise== "n" {
	//		flag = 1
	//		fmt.Println("-----------------------取消删除-----------------------")
	//	}
	//}
	choise := ""
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&choise)
		if choise == "Y"|| choise == "N"|| choise == "y"|| choise == "n"{
			break
		}
		fmt.Println("你的输入有误,请确认是否退出(Y/N):")
	}
	if choise == "Y" || choise== "y"{
		if this.customerService.Delete(id){
			fmt.Println("-----------------------删除完成-----------------------")
		} else{
			fmt.Println("-----------------------删除失败-----------------------")
		}
	}else if choise == "N" || choise== "n" {
			fmt.Println("-----------------------取消删除-----------------------")
	}


}


func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y"|| this.key == "N"|| this.key == "y"|| this.key == "n"{
			break
		}
		fmt.Println("你的输入有误,请确认是否退出(Y/N):")
	}
	if this.key == "Y"|| this.key == "y"{
		this.loop =false
	}
}
//显示主菜单
func (this *customerView) mainMenu() {
	for  {
		fmt.Println("\n-----------------------客户信息管理系统-----------------------")
		fmt.Println("                          1.添加客户")
		fmt.Println("                          2.修改客户")
		fmt.Println("                          3.删除客户")
		fmt.Println("                          4.客户列表")
		fmt.Println("                          5.退出")
		fmt.Print("请选择(1-5):")



		fmt.Scanln(&this.key)
		switch this.key {
		case "1":this.Add()
		case "2":this.Update()
		case "3":this.Del()
		case "4":this.List()
		case "5":this.exit()
		default:
			fmt.Println("你的输入有误,请输入正确选项")
		}
		if !this.loop{
			break
		}
	}
	fmt.Println("你退出了客户管理系统")



}


func main()  {
	customerView := customerView{
		key: "",
		loop: true,
	}
	//完成对customerServie字段的初始化
	customerView.customerService = service.NewcustomerService()
	customerView.mainMenu()
}