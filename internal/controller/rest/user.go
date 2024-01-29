package rest

import (
	"encoding/json"
	"example/internal/domain"
	"example/pkg/jwt"
	"io"

	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	FullName string
	Password string
}

func (h *Handler) LoginUserHandler(ctx *gin.Context) {
	var req loginRequest

	// Decoding requested body to Go object
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	passed, err := h.userService.LoginUser(req.FullName, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid password or fullname "})
		return
	}

	if passed {
		token, err := jwt.CreateToken(req.FullName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Phone number not found: " + err.Error()})
			return
		}

		response := gin.H{"access_token": token}

		ctx.JSON(http.StatusOK, response)

	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func (h *Handler) SignUpUserHandler(c *gin.Context) {
	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var user domain.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	err = h.userService.SignUp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *Handler) GetPagesUser(c *gin.Context) {
	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var inp domain.GetPaginationInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	products, err := h.userService.FindAll(int(inp.Page), int(inp.Limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve"})
		return
	}

	c.JSON(http.StatusOK, products)
}
