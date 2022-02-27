package main

import (
	"net/http"

	pb "user.proto"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	u := &pb.User{
		Name: proto.String("Nacho Herrera")
		Id: proto.Int32(123),
		Email: proto.String("ignacioromanherrera@gmail.com")
	}

	body, err := proto.Marshal(u)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/x-protobuf")
	res.Write(body)
}
