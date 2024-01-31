package rest

import (
	"encoding/json"
	"example/internal/domain"
	
	"io"

	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type User struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

// @Summary Login a user
// @Description Login a user by providing their full name and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body loginRequest true "loginRequest object that needs to be login"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/login [post]
func (h *Handler) LoginUserHandler(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	err := h.userService.LoginUser(req.UserName, req.Password)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.GenerateToken(req.UserName,req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "UserName not found: " + err.Error()})
		return
	}

	response := gin.H{"access_token": token}

	ctx.JSON(http.StatusOK, response)

}

// @Summary Sign up a new user
// @Description Sign up a new user with the provided JSON data
// @Tags auth
// @Accept json
// @Produce json
// @Param user body User true "User object that needs to be signed up"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/sign-up [post]
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

// @Summary Get paginated list of users
// @Security ApiKeyAuth
// @Description Get a paginated list of users based on provided input
// @Tags users
// @Accept json
// @Produce json
// @Param input body GetPaginationInput true "Pagination Input"
// @Success 200 {array} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/page [get]
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
