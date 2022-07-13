package middleware

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

func AuthorizeKitchenStaff(findKitchenStaffByEmail func(string) (*models.KitchenStaff, error), tokenInBlacklist func(*string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		var kitchenStaff *models.KitchenStaff
		var errors error
		secret := os.Getenv("JWT_SECRET")
		accToken := GetTokenFromHeader(c)
		accessToken, accessClaims, err := AuthorizeToken(&accToken, &secret)
		if err != nil {
			log.Printf("authorize access token errors: %s\n", err.Error())
			RespondAndAbort(c, "", http.StatusUnauthorized, nil, []string{"unauthorized"})
			return
		}

		if tokenInBlacklist(&accessToken.Raw) || IsTokenExpired(accessClaims) {

			RespondAndAbort(c, "", http.StatusBadRequest, nil, []string{"unauthorized"})
		}

		if email, ok := accessClaims["user_email"].(string); ok {
			if kitchenStaff, errors = findKitchenStaffByEmail(email); errors != nil {
				if inactiveErr, ok := err.(helpers.InActiveUserError); ok {
					RespondAndAbort(c, "", http.StatusBadRequest, nil, []string{inactiveErr.Error()})
					return
				}
				log.Printf("find user by email errors: %v\n", err)
				RespondAndAbort(c, "", http.StatusNotFound, nil, []string{"user not found"})
				return
			}
		} else {
			log.Printf("user email is not string\n")
			RespondAndAbort(c, "", http.StatusInternalServerError, nil, []string{"internal server errors"})
			return
		}

		// set the user and token as context parameters.
		c.Set("user", kitchenStaff)
		c.Set("access_token", accessToken.Raw)
		// calling next handler
		c.Next()
	}
}

func AuthorizeFoodBenefactor(findFoodBenefactorByEmail func(string) (*models.FoodBeneficiary, error), tokenInBlacklist func(*string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		var foodBenefactor *models.FoodBeneficiary
		var errors error
		secret := os.Getenv("JWT_SECRET")
		accToken := GetTokenFromHeader(c)
		accessToken, accessClaims, err := AuthorizeToken(&accToken, &secret)
		if err != nil {
			log.Printf("authorize access token errors: %s\n", err.Error())
			RespondAndAbort(c, "", http.StatusUnauthorized, nil, []string{"unauthorized"})
			return
		}

		if tokenInBlacklist(&accessToken.Raw) || IsTokenExpired(accessClaims) {
			c.AbortWithStatusJSON(http.StatusBadRequest, "unauthorized route ")
		}

		if email, ok := accessClaims["user_email"].(string); ok {
			if foodBenefactor, errors = findFoodBenefactorByEmail(email); errors != nil {
				if inactiveErr, ok := err.(helpers.InActiveUserError); ok {
					RespondAndAbort(c, "", http.StatusBadRequest, nil, []string{inactiveErr.Error()})
					return
				}
				log.Printf("find user by email errors: %v\n", err)
				RespondAndAbort(c, "", http.StatusNotFound, nil, []string{"user not found"})
				return
			}
		} else {
			log.Printf("user email is not string\n")
			RespondAndAbort(c, "", http.StatusInternalServerError, nil, []string{"internal server errors"})
			return
		}

		// set the user and token as context parameters.
		c.Set("user", foodBenefactor)
		c.Set("access_token", accessToken.Raw)

		// calling next handler
		c.Next()
	}
}

func AuthorizeAdmin(findAdminByEmail func(string) (*models.Admin, error), tokenInBlacklist func(*string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		var admin *models.Admin
		var errors error
		secret := os.Getenv("JWT_SECRET")
		accToken := GetTokenFromHeader(c)
		accessToken, accessClaims, err := AuthorizeToken(&accToken, &secret)
		if err != nil {
			log.Printf("authorize access token errors: %s\n", err.Error())
			RespondAndAbort(c, "", http.StatusUnauthorized, nil, []string{"unauthorized"})
			return
		}

		if tokenInBlacklist(&accessToken.Raw) || IsTokenExpired(accessClaims) {
			c.AbortWithStatusJSON(http.StatusBadRequest, "unauthorized route ")
		}

		if email, ok := accessClaims["user_email"].(string); ok {
			if admin, errors = findAdminByEmail(email); errors != nil {
				if inactiveErr, ok := err.(helpers.InActiveUserError); ok {
					RespondAndAbort(c, "", http.StatusBadRequest, nil, []string{inactiveErr.Error()})
					return
				}
				log.Printf("find user by email errors: %v\n", err)
				RespondAndAbort(c, "", http.StatusNotFound, nil, []string{"user not found"})
				return
			}
		} else {
			log.Printf("user email is not string\n")
			RespondAndAbort(c, "", http.StatusInternalServerError, nil, []string{"internal server errors"})
			return
		}

		// set the user and token as context parameters.
		c.Set("user", admin)
		c.Set("access_token", accessToken.Raw)

		// calling next handler
		c.Next()
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	passwordHash := string(hashedPassword)
	return passwordHash, nil
}
