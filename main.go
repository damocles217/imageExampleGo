// ! Important, this app requires to be runned in sudo
// ! Example go build -o server main.go && sudo ./server
// ! Exapmle 2: sudo go run main.go
package main

import "github.com/damocles217/images_service/images"

func main() {
	r := images.CreateServer()
	r.Run("192.168.1.65:7700")
}
