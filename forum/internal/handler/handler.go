package handler

import (
	"forum/internal/service"
	"html/template"
	"net/http"
	"path/filepath"
)

type Handler struct {
	Mux     *http.ServeMux
	Temp    *template.Template
	Service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Mux:     http.NewServeMux(),
		Temp:    template.Must(template.ParseGlob("./front/html/*.html")),
		Service: services,
	}
}

func (h *Handler) InitRoutes() {
	h.Mux.HandleFunc("/", h.middleWareGetUser(h.homePage))

	h.Mux.HandleFunc("/signup", h.signUp)
	h.Mux.HandleFunc("/signin", h.signIn)

	h.Mux.HandleFunc("/post/", h.middleWareGetUser(h.postPage))
	h.Mux.HandleFunc("/post/create", h.middleWareGetUser(h.createPost))
	h.Mux.HandleFunc("/post/myPost", h.middleWareGetUser(h.myPost))
	h.Mux.HandleFunc("/post/myLikedPost", h.middleWareGetUser(h.myLikedPost))

	h.Mux.HandleFunc("/emotion/post/", h.middleWareGetUser(h.emotionPost))
	h.Mux.HandleFunc("/emotion/comment/", h.middleWareGetUser(h.emotionComment))

	h.Mux.HandleFunc("/logout", h.logOut)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./front/static/")})
	h.Mux.Handle("/static", http.NotFoundHandler())
	h.Mux.Handle("/static/", http.StripPrefix("/static", fileServer))
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
