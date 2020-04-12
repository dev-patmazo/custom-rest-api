package router

// func A(inner http.Handler) http.Handler {
// 	log.Print("A: called")
// 	mw := func(w http.ResponseWriter, r *http.Request) {
// 		log.Print("A: before")
// 		inner.ServeHTTP(w, r)
// 		log.Print("A: after")
// 	}
// 	return http.HandlerFunc(mw)
// }
