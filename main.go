// package main

// import "fmt"

// // ==========================================
// // 1. THE CONTRACT (DOMAIN LAYER)
// // ==========================================
// // Purpose: Defines the required behaviors for storing users.
// // Any struct that has a CreateUser(name string) error method
// // automatically satisfies this contract.
// type UserRepository interface {
// 	CreateUser(name string) error
// }

// // ==========================================
// // 2. DATA ACCESS (REPOSITORY LAYER) - IN MEMORY
// // ==========================================
// // Purpose: A concrete implementation of the UserRepository contract.
// // It handles storing data temporarily in computer memory.
// type InMemoryUserRepository struct {
// 	users []string // The specific storage mechanism (a slice)
// }

// func (inmem *InMemoryUserRepository) CreateUser(name string) error {
// 	inmem.users = append(inmem.users, name)
// 	return nil
// }

// // ==========================================
// // 3. BUSINESS LOGIC (SERVICE LAYER)
// // ==========================================
// // Purpose: Handles the core rules of your application.
// // Notice how `repo` is type `UserRepository` (the interface).
// // The Service layer refuses to know about slices or MongoDB;
// // it only knows about the Contract.
// type UserService struct {
// 	repo UserRepository
// }

// func (user *UserService) RegisterUser(name string) error {
// 	// This is where you would put business rules, like checking
// 	// if the name is too short before saving it.

// 	// Call the interface method. We don't care how it gets saved!
// 	user.repo.CreateUser(name)
// 	return nil
// }

// // ==========================================
// // 2. DATA ACCESS (REPOSITORY LAYER) - MONGODB
// // ==========================================
// // Purpose: A second concrete implementation of the UserRepository contract.
// // It handles connecting to and writing to a real database.
// type MongoUserRepository struct {
// 	connectionString string
// }

// func (mongoRepo *MongoUserRepository) CreateUser(name string) error {
// 	// The specific logic required to save to Mongo lives here, safely hidden
// 	// from the Service layer.
// 	fmt.Printf("Saving %s to MongoDB at %s...\n", name, mongoRepo.connectionString)
// 	return nil
// }

// // ==========================================
// // 4. DEPENDENCY INJECTION (ENTRY POINT)
// // ==========================================
// // Purpose: This is where we wire the application together. We pick exactly
// // which Data Access implementation we want to use and "inject" it into the Service.
// func main() {
// 	// To switch back to memory storage, you would just uncomment this code
// 	// and comment out the MongoDB code below. The Service wouldn't care!
// 	// inmemRepo := &InMemoryUserRepository{}
// 	// service := &UserService{repo: inmemRepo}

// 	// Initialize our specific Data layer (MongoDB)
// 	mongoRepo := &MongoUserRepository{connectionString: "mongodb://localhost:27017"}

// 	// Inject the Data layer into our Service layer
// 	service := &UserService{repo: mongoRepo}

// 	// Execute our business logic
// 	service.RegisterUser("Sharar")
// }

package main

import (
	"myapp/repository"
	"myapp/service"
)

// ==========================================
// 4. DEPENDENCY INJECTION (ENTRY POINT)
// ==========================================
// Purpose: This is where we wire the application together. We pick exactly
// which Data Access implementation we want to use and "inject" it into the Service.
func main() {
	// To switch back to memory storage, you would just use:
	// inmemRepo := repository.NewInMemoryUserRepository()
	// userService := service.NewUserService(inmemRepo)

	// Initialize our specific Data layer (MongoDB)
	mongoRepo := repository.NewMongoUserRepository("mongodb://localhost:27017")

	// Inject the Data layer into our Service layer
	userService := service.NewUserService(mongoRepo)

	// Execute our business logic
	userService.RegisterUser("Sharar")
}
