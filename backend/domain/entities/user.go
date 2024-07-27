package entities

// このdomainディレクトリにはビジネスロジックやエンティティの定義を含みます。
// User represents a user entity
type User struct {
	ID       int
	Email    string
	Name     string
	Password string
}
