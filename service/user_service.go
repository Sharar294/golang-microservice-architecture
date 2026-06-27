package service

import "myapp/domain"

// ==========================================
// 3. BUSINESS LOGIC (SERVICE LAYER)
// ==========================================
// Purpose: Handles the core rules of your application.
// Notice how `repo` is type `domain.UserRepository` (the interface).
type UserService struct {
	repo domain.UserRepository
}

// Constructor to inject the repository
func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (user *UserService) RegisterUser(name string) error {
	// Call the interface method. We don't care how it gets saved!
	user.repo.CreateUser(name)
	return nil
}
