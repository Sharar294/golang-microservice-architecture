package repository

// ==========================================
// 2. DATA ACCESS (REPOSITORY LAYER) - IN MEMORY
// ==========================================
// Purpose: A concrete implementation of the UserRepository contract.
// It handles storing data temporarily in computer memory.
type InMemoryUserRepository struct {
	users []string // The specific storage mechanism (a slice)
}

// Constructor to safely initialize the struct from other packages
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []string{}, // Initialize the empty slice
	}
}

func (inmem *InMemoryUserRepository) CreateUser(name string) error {
	inmem.users = append(inmem.users, name)
	return nil
}
