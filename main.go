package main


var api  = Api{
	Addr: ":8080",
}
func main() {
	if err := api.init(); err != nil {
		panic(err)
	}
	if err := api.Start(); err != nil {
		panic(err)
	}
}
