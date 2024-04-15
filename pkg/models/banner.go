package models

import "time"

type UserBanner struct {
	Content string `json:"content"`
}

type Banner struct {
	BannerId  uint64            `json:"banner_id"`
	TagIds    []uint64          `json:"tag_ids"`
	FeatureId uint64            `json:"feature_id"`
	Content   map[string]string `json:"content"`
	IsActive  bool              `json:"is_active"`
}

type BannerRequest struct {
	BannerId  uint64   `json:"banner_id"`
	TagIds    []uint64 `json:"tag_ids"`
	FeatureId uint64   `json:"feature_id"`
	Content   string   `json:"content"`
	IsActive  bool     `json:"is_active"`
}

type BannerResponse struct {
	BannerId  uint64     `json:"banner_id"`
	TagIds    []uint64   `json:"tag_ids"`
	FeatureId uint64     `json:"feature_id"`
	Content   string     `json:"content"`
	IsActive  bool       `json:"is_active"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
