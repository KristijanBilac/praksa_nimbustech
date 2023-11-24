package main

import (
	"fmt"
)

// User struct represents a user.
type User struct {
	ID   int
	Name string
}

// UserService provides operations related to users.
type UserService struct {
}

// NewUserService creates a new instance of UserService with any necessary dependencies.
func NewUserService() *UserService {
	return &UserService{}
}

// GetUserByID retrieves a user by ID.
func (us *UserService) GetUserByID(userID int) *User {
	// In a real application, this would fetch a user from a database or some external service.
	// For simplicity, we just create a dummy user here.
	return &User{ID: userID, Name: "John Doe"}
}

// UserController handles HTTP requests related to users.
type UserController struct {
	UserService *UserService // Dependency injection via struct field
}

// NewUserController creates a new instance of UserController with the provided UserService dependency.
func NewUserController(userService *UserService) *UserController {
	return &UserController{UserService: userService}
}

// GetUserHandler handles HTTP requests to get a user by ID.
func (uc *UserController) GetUserHandler(userID int) {
	user := uc.UserService.GetUserByID(userID)
	fmt.Printf("User: %+v\n", user)
	// In a real application, you would return the user as JSON or in some other format.
}

func main() {
	// Dependency injection for UserService
	userService := NewUserService()

	// Dependency injection for UserController
	userController := NewUserController(userService)

	// Simulating an HTTP request to get a user by ID
	userController.GetUserHandler(123)
}
