package ports

type HashPort interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}