package route

import (
	"gd-blog/src/ioc"
	"net/http"
)

func InitRoutes() error {
	blogHandler, err := ioc.Provide("blogHandler")
	if err != nil {
		return err
	}
	handler := blogHandler.(BlogHandler)
	http.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.HandleListBlog()(w, r)
		case http.MethodPost:
			handler.HandleCreateBlog()(w, r)
		}
	})
	http.HandleFunc("/blogs/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.HandleGetBlog()(w, r)
		case http.MethodPut:
			handler.HandleUpdateBlog()(w, r)
		}
	})
	http.HandleFunc("/blogs/{id}/like", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.HandleLikeBlog()(w, r)
		case http.MethodDelete:
			handler.HandleUnLikeBlog()(w, r)
		}
	})
	http.HandleFunc("/blogs/{id}/comments", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.HandleListComment()(w, r)
		case http.MethodPost:
			handler.HandleCreateComment()(w, r)
		case http.MethodDelete:
			handler.HandleDeleteComment()(w, r)
		}
	})
	return nil
}
