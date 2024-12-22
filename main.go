package main

import (
	"AnA-Roaming/ana-authenticator/services"
	repo_config "AnA-Roaming/repo-config"
	logger_infra "AnA-Roaming/repo-infra/logger-infra"
	repo_custom_logger "AnA-Roaming/repo-infra/logger-infra/repo-custom-logger"
	mongo_infra "AnA-Roaming/repo-infra/mongo-infra"
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func StartUpModules() fx.Option {
	return fx.Options(
		fx.Provide(
			repo_config.NewApplication,
			repo_config.NewConfig,
			logger_infra.NewLogger,
			repo_custom_logger.NewRepoCustomLogger,
			mongo_infra.NewMongoInfra,
			services.NewServices,
		),
		fx.Invoke(),
	)
}

func Init() {
	// This is the entry point of the application
	var logger = new(zap.SugaredLogger)

	fxApp := fx.New(
		StartUpModules(),
		fx.Populate(&logger),
	)

	if err := fxApp.Start(context.Background()); err != nil {
		fmt.Print(fmt.Sprintf("Auth-Service: StartUp Error: %s", err.Error()))
	}

	// Start the pprof server
	go func() {
		logger.Infof(" ListenAndServe ERROR=%v", http.ListenAndServe("0.0.0.0:5041", nil))
	}()

	if err := fxApp.Err(); err != nil {
		logger.Infof("fxApp startup error in Auth-Service, ERROR is %v", err)
	}
	logger.Infof("Started AnA:Authenticator Service...")
	<-fxApp.Done()
}

func main() {
	Init()
}
