package server

import (
	"encoding/json"
	"net/http"
)

// UserRequest represents the request structure for getting a user.
type UserRequest struct {
	ID int `json:"id" example:"1" binding:"required"` // User ID
}

// UserResponse represents the response structure for a user.
type UserResponse struct {
	ID    int    `json:"id"`    // User ID
	Name  string `json:"name"`  // User's name
	Email string `json:"email"` // User's email
}

// ErrorResponse represents the error response structure.
type ErrorResponse struct {
	Message string `json:"message"` // Error message
}

// GetUserHandler handles GET requests for fetching a user by ID.
// @Summary Get user by ID
// @Description Get details of a user by their ID
// @ID get-user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponse
// @Failure 404 {object} ErrorResponse
// @Router /user/{id} [get]
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/"):]

	// Here you would typically fetch the user from a database
	// For demonstration, we assume the user ID is 1
	if id != "1" { // Example check for user existence
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := UserResponse{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
