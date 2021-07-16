package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/quang2906/book_store_be/controller"
)

func ConfigCategoryRouter(router *mux.Router) {
	category := router.PathPrefix("/v1/categories").Subrouter()
	category.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllCategory)
	category.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetCategoryById)
	category.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateCategory)
	category.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateCategory)
	category.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteCategoryById)

}

func ConfigProductsRouter(router *mux.Router) {
	product := router.PathPrefix("/v1/products").Subrouter()

	product.Path("").Methods(http.MethodGet).HandlerFunc(controller.SearchProduct)
	product.Path("/sort").Methods(http.MethodGet).HandlerFunc(controller.SortProduct)
	// product.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllProducts)
	product.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetProductById)
	product.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateProduct)
	product.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateProductById)
	product.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteProductById)

}

func ConfigUserRouter(router *mux.Router) {
	user := router.PathPrefix("/v1/users").Subrouter()
	user.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllUsers)
	user.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetUserById)
	user.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateUser)
	user.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateUser)
	user.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteUserById)

}
func ConfigOrderRouter(router *mux.Router) {
	order := router.PathPrefix("/v1/orders").Subrouter()
	order.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllOrder)
	order.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetOrderById)
	order.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateOrder)
	order.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateOrderById)
	order.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteOrderById)
}
