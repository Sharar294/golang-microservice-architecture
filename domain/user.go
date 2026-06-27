package domain

// ==========================================
// 1. THE CONTRACT (DOMAIN LAYER)
// ==========================================
// Purpose: Defines the required behaviors for storing users.
// Any struct that has a CreateUser(name string) error method
// automatically satisfies this contract.
type UserRepository interface {
	CreateUser(name string) error
}
