package dtoGrpcUser

type CreateRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
}

type UpdateRequest struct {
	Id    int64
	Name  string
	Email string
}
