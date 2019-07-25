package account

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Account struct {
	UUID      string  `json:"uuid"`
	CPF       string  `json:"cpf"`
	Name      string  `json:"name"`
	Surname   string  `json:"surname"`
	Segment   string  `json:"segment"`
	Balance   float64 `json:"balance"`
	CreatedAt int64   `json:"createdat"`
	UpdatedAt int64   `json:"updateat"`
}

// Generate a random Bank Account
func (r *resource) random(c echo.Context) error {
	account, err := NewRandomAccount()
	if err != nil {
		r.log.Error("Error generating random account")
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	if account != nil && account.isBlacklisted() {
		return c.JSON(http.StatusForbidden, account)
	}
	if err = r.db.create(account); err != nil {
		r.log.Error("Error inserting account into MongoDB")
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, account)
}

/**
 * Create a Bank Account
 * @param Account
 */
func (r *resource) new(c echo.Context) error {
	// Bind Account Payload
	account := new(Account)
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't bind JSON payload to account Struct")
	}

	// Verify Data
	if !account.isValid() {
		return c.JSON(http.StatusBadRequest, "This account is invalid")
	}
	if account.isBlacklisted() {
		return c.JSON(http.StatusForbidden, "This account is blacklisted")
	}

	// Complete Creation Date and UUID
	account.UUID = uuid.New().String()
	account.CreatedAt = time.Now().Unix()
	account.UpdatedAt = time.Now().Unix()

	// Create Account
	if err := r.db.create(account); err != nil {
		r.log.Error("Error updating account into MongoDB")
		c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, account)
}

/**
 * Read a Bank Account
 * @param :uuid string
 */
func (r *resource) get(c echo.Context) error {
	result := new(Account)
	if err := r.db.read("uuid", c.Param("uuid"), result); err != nil {
		r.log.Error("Error inserting account into DynamoDB")
		c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

/**
 * Read a Bank Account
 * @param :uuid string
 */
func (r *resource) getAll(c echo.Context) error {
	all, err := r.db.readAll()
	if err != nil {
		r.log.Error("Error inserting account into DynamoDB")
		c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, all)
}

/**
 * Update a Bank Account
 * @param Account
 */
func (r *resource) set(c echo.Context) error {
	account := new(Account)
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't bind JSON payload to account Struct")
	}
	updated, err := r.db.update("uuid", c.Param("uuid"), account.Balance)
	if err != nil {
		r.log.Error("Error updating account into MongoDB")
		c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, updated)
}

/**
 * Delete a Bank Account
 * @param :uuid string
 */
func (r *resource) delete(c echo.Context) error {
	result := new(Account)
	deleted, err := r.db.delete("uuid", c.Param("uuid"), result)
	if err != nil {
		r.log.Error("Error inserting account into DynamoDB")
		c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, deleted)
}

func NewRandomAccount() (*Account, error) {
	account := new(Account)
	account.UUID = uuid.New().String()
	account.Name = randomdata.FirstName(2)
	account.Surname = randomdata.LastName()
	account.CreatedAt = time.Now().Unix()
	account.UpdatedAt = time.Now().Unix()
	segment := rand.Intn(3)
	segmentType := ""
	switch segment {
	case 0:
		segmentType = "Varejo"
	case 1:
		segmentType = "Uniclass"
	case 2:
		segmentType = "Personalité"
	}
	account.Segment = segmentType
	integerPart := strconv.Itoa(rand.Intn(20000))
	decimal := rand.Intn(99)
	decimalString := ""
	if decimal < 10 {
		decimalString = "0" + strconv.Itoa(decimal)
	} else {
		decimalString = strconv.Itoa(decimal)
	}
	consolidatedPosition, err := strconv.ParseFloat(integerPart+"."+decimalString, 64)
	if err != nil {
		return nil, err
	}
	account.Balance = consolidatedPosition
	return account, nil
}

func (a *Account) isBlacklisted() bool {
	return false
}

func (a *Account) isValid() bool {
	return true
}
