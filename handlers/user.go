package handlers

import (
	"net/http"
	"viventis/schemas"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// @Summary Criar um novo usuário
// @Description Cria um novo usuário e armazena no banco de dados.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "Create User"
// @Sucess 201 {object} schemas.User
// @Failure 400 {object} gin.H
// @Router /users [post]
func createUser(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	user := schemas.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Erro ao realizar o hash da senha: " + err.Error()})
	}

	if err := schemas.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", createUser)
}
