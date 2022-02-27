// go:generate codecgen -o user_generated user.go

package user

// User struct is annotated for codec instead of JSON
type User struct {
	Name string `codec:"name"`
	Email string `codec:",omitempty"`
}

func main() {
	// Encoding
	jh := new(codec.JsonHandle)
	u := &user.User {
		Name: "Nacho Herrera"
		Email: "ignacioromanherrera@gmail.com"
	}

	var out[]byte
	err := codec.NewEncoderBytes(&out, jh).Encode(&u)
	if err != nil { panic(err) } 

	fmt.Println(string(out))

	// Decoding
	var u2 user.User
	err = codec.NewDecoderBytes(out, jh).Decode(&u2)
	if err != nil { panic(err) }

	fmt.Println(u2)
}
