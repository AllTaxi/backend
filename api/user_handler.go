package api

import (
	"encoding/json"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gitlab.com/golang-team-template/monolith/model"
	"gitlab.com/golang-team-template/monolith/pkg/utils"
	"go.uber.org/zap"
	// rds "gitlab.com/golang-team-template/monolith/storage/redis"
)

type request struct {
	Email string `json:"email"`
}
type response struct {
	ID string `json:"id"`
}

// SendCode method for verify user.
// @Description Send verification code for a new user.
// @Summary send verification a new user
// @Tags register
// @Accept json
// @Produce json
// @Param register body request true "register"
// @Success 200 {object} model.UserRegister
// @Router /send-code/ [post]
func (api *api) sendCode(w http.ResponseWriter, r *http.Request) {

	var body request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, "error parsing request body", http.StatusBadRequest)
		return
	}

	// code, err := utils.SendEmail(email)
	// if err != nil {
	// 	fmt.Println(err, code)
	// 	return nil, err
	// }
	// fmt.Println("Send email work")
	err := api.redisStorage.SetWithTTL(body.Email, "7777", 600)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated)
}

// verifyUser method to Verify a user.
// @Description Verify a user.
// @Summary verify a user
// @Tags register
// @Accept json
// @Produce json
// @Param email path string true "email"
// @Param code path string true "code"
// @Success 200 {object} response
// @Router /verify/{email}/{code}/ [get]
func (api *api) verifyUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	// code := params["code"]

	val, err := api.redisStorage.Get(email)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	checkValue, err := redis.String(val, err)
	if err != nil {
		api.logger.Error("failed to get value from redis", zap.Error(err))
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if checkValue != "7777" {
		writeError(w, "Wrong verification code", http.StatusBadRequest)
		return
	}
	err = api.redisStorage.SetWithTTL(email, "verified", 600)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK)
}

// registerUser method to create a new user.
// @Description Create a new user.
// @Summary creates a new user
// @Tags register
// @Accept json
// @Produce json
// @Param register body model.UserRequest true "register"
// @Success 200 {object} model.User
// @Router /users/register/ [post]
func (api *api) registerUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body model.User
	claimsMap := make(map[string]string, 0)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, "error parsing request body", http.StatusBadRequest)
		return
	}
	val, err := api.redisStorage.Get(body.Email)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	checkValue, err := redis.String(val, err)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if checkValue != "verified" {
		writeError(w, "You are not verified user", http.StatusBadRequest)
		return
	}

	body.ID = uuid.New().String()

	claimsMap["role"] = "user"
	tokens, err := utils.GenerateNewTokens(body.ID, claimsMap)

	if err != nil {
		writeError(w, "error token", http.StatusInternalServerError)
	}

	body.AccessToken = tokens.Access
	user, err := api.userService.SignUp(ctx, &body)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, user)
}

//_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
