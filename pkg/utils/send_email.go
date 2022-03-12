package utils

// import (
//   "errors"
//   "fmt"
//   "log"
//   "net"
//   "regexp"

//   gomail "gopkg.in/gomail.v2"

//   "strings"

//   helper "github.com/sendgrid/sendgrid-go/helpers/mail"

//   "gitlab.com/golang-team-template/monolith/configs"
// )

// var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// var cfg = configs.Config()

// //IsEmailValid ...
// func IsEmailValid(e string) bool {
//   if len(e) < 3 && len(e) > 254 {
//     return false
//   }
//   if !emailRegex.MatchString(e) {
//     return false
//   }
//   parts := strings.Split(e, "@")
//   mx, err := net.LookupMX(parts[1])
//   if err != nil || len(mx) == 0 {
//     return false
//   }
//   return true
// }

// // SendEmailSendGrid ...
// func SendEmailSendGrid(email, code string) (err error) {

//   if !IsEmailValid(email) {
//     err = errors.New("email is not valid: " + email)
//     return
//   }

//   emailBody, err := ParseTemplate("./html/email.html", map[string]string{"code": code})
//   if err != nil {
//     log.Println("Error in parsing template!")
//     return
//   }

//   from := helper.NewEmail("ShelfIsh", cfg.SendgridEmail)
//   to := helper.NewEmail("client", email)
//   subject := "Verification code from ShelfIsh"
//   plainTextContent := emailBody
//   htmlContent := emailBody
//   message := helper.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

//   response, err := clients.SendGrid().Send(message)

//   if err != nil {
//     log.Println("error sending: ", err)
//   } else {
//     log.Println("Verification code sent!")
//     fmt.Println(response.StatusCode)
//     fmt.Println(response.Body)
//     // fmt.Println(response.Headers)
//   }

//   return
// }

// //SendEmailSMTP ...
// func SendEmailSMTP(email, code, fullName string) (err error) {
//   if !IsEmailValid(email) {
//     err = errors.New("email is not valid: " + email)
//     return
//   }

//   emailBody, err := ParseTemplate("./html/email.html", map[string]string{"code": code, "full_name": fullName})
//   if err != nil {
//     log.Println("Error in parsing template!")
//     return
//   }

//   m := gomail.NewMessage()
//   m.SetHeader("From", cfg.EmailFromHeader)
//   m.SetHeader("To", email)
//   m.SetHeader("Subject", "Verification email")
//   m.SetBody("text/html", emailBody)

//   fmt.Println("Sending ....")

//   // Send the email to
//   dialer := gomail.NewPlainDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPUserPass)
//   if err := dialer.DialAndSend(m); err != nil {
//     fmt.Println("err in sending email:", err)
//   }
//   return nil
// }