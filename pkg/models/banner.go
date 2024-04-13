package models

import "time"

type Banner struct {
	Id        uint64                 `json:"id"`
	Url       string                 `json:"url"`
	TagIds    []uint64               `json:"tag_ids"`
	FeatureId uint64                 `json:"feature_id"`
	Content   map[string]interface{} `json:"content,omitempty"`
}

type UserBanner struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

type BannerRequest struct {
	BannerId  uint64                 `json:"banner_id"`
	TagIds    []uint64               `json:"tag_ids"`
	FeatureId uint64                 `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

type BannerResponse struct {
	BannerId  uint64                 `json:"banner_id"`
	TagIds    []uint64               `json:"tag_ids"`
	FeatureId uint64                 `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
	CreatedAt *time.Time             `json:"created_at"`
	UpdatedAt *time.Time             `json:"updated_at"`
}
