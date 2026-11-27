package sqlite

import "testing"

func newTestDB(t *testing.T) Db {
	t.Helper()

	db, err := NewDB("testdata/cv_test.db")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("close test database: %v", err)
		}
	})

	return db
}
