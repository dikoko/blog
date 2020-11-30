package user

import(
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dikoko/blog/depinv/v2/entity"
	"github.com/dikoko/blog/depinv/v2/strg"
)

func TestUserManager(t *testing.T) {
	t.Parallel()

	var (
		err error
		storage = strg.NewStorage()
		userMan = NewUserManager(storage)
		testUser = &entity.User{
			ID: "1",
			Name: "Diko",
		}
	)

	err = userMan.AddUser(testUser)
	assert.NoError(t, err)

	user, err := userMan.GetUser(testUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, testUser, user)

	err = userMan.DeleteUser(testUser.ID)
	assert.NoError(t, err)
}
