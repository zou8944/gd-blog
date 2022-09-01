package route

import (
	"encoding/json"
	"gd-blog/src/domain/service"
	"gd-blog/src/gdlog"
	"net/http"
	"strconv"
)

type BlogHandler struct {
	blogDomainService service.BlogDomainService
}

func NewBlogHandler(blogDomainService service.BlogDomainService) BlogHandler {
	return BlogHandler{blogDomainService: blogDomainService}
}

func (bs *BlogHandler) HandleListBlog() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		separateId, _ := strconv.Atoi(r.FormValue("separateId"))
		limit, _ := strconv.Atoi(r.FormValue("limit"))
		blogs, err := bs.blogDomainService.ListBlog(separateId, limit)
		if err != nil {
			gdlog.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bytes, err := json.Marshal(blogs)
		if err != nil {
			gdlog.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(bytes)
	}
}

func (bs *BlogHandler) HandleCreateBlog() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (bs *BlogHandler) HandleGetBlog() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (bs *BlogHandler) HandleUpdateBlog() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (bs *BlogHandler) HandleLikeBlog() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (bs *BlogHandler) HandleUnLikeBlog() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (bs *BlogHandler) HandleCreateComment() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (bs *BlogHandler) HandleListComment() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (bs *BlogHandler) HandleDeleteComment() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
