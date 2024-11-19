package handlers

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/helpers"
	"go-boilerplate/infrastructure/ent"
	"go-boilerplate/infrastructure/ent/user"
	"go-boilerplate/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	client    *ent.Client
	secretKey string
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

func NewAuthHandler(client *ent.Client, secretKey string) *AuthHandler {
	return &AuthHandler{
		client:    client,
		secretKey: secretKey,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	exists, err := h.client.User.Query().
		Where(user.Email(req.Email)).
		Exist(c.Request.Context())
	if helpers.HandleServerErr(c, err, "database error") {
		return
	}
	if exists {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if helpers.HandleServerErr(c, err, "password hashing failed") {
		return
	}

	u, err := h.client.User.Create().
		SetEmail(req.Email).
		SetPassword(string(hashedPassword)).
		SetName(req.Name).
		Save(c.Request.Context())
	if helpers.HandleServerErr(c, err, "failed to create user") {
		return
	}

	token, err := jwt.GenerateToken(u.ID, u.Email, h.secretKey)
	if helpers.HandleServerErr(c, err, "failed to generate token") {
		return
	}

	c.JSON(201, gin.H{"token": token})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u, err := h.client.User.Query().
		Where(user.Email(req.Email)).
		Only(c.Request.Context())
	if helpers.HandleServerErr(c, err, "database error") {
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(u.ID, u.Email, h.secretKey)
	if helpers.HandleServerErr(c, err, "failed to generate token") {
		return
	}

	c.JSON(200, gin.H{"token": token})
}
