package main
 import (
 	"fmt"
 	"math/rand"
 	"time"
 )
 var currentPhone string
 var currentCode string
 func generateCode ()string{
 	rand.Seed(time.Now().UnixNano()) 
 	randomNum := rand.Intn(900000) + 100000
    return fmt.Sprintf("%d", randomNum)
 }
 func checkPhone(phone string) bool {
 	if len(phone) != 11 {
 		return false
 	}
 	for _, char := range phone {
 		if char < '0' || char > '9' {
 			return false
 		}
 	}
 	return true
 }
 // 处理“获取验证码”逻辑
 func getCode(phone string) {
 	currentCode = generateCode() // 生成验证码并存储
 	currentPhone = phone         // 关联当前手机号
 	fmt.Printf("\n 验证码已生成！手机号【%s】的验证码为：%s（有效期5分钟）\n", phone, currentCode)
    return
 }
 	// 接收用户输入的验证码并检验
    func loginWithCode(phone string) {
 	var inputCode string
 	fmt.Print("\n请输入您的验证码（6位数字）：")
 	fmt.Scanln(&inputCode)
 	// 检验验证码格式和是否正确
 	if len(inputCode) != 6 {
 		fmt.Println(" 验证码格式错误！必须是6位数字。")
 		return
 	}
 	if inputCode != currentCode {
 		fmt.Println(" 验证码输入错误！登录失败。")
 		return
 	}
 	fmt.Printf(" 恭喜！手机号【%s】通过验证码登录成功！\n", phone)
 }
 func main() {
 	fmt.Println("=== 手机验证码登录系统 ===")
 	// 第一步：获取并检验手机号
 	var phone string
 	fmt.Print("请输入您的手机号：")
 	fmt.Scanln(&phone)
 	// 校验手机号格式
 	if !checkPhone(phone) {
 		fmt.Println(" 手机号格式错误！必须是11位数字。")
 		return
 	}
 	// 第二步：选择操作（1=登录，2=获取验证码）
 	var option string
 	fmt.Print("\n请选择操作：1-使用验证码登录  2-获取验证码：")
 	fmt.Scanln(&option)
 	switch option {
 	case "1":
 		loginWithCode(phone) // 执行登录逻辑
 	case "2":
 		getCode(phone) // 执行获取验证码逻辑
 	default:
 		fmt.Println(" 未知操作！请输入1或2。")
 	}
 }
