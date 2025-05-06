package repo

import (
	"Go/global"
	"Go/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{
	sqlc *database.Queries
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

// GetUserByEmail implements IUserRepository.
func (up *userRepository) GetUserByEmail(email string) bool {

	// bỏ
	// Select * from user where email = ?? Order by email
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected

	// user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)

	// if err != nil {
    //     return false
    // }

	return false
}