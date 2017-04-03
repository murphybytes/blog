package journal

// Database is an interface that exposes methods used to interact with
// mysql.
type Database interface {
	// Migrate creates database used by the application
	Migrate() error
	// Close frees up database connection
	Close() error
}
