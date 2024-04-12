package usecase

import (
	"avito-banner/configs"
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
