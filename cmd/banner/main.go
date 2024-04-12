package main

import (
	"avito-banner/configs"
	"avito-banner/configs/logger"
	delivery "avito-banner/services/announcement/delivery/http"
	"avito-banner/services/announcement/usecase"
	"os"
	"sync"
)

func main() {
	log := logger.GetLogger()
	err := configs.InitEnv()
	if err != nil {
		log.Errorf("Init env error: %s", err.Error())
		return
	}

	grpcCfg, err := configs.GetGrpcConfig()
	if err != nil {
		log.Errorf("failed to parse grpc configs file: %s", err.Error())
		return
	}

	psxCfg, err := configs.GetMarketPsxConfig()
	if err != nil {
		log.Errorf("Create psx config error: %s", err.Error())
		return
	}

	core, err := usecase.GetCore(grpcCfg, psxCfg, log)
	if err != nil {
		log.Errorf("Create core error: %s", err.Error())
		return
	}

	api := delivery.GetApi(core, log)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		port := os.Getenv("BANNER_APP_PORT")
		log.Infof("Server banner running on %s", port)

		err = api.ListenAndServe(port)
		if err != nil {
			log.Errorf("Listen and serve error: %s", err.Error())
			return
		}
	}()

	wg.Wait()
}
