package helper

var version = "v1.0.0"
var Application = "golang"

func SayHello(name string) string {
	return "Hello " + name
}

func sayGoodBye(name string) string {
	return "Goodbye, " + name
}
