package managers

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeDB(t *testing.T) {
	testcases := []struct {
		name          string
		db_host       string
		db_port       string
		db_user       string
		db_pass       string
		db_name       string
		expectedError error
	}{
		{
			name:          "TestInitializeDB",
			db_host:       "",
			db_port:       "3306",
			db_user:       "root",
			db_pass:       "root",
			db_name:       "kts",
			expectedError: fmt.Errorf("error initializing datbase connection: environment variables not set"),
		},
		{
			name:          "TestInitializeDB",
			db_host:       "localhost",
			db_port:       "3306",
			db_user:       "root",
			db_pass:       "root",
			db_name:       "kts",
			expectedError: fmt.Errorf("error pinging database: dial tcp [::1]:3306: connect: connection refused"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("DB_HOST", tc.db_host)
			t.Setenv("DB_PORT", tc.db_port)
			t.Setenv("DB_USER", tc.db_user)
			t.Setenv("DB_PASSWORD", tc.db_pass)
			t.Setenv("DB_NAME", tc.db_name)

			_, err := InitializeDB()
			assert.Equal(t, tc.expectedError, err)

			t.Cleanup(func() {
				os.Unsetenv("DB_HOST")
				os.Unsetenv("DB_PORT")
				os.Unsetenv("DB_USER")
				os.Unsetenv("DB_PASSWORD")
				os.Unsetenv("DB_NAME")
			})
		})
	}

}
