package main

import "fmt"

/*
 * 問題点： UserManager が複数の責務を持ちすぎて、メンテンス性が低い
 * 改善方法： 単一責務の原則に従い、変更する理由が同じものは集める、変更する理由が異なるものは分ける
 */

type User struct {
	Name  string
	Email string
}

func (u *User) Display() {
	// ユーザー情報を表示する処理
	fmt.Printf("User: %s, Email: %s\n", u.Name, u.Email)
}

type UserRepository struct{}

func (ur *UserRepository) Save(user *User) {
	// データベースにユーザー情報を保存する処理
	fmt.Printf("Save user %s to database.\n", user.Name)
}

type EmailSender struct{}

func (es *EmailSender) SendWelcomeEmail(user *User) {
	// ユーザーにウェルカムメールを送信する処理
	fmt.Printf("Sending welcome email to %s.\n", user.Email)
}

func main() {
	user := &User{
		Name:  "JohnDoe",
		Email: "john@example.com",
	}

	userRepo := UserRepository{}
	emailSender := EmailSender{}

	userRepo.Save(user)
	emailSender.SendWelcomeEmail(user)
	user.Display()
}
