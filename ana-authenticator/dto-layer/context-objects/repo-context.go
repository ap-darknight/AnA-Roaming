package context_objects

import (
	repo_custom_logger "AnA-Roaming/repo-infra/logger-infra/repo-custom-logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type RepoContext struct {
	ReqID        string
	Logger       *repo_custom_logger.CustomLogger
	UserContext  UserContext
	StartTime    time.Time
	RootFunction string
	GinContext   *gin.Context
}

func NewRepoContext(log *repo_custom_logger.CustomLogger, rootFunction string, userData UserContext) RepoContext {
	start := time.Now()
	flowUUID, _ := uuid.NewUUID()
	log.ReqID = flowUUID.String()
	log.RootFunction = rootFunction
	basicContext := RepoContext{
		ReqID:        flowUUID.String(),
		Logger:       log,
		UserContext:  userData,
		StartTime:    start,
		RootFunction: rootFunction,
		GinContext:   nil,
	}
	return basicContext
}

func NewRepoContextWithGin(log *repo_custom_logger.CustomLogger, rootFunction string, c *gin.Context) RepoContext {
	start := time.Now()
	flowUUID, _ := uuid.NewUUID()
	log.ReqID = flowUUID.String()
	log.RootFunction = rootFunction
	contextData := RepoContext{
		ReqID:        flowUUID.String(),
		Logger:       log,
		UserContext:  UserContext{},
		StartTime:    start,
		RootFunction: rootFunction,
		GinContext:   c,
	}

	return contextData
}
