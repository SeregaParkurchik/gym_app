package main

import (
	"fmt"
)

const grpcPort = 50051

type server struct {
	desc.mustEmbedUnimplementedAdministratorV1Server
}

func main() {
	fmt.Println("serega")
}
