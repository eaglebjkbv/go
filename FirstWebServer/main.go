package main

import ("fmt"
		"net/http")
func index_handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"İlk Web Applikasyonum")
}
func about_handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hakkımda")
}

func main(){

	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about",about_handler)
	http.ListenAndServe(":8000",nil)




}