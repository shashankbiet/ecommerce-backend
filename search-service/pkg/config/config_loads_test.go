package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadSecretFilesSetsValueFromFile(t *testing.T) {
	tempDir := t.TempDir()
	secretFile := filepath.Join(tempDir, "elastic_password.txt")
	if err := os.WriteFile(secretFile, []byte(" search-secret \n"), 0o600); err != nil {
		t.Fatalf("write secret file: %v", err)
	}

	t.Setenv("SEARCH_TEST_SECRET", "")
	t.Setenv("SEARCH_TEST_SECRET_FILE", secretFile)

	if err := loadSecretFiles(); err != nil {
		t.Fatalf("loadSecretFiles() error = %v", err)
	}

	if got := os.Getenv("SEARCH_TEST_SECRET"); got != "search-secret" {
		t.Fatalf("SEARCH_TEST_SECRET = %q, want %q", got, "search-secret")
	}
}

func TestLoadSecretFilesDoesNotOverrideExistingValue(t *testing.T) {
	tempDir := t.TempDir()
	secretFile := filepath.Join(tempDir, "secret.txt")
	if err := os.WriteFile(secretFile, []byte("from-file"), 0o600); err != nil {
		t.Fatalf("write secret file: %v", err)
	}

	t.Setenv("SEARCH_TEST_SECRET_EXISTS", "already-set")
	t.Setenv("SEARCH_TEST_SECRET_EXISTS_FILE", secretFile)

	if err := loadSecretFiles(); err != nil {
		t.Fatalf("loadSecretFiles() error = %v", err)
	}

	if got := os.Getenv("SEARCH_TEST_SECRET_EXISTS"); got != "already-set" {
		t.Fatalf("SEARCH_TEST_SECRET_EXISTS = %q, want %q", got, "already-set")
	}
}

func TestLoadSecretFilesReturnsErrorForMissingFile(t *testing.T) {
	t.Setenv("SEARCH_TEST_MISSING_FILE", filepath.Join(t.TempDir(), "missing.txt"))
	t.Setenv("SEARCH_TEST_MISSING", "")

	err := loadSecretFiles()
	if err == nil {
		t.Fatal("loadSecretFiles() error = nil, want non-nil")
	}
	if !strings.Contains(err.Error(), "SEARCH_TEST_MISSING_FILE") {
		t.Fatalf("error %q does not mention env var", err.Error())
	}
}
