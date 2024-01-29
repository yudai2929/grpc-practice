package usecase

type AuthUsecase interface {
	SignUp(SignUpInput) (*SignUpOutput, error)
	SignIn(SignInInput) (*SignInOutput, error)
}
