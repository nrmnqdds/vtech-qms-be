package internal

import (
	"os"
	"testing"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Setup: Set environment variables for testing
	os.Setenv("PASETO_SECRET_KEY", "faf467491275e04532b6f2b8661b3c93e0896504fc54c129e99b8f10d70a9bf8be7f9772644f4bdb17d8da12d197882711b82f172a579aaff0782a1297f5e6a6")
	os.Setenv("PASETO_PUBLIC_KEY", "be7f9772644f4bdb17d8da12d197882711b82f172a579aaff0782a1297f5e6a6")

	// Run tests
	code := m.Run()

	// Cleanup: Unset environment variables
	os.Unsetenv("PASETO_SECRET_KEY")
	os.Unsetenv("PASETO_PUBLIC_KEY")

	os.Exit(code)
}

func TestGeneratePasetoToken(t *testing.T) {
	testCases := []struct {
		name          string
		userID        string
		expectedError bool
	}{
		{
			name:          "Valid User ID",
			userID:        "user123",
			expectedError: false,
		},
		{
			name:          "Empty User ID",
			userID:        "",
			expectedError: false, // PASETO allows empty strings
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := GeneratePasetoToken(tc.userID)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}

func TestDecodePasetoToken(t *testing.T) {
	testCases := []struct {
		setupToken     func() string
		name           string
		expectedUserID string
		expectedError  bool
	}{
		{
			name: "Valid Token",
			setupToken: func() string {
				token, err := GeneratePasetoToken("user123")
				require.NoError(t, err)
				return token
			},
			expectedError:  false,
			expectedUserID: "user123",
		},
		{
			name: "Expired Token",
			setupToken: func() string {
				// Create a token that's already expired
				token := paseto.NewToken()
				token.SetIssuedAt(time.Now().Add(-48 * time.Hour))
				token.SetNotBefore(time.Now().Add(-48 * time.Hour))
				token.SetExpiration(time.Now().Add(-24 * time.Hour))
				token.SetIssuer("quddus")
				token.SetString("user-id", "expireduser")

				secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(os.Getenv("PASETO_SECRET_KEY"))
				require.NoError(t, err)

				return token.V4Sign(secretKey, nil)
			},
			expectedError: true,
		},
		{
			name: "Invalid Issuer",
			setupToken: func() string {
				token := paseto.NewToken()
				token.SetIssuedAt(time.Now())
				token.SetNotBefore(time.Now())
				token.SetExpiration(time.Now().Add(24 * time.Hour))
				token.SetIssuer("invalid-issuer")
				token.SetString("user-id", "testuser")

				secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(os.Getenv("PASETO_SECRET_KEY"))
				require.NoError(t, err)

				return token.V4Sign(secretKey, nil)
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token := tc.setupToken()

			userID, err := DecodePasetoToken(token)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Empty(t, userID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedUserID, userID)
			}
		})
	}
}

func TestDecodePasetoToken_MissingEnvironmentVariables(t *testing.T) {
	// Temporarily unset environment variables
	os.Unsetenv("PASETO_PUBLIC_KEY")
	defer os.Setenv("PASETO_PUBLIC_KEY", "be7f9772644f4bdb17d8da12d197882711b82f172a579aaff0782a1297f5e6a6")

	token, err := GeneratePasetoToken("user123")
	require.NoError(t, err)

	_, err = DecodePasetoToken(token)
	assert.Error(t, err, "Should return an error when PASETO_PUBLIC_KEY is not set")
}
