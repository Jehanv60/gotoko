package controllers

import (
	"net/http"

	"github.com/Jehanv60/gotoko/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Product(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})
	productmodel := models.Product{}
	product, err := productmodel.Getproduct(server.DB)
	if err != nil {
		return
	}
	_ = render.HTML(w, http.StatusOK, "product", map[string]interface{}{
		"product": product,
	})
}
