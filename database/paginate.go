package database

import (
	"math"
)

// Struktur untuk response pagination
type PaginatedResponse struct {
	Status string         `json:"status"`
	Data   interface{}    `json:"data"`
	Meta   PaginationMeta `json:"meta"`
}

// Metadata untuk pagination
type PaginationMeta struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

// Request struct untuk pagination
type PaginationRequest struct {
	Length int         `json:"length,omitempty"`
	Page   int         `json:"page,omitempty"`
	Dest   interface{} `json:"dest,omitempty"`
}

// Fungsi Paginate untuk QueryBuilder
func (qb *QueryBuilder) Paginate(request PaginationRequest) (PaginatedResponse, error) {
	// Pastikan page minimal 1
	if request.Page < 1 {
		request.Page = 1
	}

	// Pastikan length memiliki nilai default jika kosong
	if request.Length <= 0 {
		request.Length = 10 // Default 10 item per halaman
	}

	// Hitung total data menggunakan model yang sesuai
	var total int64
	if err := qb.db.Model(request.Dest).Count(&total).Error; err != nil {
		return PaginatedResponse{}, err
	}

	// Hitung offset berdasarkan page dan length
	offset := (request.Page - 1) * request.Length

	// Ambil data sesuai pagination
	if err := qb.db.Limit(request.Length).Offset(offset).Find(request.Dest).Error; err != nil {
		return PaginatedResponse{}, err
	}

	// Hitung total halaman
	totalPages := int(math.Ceil(float64(total) / float64(request.Length)))

	// Kembalikan response pagination
	return PaginatedResponse{
		Status: "success",
		Data:   request.Dest,
		Meta: PaginationMeta{
			Total:       int(total),
			PerPage:     request.Length,
			CurrentPage: request.Page,
			TotalPages:  totalPages,
		},
	}, nil
}
