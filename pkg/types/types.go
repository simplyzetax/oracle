package types

// Config holds the application configuration
type Config struct {
	APIKey string
	Model  string
}

// Question represents a user question
type Question struct {
	Text      string
	Model     string
	Timestamp int64
}

// Response represents an AI response
type Response struct {
	Text      string
	Model     string
	Timestamp int64
	Error     error
}
