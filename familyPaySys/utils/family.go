package utils

import (
	"fmt"
)


type familyAccount struct {
	//声明一个变量key记录输入选项
	key string
	//声明一个变量是否退出for
	loop bool
	balance float64
	//定义变化金额
	money float64
	//note
	note string
	flag bool
	//"收支详情"
	//details :="收支\t账户余额\t收支金额\t说明"
	details string
}


//编写一个工厂模式
func NewFamiluAccount() *familyAccount {

	return &familyAccount{
		key: "",
		loop: true,
		balance: 1000.0,
		money: 0.0,
		note: "",
		flag: false,
		details: "收支\t账户金额\t收支金额\t说    明",
	}
}




//将显示明细封装到一个方法中
func (this *familyAccount) showDetails()  {
	fmt.Println("-----------------------当前收支明细-----------------------")
	if this.flag{
		fmt.Println(this.details)
	}else {
		fmt.Println("当前没有收支情况")
	}

}
//收入
func (this *familyAccount) income() {
	fmt.Println("本次收入金额:")
	_, _ = fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)

	//存储到details中
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",this.balance,this.money,this.note)
	this.flag =true

}
//支出
func (this *familyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	//判断余额
	if this.money >this.balance{
		fmt.Println("余额不足")
	}

	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag =true
}
func (this *familyAccount) exit()  {
	fmt.Println("你确定要退出吗?y/n")
	choice :=""
	for{
		fmt.Scanln(&choice)
		if choice =="y"||choice=="n"{
			break
		}
		fmt.Println("输入有误,请重新输入")
	}
	if choice == "y"{
		this.loop =false
	}


}
//显示主菜单
func (this *familyAccount) MAinMenu()  {
	for  {
		fmt.Println("\n-----------------------家庭收支明细-----------------------")
		fmt.Println("                          1.收支明细")
		fmt.Println("                          2.登记收入")
		fmt.Println("                          3.登记支出")
		fmt.Println("                          4.退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":this.showDetails()
		case "2":this.income()
		case "3":this.pay()
		case "4":this.exit()
		default:
			fmt.Println("请输入正确选项")
		}
		if !this.loop{
			break
		}



	}
}