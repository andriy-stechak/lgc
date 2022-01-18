// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/andriystech/lgc/api/server"
	"github.com/andriystech/lgc/config"
	"github.com/andriystech/lgc/db/repositories"
	"github.com/andriystech/lgc/facilities/mongo"
	"github.com/andriystech/lgc/facilities/ws"
	"github.com/andriystech/lgc/services"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewServer(db mongo.ClientHelper) server.HttpServer {
	serverConfig := config.GetServerConfig()
	tokensRepository := repositories.NewTokensRepository(serverConfig)
	tokenService := services.NewTokenService(tokensRepository)
	usersCollection := mongo.NewUsersCollection(db, serverConfig)
	usersRepository := repositories.NewUsersRepository(usersCollection)
	userService := services.NewUserService(usersRepository)
	connectionsRepository := repositories.NewConnectionsRepository()
	messagesCollection := mongo.NewMessagesCollection(db, serverConfig)
	messagesRepository := repositories.NewMessagesRepository(messagesCollection)
	upgraderHelper := ws.NewUpgrader(serverConfig)
	webSocketService := services.NewWebSocketService(connectionsRepository, messagesRepository, usersRepository, upgraderHelper)
	httpServer := server.NewHttpServer(tokenService, userService, webSocketService, serverConfig)
	return httpServer
}

// wire.go:

var collectionsSet = wire.NewSet(mongo.NewMessagesCollection, mongo.NewUsersCollection)

var repositoriesSet = wire.NewSet(repositories.NewConnectionsRepository, repositories.NewMessagesRepository, repositories.NewTokensRepository, repositories.NewUsersRepository)

var servicesSet = wire.NewSet(services.NewTokenService, services.NewUserService, services.NewWebSocketService)
