package helpers

import (
	"context"
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func CheckType(c *gin.Context, userId string) string {
	uid := c.GetString("uid")
	userType := c.GetString("user_type")
	if userType == "ADMIN" && uid != userId {
		return ""
	} else {
		return "Unauthorized access to the server"
	}
}

func Hash(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashed)
}

func VerifyingOfPassword(userPassword, foundUserPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(foundUserPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		check = false
		msg = "Email or password is incorrect"
	}
	return check, msg
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func TakeName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		var foundUser User
		err = userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": foundUser.Name})
	}
}

func SendTokenToEmail(email, token, resetTokenID string) error {
	senderEmail, frontendAPI, smtpPass := os.Getenv("EMAIL_SENDER"), os.Getenv("FRONTEND_API"), os.Getenv("SMTP_PASS")
	if senderEmail == "" || frontendAPI == "" || smtpPass == "" {
		return fmt.Errorf("sender's email or frontend_api is null")
	}
	resetURL := fmt.Sprintf("%s/reset-password?id=%s&token=%s", frontendAPI, resetTokenID, token)
	message := gomail.NewMessage()

	message.SetHeader("From", senderEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Reset Your CRM Password")

	message.SetBody("text/plain", fmt.Sprintf(
		"You requested a password reset.\n\n"+
			"Open the following link to reset your password:\n%s\n\n"+
			"This link expires in 30 minutes.\n"+
			"If you did not request a password reset, ignore this email.",
		resetURL,
	))

	message.AddAlternative("text/html", fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Password Reset</title>
		</head>
		<body style="margin:0;padding:0;background-color:#f4f4f5;
					font-family:Arial,sans-serif;color:#18181b;">
			<div style="max-width:560px;margin:40px auto;padding:32px;
						background:#ffffff;border-radius:12px;">
				<h2 style="margin-top:0;">Reset your password</h2>

				<p>You requested a password reset for your CRM account.</p>

				<p>Click the button below to choose a new password.</p>

				<div style="margin:28px 0;">
					<a href="%s"
					style="display:inline-block;padding:14px 24px;
							background:#18181b;color:#ffffff;
							text-decoration:none;border-radius:6px;
							font-weight:bold;">
						Reset Password
					</a>
				</div>

				<p style="font-size:14px;color:#52525b;">
					This link expires in 30 minutes.
				</p>

				<p style="font-size:14px;color:#52525b;">
					If you did not request this, you can safely ignore this email.
				</p>

				<hr style="border:none;border-top:1px solid #e4e4e7;margin:24px 0;">

				<p style="font-size:12px;color:#71717a;">
					If the button doesn't work, copy this URL into your browser:
					<br>
					<a href="%s" style="color:#2563eb;word-break:break-all;">%s</a>
				</p>
			</div>
		</body>
		</html>
		`, html.EscapeString(resetURL), html.EscapeString(resetURL), html.EscapeString(resetURL)))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, smtpPass)

	err := dialer.DialAndSend(message)
	if err != nil {
		return err
	}
	return nil
}
