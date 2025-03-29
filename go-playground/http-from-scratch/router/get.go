package router

func HandleGet(path string) (string, string) {
	switch path {
	case "/":
		return "200 OK", "Welcome to the homepage"

	case "/about":
		return "200 OK", "This is a raw Go server build with LOVE!"

	default:
		return "404 Not Found", "404 Not Found"
	}

}
