package sqlite

import (
	"context"
	"strings"
	"testing"
)

func TestGetResume(t *testing.T) {
	db := newTestDB(t)
	ctx := context.Background()

	resume, err := db.GetResume(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}

	if !resume.Residence.Valid {
		t.Fatal("expected resume residence to be valid")
	}

	if !strings.Contains(resume.Residence.Value, "Астана") {
		t.Fatal("expected resume residence to contain 'Астана'")
	}
}
