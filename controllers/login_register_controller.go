package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mahasiswa/config"
	"mahasiswa/models"
	"net/http"
	"strconv"
)

func Login(c echo.Context) error {
	account := &models.Account{}
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
		})
	}
	authAccount, err := models.Login(config.DB, account.Username, account.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Invalid username or password",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to login",
			"error":   err.Error(),
		})
	}
	token, err := models.GenerateToken(strconv.Itoa(authAccount.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to generate token",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success",
		"token":   token,
	})
}

func Register(c echo.Context) error {
	account := &models.Account{}
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
		})
	}

	if account.Username == "" || account.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Username and password are required",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
	}

	account.Password = string(hashedPassword)

	createdAccount, err := models.Register(config.DB, *account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to register",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, createdAccount)
}
