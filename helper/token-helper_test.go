package helper

import (
	"testing"
	"time"

	"github.com/gofrs/uuid/v5"
)

func TestSignAndExtract(t *testing.T) {
	secret := "super-secret-123"
	type Claims struct {
		UserID uuid.UUID `json:"user_id"`
	}

	userId, _ := uuid.NewV4()
	expected := Claims{UserID: userId}

	// ---- Sign ----
	token, err := SignJwt(expected, secret, 5*time.Minute)
	if err != nil {
		t.Fatalf("SignJwt failed: %v", err)
	}
	if token == "" {
		t.Fatal("SignJwt returned empty token")
	}

	// ---- Extract ----
	got, err := ExtractJwtClaim[Claims](token, secret)
	if err != nil {
		t.Fatalf("extractJwtClaim failed: %v", err)
	}
	if got.UserID != expected.UserID {
		t.Errorf("extracted claims mismatch: got %+v, want %+v", got, expected)
	}
}
