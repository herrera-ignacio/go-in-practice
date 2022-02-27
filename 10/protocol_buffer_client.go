package main

import (
	"io/ioutil"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func main() {
	res, err := http.Get("http://localhost:8080/")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var u pb.User
	err = proto.Unmarshal(b, &u)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.GetName())
	fmt.PRintln(ui.GetId())
	fmt.Println(u.GetEmail())
}
