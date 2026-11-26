package sqlite

import (
	"path/filepath"
	"testing"
)

func TestNewDB(t *testing.T) {
	sName := filepath.Join(t.TempDir(), "cv.db")

	db, err := NewDB(sName)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = db.Close()
	})
}
