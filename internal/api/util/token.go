package util

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// TODO create errors

func UsernameFromToken(c echo.Context) (string, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return "", errors.New("valid token is required")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("valid token claims are required")
	}
	username := claims["username"].(string)
	if !ok {
		return "", errors.New("token claims do not have username")
	}
	if username == "" {
		return "", errors.New("failed to get username from token claims")
	}
	return username, nil
}

func UserIDFromToken(c echo.Context) (int, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return 0, errors.New("valid token is required")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("valid token claims are required")
	}
	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("token claims do not have user id")
	}
	if userID == 0 {
		return 0, errors.New("failed to get user id from token claims")
	}
	return int(userID), nil
}
