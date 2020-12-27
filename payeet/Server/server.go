package main

import (
	"log"
	"net"

	"galil-maaravi-802-payeet/payeet/Server/services"
	"galil-maaravi-802-payeet/payeet/Server/util"
	pb "galil-maaravi-802-payeet/payeet/protos/go"

	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

func accessibleRoles() map[string][]string {
	const path = "/payeet.payeet/"
	return map[string][]string{
		path + "GetUserInfo":     {"user"},
		path + "TransferBalance": {"user"},
		path + "GetBalance":      {"user"},
	}

}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load config : ", err)
	}

	log.Printf("Starting server on port [%s]", config.Port)

	userStore := services.NewMongoUserStore(config.ConnectionString, config.DBName, config.CollectionName)
	userStore.Connect()
	defer userStore.Disconnect()
	jwtManger, err := services.NewJWTManager(config.AccessTokenDuration, config.RefreshTokenDuration, config.AccessTokenKey, config.RefreshTokenKey)

	if err != nil {
		log.Panic(err)
	}

	authServer := services.NewAuthServer(userStore, jwtManger)
	logic := services.NewPayeetServer(userStore, jwtManger)

	interceptor := services.NewAuthInterceptor(jwtManger, accessibleRoles())
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
	)

	pb.RegisterPayeetAuthServer(srv, authServer)
	pb.RegisterPayeetServer(srv, logic)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Done! ✅")

	log.Printf("Serving... 🥳")
	if e := srv.Serve(lis); e != nil {
		log.Panic(e)
	}

}
