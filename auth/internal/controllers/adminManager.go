package controllers

import (
	"context"
	"go/auth-service/internal/helpers"
	"go/auth-service/internal/models"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PromoteAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		authHeader := c.Request.Header.Get("Authorization")

		err := helpers.CheckAdmin(authHeader)
		if err != nil {
			c.JSON(403, gin.H{"error": "Access denied: only admins can promote users"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
		defer cancel()

		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"user_id": id}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(404, gin.H{"error": "User not found"})
			} else if err == context.DeadlineExceeded {
				c.JSON(504, gin.H{"error": "Database timeout"})
			} else {
				c.JSON(500, gin.H{"error": "Database error"})
			}
			return
		}

		updateUser, err := helpers.UpdateMongo(id, "ADMIN")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "User promoted to admin",
			"user_id": updateUser.ID,
			"email":   updateUser.Email,
			"role":    updateUser.Type,
		})
	}
}

func DeleteAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		authHeader := c.Request.Header.Get("Authorization")

		err := helpers.CheckAdmin(authHeader)
		if err != nil {
			c.JSON(403, gin.H{"error": "Access denied: only admins can demote users"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
		defer cancel()

		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"user_id": id}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(404, gin.H{"error": "User not found"})
			} else if err == context.DeadlineExceeded {
				c.JSON(504, gin.H{"error": "Database timeout"})
			} else {
				c.JSON(500, gin.H{"error": "Database error"})
			}
			return
		}

		updateUser, err := helpers.UpdateMongo(id, "USER")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Admin demoted to user",
			"user_id": updateUser.ID,
			"email":   updateUser.Email,
			"role":    updateUser.Type,
		})
	}
}

func AdminPresence() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	filter := bson.M{"type": "ADMIN"}
	cursor, err := userCollection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	adminCount := 0
	for cursor.Next(ctx) {
		var admin models.User
		if err := cursor.Decode(&admin); err != nil {
			return err
		}
		if admin.Type != nil && *admin.Type == "ADMIN" {
			adminCount++
		}
	}

	if adminCount == 0 {
		file, err := os.ReadFile("/app/config/conf.yaml")
		if err != nil {
			return err
		}

		name, email, password, phone, err := takeAdminFromFile(file)
		if err != nil {
			return err
		}

		err = AdminCreatingWithContext(ctx, &name, &email, &password, &phone)
		if err != nil {
			return err
		}

		err = helpers.SendToNotifier(name, email, phone)
		if err != nil {
			return err
		}
	}

	return nil
}

func takeAdminFromFile(file []byte) (string, string, string, string, error) {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	data := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			data[key] = value
		}
	}

	return data["name"], data["email"], data["password"], data["phone"], nil
}
