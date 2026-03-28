package repositories

import (
    "github.com/frientia/gin-firebase-backend/config"
    "github.com/frientia/gin-firebase-backend/models"
)

// ProductRepository bertugas mengelola data produk
type ProductRepository struct{}

// Constructor untuk membuat instance baru
func NewProductRepository() *ProductRepository {
    return &ProductRepository{}
}

// FindAll mengambil semua produk aktif dengan pagination dan filter kategori
func (r *ProductRepository) FindAll(page, limit int, category string) ([]models.Product, int64, error) {
    var products []models.Product
    var total int64

    query := config.DB.Model(&models.Product{}).Where("is_active = ?", true)

    // Filter berdasarkan kategori jika ada
    if category != "" {
        query = query.Where("category = ?", category)
    }

    // Hitung total data untuk pagination
    query.Count(&total)

    // Ambil data sesuai offset & limit
    offset := (page - 1) * limit
    result := query.Offset(offset).Limit(limit).Find(&products)

    return products, total, result.Error
}

// FindByID mengambil satu produk berdasarkan ID
func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
    var product models.Product
    result := config.DB.First(&product, id)
    return &product, result.Error
}

// Create menyimpan produk baru
func (r *ProductRepository) Create(product *models.Product) error {
    return config.DB.Create(product).Error
}

// Update memperbarui produk
func (r *ProductRepository) Update(product *models.Product) error {
    return config.DB.Save(product).Error
}

// Delete melakukan soft-delete (tidak benar-benar menghapus dari DB)
func (r *ProductRepository) Delete(id uint) error {
    return config.DB.Delete(&models.Product{}, id).Error
}
