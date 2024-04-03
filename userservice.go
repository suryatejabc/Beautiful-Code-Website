type UserService struct {
	UserRepository IUserRepository
}

func NewUserService(repo IUserRepository) UserService {
	return UserService{UserRepository: repo}
}

func (s *UserService) GetUserByID(id int64) (*User, error) {
	return s.UserRepository.GetByID(id)
}