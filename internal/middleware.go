package internal

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Middleware is for Basic Authentication
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение заголовка Authorization
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if auth[0] != "Basic" || !authenticateUser(auth) {
			respondWithError(401, "Unauthorized", c)
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		if authenticateUser(auth) {
			c.JSON(200, gin.H{"message": "OK"})
		}

		//Прохождение middleware, если аутентификация прошла успешно
		c.Next()
	}
}

// authenticateUser is for Basic Auth
func authenticateUser(auth []string) bool {
	// Декодирование base64-строки
	decoded, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		return false
	}

	// Разбиение декодированной строки на логин и пароль
	credentials := strings.Split(string(decoded), ":")
	if len(credentials) != 2 {
		return false
	}

	// Проверка логина и пароля
	username := viper.GetString("USERNAME_BASIC")
	password := viper.GetString("PASSWORD_BASIC")

	if credentials[0] != username || credentials[1] != password {
		return false
	}
	return true
}

// respondWithError is for formatting error respond
func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
