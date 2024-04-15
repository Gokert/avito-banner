package sql_requests

const (
	GetBannersByTagFeature = "SELECT\n    b.id AS banner_id,\n    b.id_feature AS feature_id,\n    v.content,\n    v.is_active AS is_active,\n    b.created_at,\n    v.updated_at\nFROM\n    banners b\n        JOIN\n    versions v ON b.id = v.id_banner\n        JOIN\n    banner_tags bt ON b.id = bt.id_banner\nWHERE\n    b.id_feature = $1 AND bt.id_tag = $2 AND v.is_active = true\nGROUP BY\n    b.id, b.id_feature, v.content, v.is_active, b.created_at, v.updated_at OFFSET $3 LIMIT $4;"
	GetBannersByTag        = "SELECT\n    b.id AS banner_id,\n    b.id_feature AS feature_id,\n    v.content,\n    v.is_active AS is_active,\n    b.created_at,\n    v.updated_at\nFROM\n    banners b\n        JOIN\n    versions v ON b.id = v.id_banner\n        JOIN\n    banner_tags bt ON b.id = bt.id_banner\nWHERE\n bt.id_tag = $1 AND v.is_active = true\nGROUP BY\n    b.id, b.id_feature, v.content, v.is_active, b.created_at, v.updated_at OFFSET $2 LIMIT $3;"
	GetBannersFeature      = "SELECT\n    b.id AS banner_id,\n    b.id_feature AS feature_id,\n    v.content,\n    v.is_active AS is_active,\n    b.created_at,\n    v.updated_at\nFROM\n    banners b\n        JOIN\n    versions v ON b.id = v.id_banner\n        WHERE\n    b.id_feature = $1 AND v.is_active = true\nGROUP BY\n    b.id, b.id_feature, v.content, v.is_active, b.created_at, v.updated_at OFFSET $2 LIMIT $3;"
	GetBanners             = "SELECT\n    b.id AS banner_id,\n    b.id_feature AS feature_id,\n    v.content,\n    v.is_active AS is_active,\n    b.created_at,\n    v.updated_at\nFROM\n    banners b\n        JOIN\n    versions v ON b.id = v.id_banner\n         WHERE v.is_active = true \nGROUP BY\n    b.id, b.id_feature, v.content, v.is_active, b.created_at, v.updated_at OFFSET $1 LIMIT $2;"

	GetAllBannersByTagFeature = "SELECT\n    b.id AS banner_id,\n    b.id_feature AS feature_id,\n    v.content,\n    v.is_active AS is_active,\n    b.created_at,\n    v.updated_at\nFROM\n    banners b\n        JOIN\n    versions v ON b.id = v.id_banner\n        JOIN\n    banner_tags bt ON b.id = bt.id_banner\nWHERE\n    b.id_feature = $1 AND bt.id_tag = $2\nGROUP BY\n    b.id, b.id_feature, v.content, v.is_active, b.created_at, v.updated_at OFFSET $3 LIMIT $4;"
	GetAllBannersByTag        = "SELECT\n    b.id AS banner_id,\n    b.id_feature AS feature_id,\n    v.content,\n    v.is_active AS is_active,\n    b.created_at,\n    v.updated_at\nFROM\n    banners b\n        JOIN\n    versions v ON b.id = v.id_banner\n        JOIN\n    banner_tags bt ON b.id = bt.id_banner\nWHERE\n bt.id_tag = $1\nGROUP BY\n    b.id, b.id_feature, v.content, v.is_active, b.created_at, v.updated_at OFFSET $2 LIMIT $3;"
	GetAllBannersFeature      = "SELECT\n    b.id AS banner_id,\n    b.id_feature AS feature_id,\n    v.content,\n    v.is_active AS is_active,\n    b.created_at,\n    v.updated_at\nFROM\n    banners b\n        JOIN\n    versions v ON b.id = v.id_banner\n        WHERE\n    b.id_feature = $1\nGROUP BY\n    b.id, b.id_feature, v.content, v.is_active, b.created_at, v.updated_at OFFSET $2 LIMIT $3;"
	GetAllBanners             = "SELECT\n    b.id AS banner_id,\n    b.id_feature AS feature_id,\n    v.content,\n    v.is_active AS is_active,\n    b.created_at,\n    v.updated_at\nFROM\n    banners b\n        JOIN\n    versions v ON b.id = v.id_banner\n        GROUP BY\n    b.id, b.id_feature, v.content, v.is_active, b.created_at, v.updated_at OFFSET $1 LIMIT $2;"
)

const GetUserBanner = "SELECT versions.content FROM versions JOIN banners on banners.id = versions.id_banner JOIN banner_tags on banners.id = banner_tags.id_banner WHERE banner_tags.id_tag = $1 AND banners.id_feature = $2 AND versions.is_active = true"
const GetUserBannerLast = "SELECT versions.content FROM versions JOIN banners on banners.id = versions.id_banner JOIN banner_tags on banners.id = banner_tags.id_banner WHERE banner_tags.id_tag = $1 AND banners.id_feature = $2 AND versions.is_active = true ORDER BY versions.updated_at LIMIT 1"