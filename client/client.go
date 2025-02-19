package main

import "google.golang.org/grpc"

func main() {
	grpc.NewClient("localhost:9001")
}
