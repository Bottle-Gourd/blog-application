package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	"github.com/Bottle-Gourd/blog-application/global"
	"github.com/Bottle-Gourd/blog-application/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

type authServer struct{}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	login, password := in.GetLogin(), in.GetPassword()

	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()

	var user global.User

	global.DB.Collection("user").FindOne(ctx, bson.M{"$or": []bson.M{bson.M{"username": login}, bson.M{"email": login}}}).Decode(&user)
	if user == global.NilUser {
		return &proto.AuthResponse{}, errors.New("wrong Login credential provided")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return &proto.AuthResponse{}, errors.New("wrong Login credential provided")
	}

	return &proto.AuthResponse{Token: user.GetToken()}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, &authServer{})

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal("Error creating listener: ", err.Error())
	}

	server.Serve(listener)
}

//
