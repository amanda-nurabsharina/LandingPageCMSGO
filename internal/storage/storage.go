package storage

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// MaxImageSize is the max accepted upload size (10 MB).
const MaxImageSize = 10 << 20

var allowedExt = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true, ".avif": true, ".bmp": true,
}

// Storage is the "public" disk: files are served under /storage/<path>.
// Equivalent of Laravel's storage/app/public + `php artisan storage:link`.
type Storage struct {
	Root string
}

// New creates the storage root if needed.
func New(root string) (*Storage, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(abs, 0o755); err != nil {
		return nil, err
	}
	return &Storage{Root: abs}, nil
}

// SaveImage validates and stores an uploaded image inside dir, returning the
// relative path (e.g. "hero/3f2a....png") that is stored in the database.
func (s *Storage) SaveImage(fh *multipart.FileHeader, dir string) (string, error) {
	if fh.Size > MaxImageSize {
		return "", fmt.Errorf("ukuran file maksimal %d MB", MaxImageSize>>20)
	}
	ext := strings.ToLower(filepath.Ext(fh.Filename))
	if !allowedExt[ext] {
		return "", errors.New("file harus berupa gambar (jpg, jpeg, png, gif, webp, avif, bmp)")
	}

	src, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	head := make([]byte, 512)
	n, _ := io.ReadFull(src, head)
	ct := http.DetectContentType(head[:n])
	if !strings.HasPrefix(ct, "image/") && ext != ".avif" {
		return "", errors.New("file yang diunggah bukan gambar yang valid")
	}
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	name, err := randomName()
	if err != nil {
		return "", err
	}
	rel := filepath.ToSlash(filepath.Join(dir, name+ext))
	full, err := s.resolve(rel)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return "", err
	}

	dst, err := os.Create(full)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		_ = os.Remove(full)
		return "", err
	}
	return rel, nil
}

// Delete removes a stored file (ignores missing files / invalid paths).
func (s *Storage) Delete(rel string) {
	if rel == "" {
		return
	}
	full, err := s.resolve(rel)
	if err != nil {
		return
	}
	_ = os.Remove(full)
}

// resolve converts a relative storage path into an absolute path, refusing
// anything that escapes the storage root.
func (s *Storage) resolve(rel string) (string, error) {
	clean := filepath.Clean(filepath.FromSlash(rel))
	if filepath.IsAbs(clean) || strings.HasPrefix(clean, "..") {
		return "", errors.New("invalid path")
	}
	full := filepath.Join(s.Root, clean)
	if !strings.HasPrefix(full, s.Root+string(os.PathSeparator)) {
		return "", errors.New("invalid path")
	}
	return full, nil
}

func randomName() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(b)), nil
}

