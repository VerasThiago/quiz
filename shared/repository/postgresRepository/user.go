package postgresrepository

import "github.com/verasthiago/quiz/shared/models"

func (p *PostgresRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if record := p.db.Where("email = ?", email).First(&user); record.Error != nil {
		return nil, record.Error
	}
	return &user, nil
}

func (p *PostgresRepository) CreateUser(user *models.User) error {
	return p.db.Create(user).Error
}

func (p *PostgresRepository) UpdateUser(user *models.User) error {
	if user.Password != "" {
		if err := user.HashPassword(user.Password); err != nil {
			return err
		}
	}
	return p.db.Model(user).Updates(user).Error
}

func (p *PostgresRepository) DeleteUser(userID string) error {
	if err := p.db.Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
		return nil
	}
	return p.DeleteQuizzesByUserID(userID)
}
