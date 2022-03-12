package api

import (
	"gitlab.com/golang-team-template/monolith/service"
	"gitlab.com/golang-team-template/monolith/storage/redis"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type routes struct {
	root    *mux.Router
	apiRoot *mux.Router
}

type api struct {
	routes        *routes
	userService   service.UserService
	deviceService service.DeviceService
	redisStorage  redis.InMemoryStorage
	logger        *zap.Logger
}

//Init starts routes
func Init(
	root *mux.Router,
	userService service.UserService,
	deviceService service.DeviceService,
	redisStorage redis.InMemoryStorage,
	logger *zap.Logger) {

	r := routes{
		root:    root,
		apiRoot: root.PathPrefix("/api").Subrouter(),
	}

	api := api{
		routes:        &r,
		userService:   userService,
		deviceService: deviceService,
		redisStorage:  redisStorage,
	}

	api.routes.root.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	api.initUser()
	api.initDevice()
}

func (api *api) initUser() {
	// api.routes.apiRoot.HandleFunc("/users/", api.sendCode).Methods("GET")
	api.routes.apiRoot.HandleFunc("/send-code/", api.sendCode).Methods("POST")
	api.routes.apiRoot.HandleFunc("/verify/{email}/{code}/", api.verifyUser).Methods("GET")
	api.routes.apiRoot.HandleFunc("/users/register/", api.registerUser).Methods("POST")
}

func (api *api) initDevice() {
	api.routes.apiRoot.HandleFunc("/devices/", api.registerDevice).Methods("POST")
	api.routes.apiRoot.HandleFunc("/devices/", api.registerDevice).Methods("GET")
}
