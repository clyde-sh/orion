package security

import (
	"time"
)

type AppRateLimiters struct {
	UserRegisterRateLimit    TokenBucketRateLimit
	UserGetAllUsersRateLimit TokenBucketRateLimit
	UserGetUserByIdRateLimit TokenBucketRateLimit
	UserUpdateUserRateLimit  TokenBucketRateLimit
	UserDeleteUserRateLimit  TokenBucketRateLimit
}

func InitializeAppRateLimiters() AppRateLimiters {
	return AppRateLimiters{
		UserRegisterRateLimit:    NewTokenBucketRateLimit(5, 10*time.Second),
		UserGetAllUsersRateLimit: NewTokenBucketRateLimit(50, 1*time.Minute),
		UserGetUserByIdRateLimit: NewTokenBucketRateLimit(50, 1*time.Minute),
		UserUpdateUserRateLimit:  NewTokenBucketRateLimit(5, 2*time.Minute),
		UserDeleteUserRateLimit:  NewTokenBucketRateLimit(3, 1*time.Hour),
	}
}
