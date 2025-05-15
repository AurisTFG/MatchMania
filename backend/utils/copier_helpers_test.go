// nolint
package utils_test

import (
	"MatchManiaAPI/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Source struct {
	Name string
	Age  int
}

type Destination struct {
	Name string
	Age  int
}

type BadDestination struct {
	// incompatible fields that copier cannot copy correctly
	Data chan bool
}

func TestMustCopy_Success(t *testing.T) {
	src := &Source{
		Name: "Aurimas",
		Age:  23,
	}

	result := utils.MustCopy[Destination](src)

	require.NotNil(t, result)
	assert.Equal(t, src.Name, result.Name)
	assert.Equal(t, src.Age, result.Age)
}

func TestMustCopy_Panic(t *testing.T) {
	assert.Panics(t, func() {
		utils.MustCopy[Destination](nil)
	})
}
