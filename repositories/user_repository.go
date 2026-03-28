package repositories

import (
    "github.com/frientia/gin-firebase-backend/config"
    "github.com/frientia/gin-firebase-backend/models"
)

// UserRepository bertugas mengelola data user
type UserRepository struct{}

// Constructor untuk membuat instance baru
func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

// FindByFirebaseUID mencari user berdasarkan Firebase UID
func (r *UserRepository) FindByFirebaseUID(uid string) (*models.User, error) {
    var user models.User
    result := config.DB.Where("firebase_uid = ?", uid).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

// FindByEmail mencari user berdasarkan email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    result := config.DB.Where("email = ?", email).First(&user)
    return &user, result.Error
}

// Create menyimpan user baru ke database
func (r *UserRepository) Create(user *models.User) error {
    return config.DB.Create(user).Error
}

// Update memperbarui data user
func (r *UserRepository) Update(user *models.User) error {
    return config.DB.Save(user).Error
}
