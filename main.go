package main

var api = Api{
	Host: "127.0.0.1",
	Port: "8080",
}

func main() {
	if err := api.Init(); err != nil {
		panic(err)
	}
	if err := api.Start(); err != nil {
		panic(err)
	}
}
