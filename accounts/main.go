package main

import (
	"fmt"
	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("KL")
	fmt.Println(account)
	//Doesm't work this way.
	//account.balance = 1000
	//We can make this method than we can use it.
	account.Deposit(2000)
	/*
		fmt.Println(account.Balance())
		err := account.Withdraw(2020)
		if err != nil {
			//log.Fatalln(err) //에러시 프로그램을 강제 종료시킴
			fmt.Println(err)
		}
	*/
	//fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account.String())

}
