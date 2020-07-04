package main

import (
	proto "encrypt_service/proto"
	"fmt"
	"github.com/micro/go-micro"
)

//func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
//	return func(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
//		fmt.Printf("encryption request at time: %v", time.Now())
//		return fn(ctx, req, rsp)
//	}
//}

// go micro
// go micro web

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("encrypter"),
		//	micro.WrapClient(logWrapper),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterEncrypterHandler(service.Server(), new(Encrypter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
