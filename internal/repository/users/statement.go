package users

const (
	createUserStmt = `
		INSERT INTO users (username, age, hashed_pass) 
		VALUES 
		($1, $2, $3)
	`
)
