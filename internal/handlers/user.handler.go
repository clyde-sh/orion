package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucsky/cuid"

	"github.com/clyde-sh/orion/internal/database"
	"github.com/clyde-sh/orion/internal/security"
	"github.com/clyde-sh/orion/internal/utils"
	shared "github.com/clyde-sh/orion/shared/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

func (h *HandlersCtx) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetClientIP(r)

	if ip != "" && !h.ratelimit.UserRegisterRateLimit.Consume(ip) {
		utils.WriteErrorJSON(w, ErrRateLimitedError, http.StatusTooManyRequests)
		return
	}

	var body shared.UserCreateUserRequestDto
	err := utils.ReadJSON(w, r, &body)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.IsValidEmail(body.Email)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.IsValidPassword(body.Password)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.IsValidName(body.Name)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	passwordHash, err := security.HashValue(body.Password)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	accountRecoveryCode, err := security.GenerateSecureCode()
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	accountRecoveryCodeHash, err := security.HashValue(accountRecoveryCode)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	id := cuid.New()

	slog.Debug(fmt.Sprintf("Inserting a new user of id %s...", id))

	user, err := h.queries.InsertUser(r.Context(), database.InsertUserParams{
		ID:           id,
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: passwordHash,
		Role:         string(RoleUser),
		RecoveryCode: accountRecoveryCodeHash,
	})
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, shared.UserCreateUserResponseDto{
		User: &shared.User{
			Id:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Role:            user.Role,
			IsEmailVerified: user.IsEmailVerified.Bool,
			CreatedAt:       timestamppb.New(user.CreatedAt.Time),
			UpdatedAt:       timestamppb.New(user.UpdatedAt.Time),
		},
	}, http.StatusCreated)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}
}

func (h *HandlersCtx) HandleAllGetUsers(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetClientIP(r)

	if ip != "" && !h.ratelimit.UserGetAllUsersRateLimit.Consume(ip) {
		utils.WriteErrorJSON(w, ErrRateLimitedError, http.StatusTooManyRequests)
		return
	}

	users, err := h.queries.ListUsers(r.Context())
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	var protoUsers []*shared.User
	for _, user := range users {
		pu := &shared.User{
			Id:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Role:            user.Role,
			IsEmailVerified: user.IsEmailVerified.Bool,
			Is_2FaEnabled:   user.Is2faEnabled.Bool,
			CreatedAt:       timestamppb.New(user.CreatedAt.Time),
			UpdatedAt:       timestamppb.New(user.UpdatedAt.Time),
		}
		protoUsers = append(protoUsers, pu)
	}

	err = utils.WriteJSON(w, &shared.UserGetAllUsersResponseDto{
		Users: protoUsers,
	}, http.StatusOK)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}
}

func (h *HandlersCtx) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetClientIP(r)

	if ip != "" && !h.ratelimit.UserGetUserByIdRateLimit.Consume(ip) {
		utils.WriteErrorJSON(w, ErrRateLimitedError, http.StatusTooManyRequests)
		return
	}

	id := chi.URLParam(r, "id")
	user, err := h.queries.FindUserById(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteErrorJSON(w, ErrNotFoundError, http.StatusNotFound)
		} else {
			utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		}
		return
	}

	err = utils.WriteJSON(w, shared.UserCreateUserResponseDto{
		User: &shared.User{
			Id:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Role:            user.Role,
			IsEmailVerified: user.IsEmailVerified.Bool,
			CreatedAt:       timestamppb.New(user.CreatedAt.Time),
			UpdatedAt:       timestamppb.New(user.UpdatedAt.Time),
		},
	}, http.StatusCreated)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}
}

func (h *HandlersCtx) HandleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetClientIP(r)

	if ip != "" && !h.ratelimit.UserUpdateUserRateLimit.Consume(ip) {
		utils.WriteErrorJSON(w, ErrRateLimitedError, http.StatusTooManyRequests)
		return
	}

	id := chi.URLParam(r, "id")

	var body shared.UserUpdateUserRequestDto
	err := utils.ReadJSON(w, r, &body)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedUser, err := h.queries.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:           id,
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: body.PasswordHash,
		RecoveryCode: body.RecoveryCode,
		Role:         body.Role,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteErrorJSON(w, ErrNotFoundError, http.StatusNotFound)
		} else {
			utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		}
		return
	}

	err = utils.WriteJSON(w, shared.UserUpdateUserResponseDto{
		User: &shared.User{
			Id:              updatedUser.ID,
			Name:            updatedUser.Name,
			Email:           updatedUser.Email,
			Role:            updatedUser.Role,
			IsEmailVerified: updatedUser.IsEmailVerified.Bool,
			CreatedAt:       timestamppb.New(updatedUser.CreatedAt.Time),
			UpdatedAt:       timestamppb.New(updatedUser.UpdatedAt.Time),
		},
	}, http.StatusOK)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}
}

func (h *HandlersCtx) HandleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetClientIP(r)

	if ip != "" && !h.ratelimit.UserDeleteUserRateLimit.Consume(ip) {
		utils.WriteErrorJSON(w, ErrRateLimitedError, http.StatusTooManyRequests)
		return
	}

	id := chi.URLParam(r, "id")
	deletedUser, err := h.queries.DeleteUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteErrorJSON(w, ErrNotFoundError, http.StatusNotFound)
		} else {
			utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		}
		return
	}

	err = utils.WriteJSON(w, &shared.UserDeleteUserResponseDto{
		User: &shared.User{
			Id:              deletedUser.ID,
			Name:            deletedUser.Name,
			Email:           deletedUser.Email,
			IsEmailVerified: deletedUser.IsEmailVerified.Bool,
			Is_2FaEnabled:   deletedUser.Is2faEnabled.Bool,
			CreatedAt:       timestamppb.New(deletedUser.CreatedAt.Time),
			UpdatedAt:       timestamppb.New(deletedUser.UpdatedAt.Time),
		},
	}, http.StatusOK)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}
}
