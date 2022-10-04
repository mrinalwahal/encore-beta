package user

const (
	GET_ALL    = `SELECT * FROM users`
	GET        = `SELECT * FROM users WHERE id=$1 LIMIT 1`
	INSERT     = `INSERT INTO users (full_name, active) VALUES ($1, $2) RETURNING id, full_name, active`
	DELETE     = `DELETE FROM users WHERE id=$1`
	DELETE_ALL = `DELETE FROM users`
)
