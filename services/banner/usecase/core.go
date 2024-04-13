package usecase

import (
	"avito-banner/configs"
	"avito-banner/pkg/models"
	_ "avito-banner/pkg/models"
	auth "avito-banner/services/authorization/delivery/proto"
	banner_repo "avito-banner/services/banner/repository/banner"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source=core.go -destination=../mocks/core_mock.go -package=mocks
type ICore interface {
	GetUserBanner(tagId uint64, featureId uint64) (*models.UserBanner, error)
	GetBanners(tagId uint64, featureId uint64, offset uint64, limit uint64) (*[]models.BannerResponse, error)
	CreateBanner(banner *models.BannerRequest) error
	CheckBanner(bannerId uint64) (bool, error)
	DeleteBanner(bannerId uint64) (bool, error)
	UpdateBanner(banner *models.BannerRequest) (bool, error)

	GetUserId(ctx context.Context, sid string) (uint64, error)
	GetRole(ctx context.Context, userId uint64) (string, error)
}

type Core struct {
	log    *logrus.Logger
	banner banner_repo.IRepository
	client auth.AuthorizationClient
}

func GetClient(address string) (auth.AuthorizationClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc connect err: %w", err)
	}
	client := auth.NewAuthorizationClient(conn)

	return client, nil
}

func GetCore(grpcCfg *configs.GrpcConfig, psxCfg *configs.DbPsxConfig, log *logrus.Logger) (*Core, error) {
	repo, err := banner_repo.GetPsxRepo(psxCfg)
	if err != nil {
		return nil, fmt.Errorf("get psx error error: %s", err.Error())
	}
	log.Info("Psx created successful")

	authRepo, err := GetClient(grpcCfg.Addr + ":" + grpcCfg.Port)
	if err != nil {
		return nil, fmt.Errorf("get auth repo error: %s", err.Error())
	}

	core := &Core{
		log:    log,
		banner: repo,
		client: authRepo,
	}

	return core, nil
}

func (c *Core) GetUserBanner(tagId uint64, featureId uint64) (*models.UserBanner, error) {
	banner, err := c.banner.GetUserBanner(tagId, featureId)
	if err != nil {
		return nil, fmt.Errorf("get user banner error: %s", err.Error())
	}

	return banner, nil
}

func (c *Core) GetBanners(tagId uint64, featureId uint64, offset uint64, limit uint64) (*[]models.BannerResponse, error) {
	banners, err := c.banner.GetBanners(tagId, featureId, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("get banners error: %s", err.Error())
	}

	return banners, nil
}

func (c *Core) CheckBanner(bannerId uint64) (bool, error) {
	check, err := c.banner.CheckBanner(bannerId)
	if err != nil {
		return false, fmt.Errorf("check banner error: %s", err.Error())
	}

	return check, nil
}

func (c *Core) DeleteBanner(bannerId uint64) (bool, error) {
	res, err := c.banner.CheckBanner(bannerId)
	if err != nil {
		return false, fmt.Errorf("update banner error: %s", err.Error())
	}

	if !res {
		return false, nil
	}

	_, err = c.banner.DeleteBanner(bannerId)
	if err != nil {
		return false, fmt.Errorf("delete banner error: %s", err.Error())
	}

	return true, nil
}

func (c *Core) UpdateBanner(banner *models.BannerRequest) (bool, error) {
	res, err := c.banner.CheckBanner(banner.BannerId)
	if err != nil {
		return false, fmt.Errorf("update banner error: %s", err.Error())
	}

	if !res {
		return false, nil
	}

	_, err = c.banner.UpdateBanner(banner)
	if err != nil {
		return false, fmt.Errorf("update banner error: %s", err.Error())
	}

	return true, nil
}

func (c *Core) CreateBanner(banner *models.BannerRequest) error {
	err := c.banner.CreateBanner(banner)
	if err != nil {
		return fmt.Errorf("create banner error: %s", err.Error())
	}

	return nil
}

func (c *Core) GetUserId(ctx context.Context, sid string) (uint64, error) {
	response, err := c.client.GetId(ctx, &auth.FindIdRequest{Sid: sid})
	if err != nil {
		return 0, fmt.Errorf("get user id err: %w", err)
	}
	return response.Value, nil
}

func (c *Core) GetRole(ctx context.Context, userId uint64) (string, error) {
	role, err := c.client.GetRole(ctx, &auth.RoleRequest{Id: userId})
	if err != nil {
		return "", fmt.Errorf("get role error: %s", err.Error())
	}

	return role.Role, nil
}
