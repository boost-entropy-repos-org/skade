package postgresql

import(
	"testing"	
	"context"

	"github.com/Mindslave/skade/internal/entities"

	"github.com/stretchr/testify/require"
)

func TestStoreFile(t *testing.T) {
	arg := entities.DbStoreFileParams {
		Filename: "testfile",
		Filesize: 100,
		FileExtension: "exe",
	}

	err := testrepo.StoreFile(context.Background(), arg)
	require.NoError(t, err)
}