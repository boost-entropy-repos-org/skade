package server

func (s *Server) routes() {
	s.HandleFunc("/", s.indexPage()).Methods("GET")
	s.HandleFunc("/login", s.loginPage()).Methods("GET", "POST")
	s.HandleFunc("/home", s.homePage()).Methods("GET")
	s.HandleFunc("/static/css/bulma.min.css", s.css()).Methods("GET")
}
