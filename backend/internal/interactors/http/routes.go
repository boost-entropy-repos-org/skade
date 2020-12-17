package server

func (h *httpInteractor) routes() {
	//regular endpoints
	h.router.HandleFunc("/", h.indexPage()).Methods("GET")
	h.router.HandleFunc("/upload", h.uploadEndpoint()).Methods("POST")
	//static files
	h.router.HandleFunc("/static/css/bulma.min.css", h.bulmaCss()).Methods("GET")
	h.router.HandleFunc("/static/css/dropzone.min.css", h.dropzoneCss()).Methods("GET")
	h.router.HandleFunc("/static/css/custom.css", h.customCss()).Methods("GET")
	h.router.HandleFunc("/static/js/dropzone.min.js", h.dropzoneJs()).Methods("GET")
}
