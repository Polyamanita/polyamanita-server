package lib_test

import (
	"testing"
	"time"

	"github.com/polyamanita/polyamanita-server/src/lib"
	"github.com/stretchr/testify/assert"
)

func TestNewSignedToken(t *testing.T) {
	id := "some id"
	key := []byte("some key")

	t.Run("when the signed token is valid", func(t *testing.T) {
		token, err := lib.NewSignedToken(id, key, time.Hour)
		assert.NoError(t, err)

		claims, valid := lib.ValidateToken(token, key)
		assert.Equal(t, true, valid)
		assert.Equal(t, id, claims.ID)
	})

	t.Run("when the signed token is expired", func(t *testing.T) {
		token, err := lib.NewSignedToken(id, key, -time.Second)
		assert.NoError(t, err)

		claims, valid := lib.ValidateToken(token, key)
		assert.Equal(t, false, valid)
		assert.Nil(t, nil, claims)
	})
}

func TestValidateToken(t *testing.T) {
	id := "some id"
	key := []byte("some key")

	t.Run("when the token is valid", func(t *testing.T) {
		validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6InNvbWUgaWQiLCJpYXQiOjE2NzE0MTI2NTF9.yBLRUqA69p5BeFnZZ-dFRHpMyASy26h5mMa6PT0BJQs"
		wantIssued := int64(1671412651)

		claims, valid := lib.ValidateToken(validToken, key)

		assert.Equal(t, true, valid)
		assert.Equal(t, wantIssued, claims.IssuedAt)
		assert.Equal(t, id, claims.ID)
	})

	t.Run("when the token is invalid", func(t *testing.T) {
		expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6InNvbWUgaWQiLCJleHAiOjE2NzE0MDkxOTksImlhdCI6MTY3MTQxMjc5OX0.JFeEQuRiU2HALwVrmB0tS4ygIsU_NbGkqC-WhrMLuH8"

		claims, valid := lib.ValidateToken(expiredToken, key)

		assert.Equal(t, false, valid)
		assert.Nil(t, claims)
	})
}
