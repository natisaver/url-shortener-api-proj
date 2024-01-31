package config

// db
const (
	Host       = "localhost"
	Port       = 5432
	User       = "postgres"
	Password   = "123"
	Dbname     = "urlDB"
	Dbtestname = "urlDBTest"
)

// context key
// of type int
// used in context to pass around objects {ctxKey: ctxValue}
type CtxKey int

// iota
// => means enums start at index 0 and auto-increment accordingly
// in our use case, these keys help us pass around different implementations of functions for testing
const (
	CtxKeyDB CtxKey = iota
	CtxKeyMockCRUDRepository
)
