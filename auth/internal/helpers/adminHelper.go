package helpers

import (
	"context"
	"fmt"
	"go/auth-service/internal/models"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const contextTime = time.Second * 10

func TokenTaking(authHeader string) (string, error) {
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return "", fmt.Errorf("your token is empty")
	}
	return token, nil
}

func CheckAdmin(authHeader string) error {
	token, err := TokenTaking(authHeader)
	if err != nil {
		return err
	}
	claims, msg := ValidateToken(token)
	if msg != "" {
		return fmt.Errorf("error: %v", msg)
	}
	if claims.UserType != "ADMIN" {
		return fmt.Errorf("access is allowed only to administrators")
	}
	return nil
}

func UpdateMongo(id, usertype string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()

	update := bson.M{"$set": bson.M{"type": usertype}}
	res, err := userCollection.UpdateOne(ctx, bson.M{"user_id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update user type: %v", err)
	}
	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("user not found")
	}

	var updatedUser models.User
	err = userCollection.FindOne(ctx, bson.M{"user_id": id}).Decode(&updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated user data")
	}

	token, refreshToken, err := CreateToken(*updatedUser.Email, *updatedUser.Name, usertype, id)
	if err != nil {
		return nil, err
	}

	err = UpdateTokens(token, refreshToken, id)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

var httpClient = &http.Client{}

func SendToNotifier(name, email, phone string) error {
	jsonData := strings.NewReader(fmt.Sprintf(`{"name": "%s", "email": "%s", "phone": "%s"}`, name, email, phone))
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", "http://notifier-service:8082/admin", jsonData)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification, status code: %d", resp.StatusCode)
	}

	return nil
}
