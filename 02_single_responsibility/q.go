package main

import "fmt"

/*
 * 以下はユーザーを作成して、データベースに保存し、ウェルカムメールを送信するコードです。
 *
 * このコードに問題はありますか？（ヒント：責務）
 * どのように改善できますか？
 */

type UserManager struct {
	UserName string
	Email    string
}

func (um *UserManager) SaveUser() {
	// データベースにユーザー情報を保存する処理
	fmt.Printf("Save user %s to database.\n", um.UserName)
}

func (um *UserManager) SendWelcomeEmail() {
	// ユーザーにウェルカムメールを送信する処理
	fmt.Printf("Sending welcome email to %s.\n", um.Email)
}

func (um *UserManager) DisplayUserInfo() {
	// ユーザー情報を表示する処理
	fmt.Printf("User: %s, Email: %s\n", um.UserName, um.Email)
}

func main() {
	userManager := UserManager{
		UserName: "JohnDoe",
		Email:    "john@example.com",
	}
	userManager.SaveUser()
	userManager.SendWelcomeEmail()
	userManager.DisplayUserInfo()
}
