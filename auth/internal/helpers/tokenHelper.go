package helpers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"go/auth-service/internal/config"
	"go/auth-service/internal/models"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const (
	localzone = "UTC+5"
	offset    = 5 * 60 * 60
)

type SignedDetails struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
	Uid      string `json:"uid"`
	jwt.RegisteredClaims
}

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var key = os.Getenv("KEY")

func hashToken(token string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(token), 12)
}

func CreateResetToken() (string, error) {
	n := 32
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func GetUserByResetToken(tokenID string) (*models.PasswordResetToken, error) {
	var tokenDetails models.PasswordResetToken

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(
		ctx,
		bson.M{"reset_token_id": tokenID},
	).Decode(&tokenDetails)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("invalid reset token")
		}

		return nil, err
	}

	return &tokenDetails, nil
}

func CreateToken(email, name, userType, userId string) (tokenWithClaims, refreshTokenWithClaims string, err error) {
	claims := SignedDetails{
		Email:    email,
		Name:     name,
		UserType: userType,
		Uid:      userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	refreshClaims := SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenWithClaims, err = token.SignedString([]byte(key))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenWithClaims, err = refreshToken.SignedString([]byte(key))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return
}

func UpdateResetToken(userId, token string) (string, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	localTime := time.FixedZone(localzone, offset)
	var updateObj primitive.D

	hashToken, err := hashToken(token)
	if err != nil {
		return "", err
	}
	resetTokenId := uuid.New().String()
	updateObj = append(updateObj, bson.E{Key: "reset_token_id", Value: resetTokenId})
	updateObj = append(updateObj, bson.E{Key: "hash_token", Value: string(hashToken)})
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: time.Now().In(localTime).Format(time.RFC3339)})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := userCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
	if err != nil {
		return "", fmt.Errorf("failed to update tokens for user %s: %v", userId, err)
	}
	if result.MatchedCount == 0 {
		return "", fmt.Errorf("user not found")
	}

	return resetTokenId, nil
}

func UpdateTokens(token, refreshToken, userId string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	localTime := time.FixedZone(localzone, offset)
	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{Key: "token", Value: token})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: refreshToken})
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: time.Now().In(localTime).Format(time.RFC3339)})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
	if err != nil {
		return fmt.Errorf("failed to update tokens for user %s: %v", userId, err)
	}

	return nil
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Sprintf("error loading .env file: %v", err)
	}

	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, fmt.Sprintf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, "invalid token"
	}

	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		return nil, "token is expired"
	}

	return claims, ""
}

func ExtractClaimsFromToken(userToken string) (*SignedDetails, error) {
	token, err := jwt.ParseWithClaims(userToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		return nil, fmt.Errorf("token is expired")
	}
	if claims.Email == "" || claims.UserType == "" {
		return nil, fmt.Errorf("missing essential claim data")
	}
	return claims, nil
}
