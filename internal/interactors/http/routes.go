package server

func (s *Server) routes() {
	s.HandleFunc("/", s.indexPage()).Methods("GET")
	s.HandleFunc("/login", s.loginPage()).Methods("GET", "POST")
	s.HandleFunc("/home", s.homePage()).Methods("GET")
	s.HandleFunc("/static/css/bulma.min.css", s.bulmaCss()).Methods("GET")
	s.HandleFunc("/static/css/dropzone.min.css", s.dropzoneCss()).Methods("GET")
	s.HandleFunc("/static/css/basic.min.css", s.basicCss()).Methods("GET")
	s.HandleFunc("/static/js/dropzone.min.js", s.dropzoneJs()).Methods("GET")
}
