package server

func (s *Server) routes() {
	//regular endpoints
	s.HandleFunc("/", s.indexPage()).Methods("GET")
	s.HandleFunc("/login", s.loginPage()).Methods("GET", "POST")
	s.HandleFunc("/upload", s.uploadEndpoint()).Methods("POST")
	//static files
	s.HandleFunc("/static/css/bulma.min.css", s.bulmaCss()).Methods("GET")
	s.HandleFunc("/static/css/dropzone.min.css", s.dropzoneCss()).Methods("GET")
	s.HandleFunc("/static/css/custom.css", s.customCss()).Methods("GET")
	s.HandleFunc("/static/js/dropzone.min.js", s.dropzoneJs()).Methods("GET")
}
