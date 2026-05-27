package main


const (
	OK = 200
	Created = 201
	BadRequest = 400
	Unauthorized = 401
	Forbidden = 403
	NotFound = 404
	InternalServerError = 500
)

func main() {
	println("OK:", OK)
	println("Created:", Created)
	println("Bad Request:", BadRequest)
	println("Unauthorized:", Unauthorized)
	println("Forbidden:", Forbidden)
	println("Not Found:", NotFound)
	println("Internal Server Error:", InternalServerError)
}