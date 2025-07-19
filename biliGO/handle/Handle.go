package main

import "net/http"

// 自定义一个handler
type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func main() {
	mh := helloHandler{}
	a := aboutHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.Handle("/hello", &mh)
	http.Handle("/about", &a)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/home", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("home"))
	})
	server.ListenAndServe()
	// 下面这句话和上面的作用是一样的
	//http.ListenAndServe("localhost:8080", nil)
}
