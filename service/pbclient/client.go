package pbclient

import (
	"context"
	"fmt"
	"log"

	proto "github.com/Anla3421/myGoProtobuf/myGoMemberServer/go"
	"google.golang.org/grpc"
)

//gPRC client 連線建立
func init() {
	fmt.Println("gRPC client initial")
	CreateConn()
}

var Client proto.GetInfoClient

func CreateConn() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	Client = proto.NewGetInfoClient(conn)
}

//gPRC client function
func GetMemberIDByJWT(JWT string) *proto.GetMemberIDByJWTRes {
	req := &proto.GetMemberIDByJWTReq{
		JWT: JWT,
	}
	res, err := Client.GetMemberIDByJWT(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC client:error while calling GetMemberIDByJWT Service: %v \n", err)
	}
	log.Printf("gRPC client:Response from GetMemberIDByJWT Service: %v", res)
	return res

}
