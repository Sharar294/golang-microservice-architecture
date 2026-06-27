package repository

import "fmt"

// ==========================================
// 2. DATA ACCESS (REPOSITORY LAYER) - MONGODB
// ==========================================
// Purpose: A second concrete implementation of the UserRepository contract.
// It handles connecting to and writing to a real database.
type MongoUserRepository struct {
	connectionString string
}

// Constructor to safely pass the connection string from main
func NewMongoUserRepository(connStr string) *MongoUserRepository {
	return &MongoUserRepository{
		connectionString: connStr,
	}
}

func (mongoRepo *MongoUserRepository) CreateUser(name string) error {
	// The specific logic required to save to Mongo lives here, safely hidden
	fmt.Printf("Saving %s to MongoDB at %s...\n", name, mongoRepo.connectionString)
	return nil
}
