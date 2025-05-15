// nolint
package utils_test

import (
	"MatchManiaAPI/constants"
	"MatchManiaAPI/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestContext() *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(nil)
	return ctx
}

func TestGetParamId_Success(t *testing.T) {
	ctx := createTestContext()
	testId := uuid.New()
	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "id",
		Value: testId.String(),
	})

	id, err := utils.GetParamId(ctx, "id")

	assert.NoError(t, err)
	assert.Equal(t, testId, id)
}

func TestGetParamId_InvalidUUID(t *testing.T) {
	ctx := createTestContext()
	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "id",
		Value: "invalid-uuid",
	})

	id, err := utils.GetParamId(ctx, "id")

	assert.Error(t, err)
	assert.Equal(t, uuid.Nil, id)
}

func TestMustGetRequestingUserId_Success(t *testing.T) {
	ctx := createTestContext()
	testId := uuid.New()
	ctx.Set(constants.RequestingUserIdKey, testId)

	result := utils.MustGetRequestingUserId(ctx)

	assert.Equal(t, testId, result)
}

func TestMustGetRequestingUserId_MissingKey(t *testing.T) {
	ctx := createTestContext()

	assert.PanicsWithValue(t, "Requesting user ID not found in context", func() {
		utils.MustGetRequestingUserId(ctx)
	})
}

func TestMustGetRequestingUserId_WrongType(t *testing.T) {
	ctx := createTestContext()
	ctx.Set(constants.RequestingUserIdKey, "not-a-uuid")

	assert.PanicsWithValue(t, "Requesting user ID is not of type uuid.UUID", func() {
		utils.MustGetRequestingUserId(ctx)
	})
}

func TestGetRequestingUserPermissions_Success(t *testing.T) {
	ctx := createTestContext()
	permissions := []string{"perm1", "perm2"}
	ctx.Set(constants.RequestingUserPermissionsKey, permissions)

	result := utils.GetRequestingUserPermissions(ctx)

	assert.Equal(t, permissions, result)
}

func TestGetRequestingUserPermissions_MissingKey(t *testing.T) {
	ctx := createTestContext()

	result := utils.GetRequestingUserPermissions(ctx)

	assert.Nil(t, result)
}

func TestGetRequestingUserPermissions_WrongType(t *testing.T) {
	ctx := createTestContext()
	ctx.Set(constants.RequestingUserPermissionsKey, "not-a-slice")

	result := utils.GetRequestingUserPermissions(ctx)

	assert.Nil(t, result)
}

func TestSetRequestingUserId(t *testing.T) {
	ctx := createTestContext()
	testId := uuid.New()

	utils.SetRequestingUserId(ctx, testId)

	val, exists := ctx.Get(constants.RequestingUserIdKey)
	assert.True(t, exists)
	assert.Equal(t, testId, val)
}

func TestSetRequestingUserPermissions(t *testing.T) {
	ctx := createTestContext()
	permissions := []string{"perm1", "perm2"}

	utils.SetRequestingUserPermissions(ctx, permissions)

	val, exists := ctx.Get(constants.RequestingUserPermissionsKey)
	assert.True(t, exists)
	assert.Equal(t, permissions, val)
}

func TestMustSetTrustedProxies_Success(t *testing.T) {
	server := gin.New()
	trustedProxies := []string{"192.168.0.1", "192.168.0.2"}

	assert.NotPanics(t, func() {
		utils.MustSetTrustedProxies(server, trustedProxies)
	})
}

func TestMustSetTrustedProxies_Panic(t *testing.T) {
	server := gin.New()
	invalidProxies := []string{"invalid-proxy"}

	assert.Panics(t, func() {
		utils.MustSetTrustedProxies(server, invalidProxies)
	})
}
