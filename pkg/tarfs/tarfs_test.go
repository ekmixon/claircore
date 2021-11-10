package tarfs

import (
	"archive/tar"
	"bytes"
	"crypto/sha256"
	"io"
	"io/fs"
	"os"
	"sync"
	"testing"
	"testing/fstest"
)

// TestFS runs some sanity checks on a tar generated from this package's
// directory.
//
// The tar is generated on demand and removed if tests fail, so modifying any
// file in this package *will* cause tests to fail once. Make sure to run tests
// twice if the Checksum tests fail.
func TestFS(t *testing.T) {
	const name = `testdata/fstest.tar`
	checktar(t, name)
	fileset := []string{
		"file.go",
		"parse.go",
		"tarfs.go",
		"tarfs_test.go",
		"testdata/.gitignore",
	}

	t.Run("Single", func(t *testing.T) {
		f, err := os.Open(name)
		if err != nil {
			t.Error(err)
		}
		t.Cleanup(func() {
			if err := f.Close(); err != nil {
				t.Error(err)
			}
		})
		sys, err := New(f)
		if err != nil {
			t.Error(err)
		}

		if err := fstest.TestFS(sys, fileset...); err != nil {
			t.Error(err)
		}
	})

	t.Run("Concurrent", func(t *testing.T) {
		f, err := os.Open(name)
		if err != nil {
			t.Error(err)
		}
		t.Cleanup(func() {
			if err := f.Close(); err != nil {
				t.Error(err)
			}
		})
		sys, err := New(f)
		if err != nil {
			t.Error(err)
		}

		const lim = 8
		var wg sync.WaitGroup
		t.Logf("running %d goroutines", lim)
		wg.Add(lim)
		for i := 0; i < lim; i++ {
			go func() {
				defer wg.Done()
				if err := fstest.TestFS(sys, fileset...); err != nil {
					t.Error(err)
				}
			}()
		}
		wg.Wait()
	})

	t.Run("Sub", func(t *testing.T) {
		f, err := os.Open(name)
		if err != nil {
			t.Error(err)
		}
		t.Cleanup(func() {
			if err := f.Close(); err != nil {
				t.Error(err)
			}
		})
		sys, err := New(f)
		if err != nil {
			t.Error(err)
		}

		sub, err := fs.Sub(sys, "testdata")
		if err != nil {
			t.Error(err)
		}
		if err := fstest.TestFS(sub, ".gitignore"); err != nil {
			t.Error(err)
		}
	})

	t.Run("Checksum", func(t *testing.T) {
		f, err := os.Open(name)
		if err != nil {
			t.Error(err)
		}
		t.Cleanup(func() {
			if err := f.Close(); err != nil {
				t.Error(err)
			}
		})
		sys, err := New(f)
		if err != nil {
			t.Error(err)
		}
		for _, n := range fileset {
			name := n
			t.Run(name, func(t *testing.T) {
				h := sha256.New()
				f, err := os.Open(name)
				if err != nil {
					t.Fatal(err)
				}
				defer f.Close()
				if _, err := io.Copy(h, f); err != nil {
					t.Error(err)
				}
				want := h.Sum(nil)

				h.Reset()
				b, err := fs.ReadFile(sys, name)
				if err != nil {
					t.Error(err)
				}
				if _, err := h.Write(b); err != nil {
					t.Error(err)
				}
				got := h.Sum(nil)

				if !bytes.Equal(got, want) {
					t.Errorf("got: %x, want: %x", got, want)
				}
			})
		}
	})
}

func checktar(t *testing.T, name string) {
	t.Helper()
	out, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	tw := tar.NewWriter(out)
	defer tw.Close()

	in := os.DirFS(".")
	if err := fs.WalkDir(in, ".", mktar(t, in, tw)); err != nil {
		t.Fatal(err)
	}
}

func mktar(t *testing.T, in fs.FS, tw *tar.Writer) fs.WalkDirFunc {
	return func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		switch {
		case d.Name() == "fstest.tar":
			return nil
		case d.Name() == "." && d.IsDir():
			return nil
		default:
		}
		t.Logf("adding %q", p)
		i, err := d.Info()
		if err != nil {
			return err
		}
		h, err := tar.FileInfoHeader(i, "")
		if err != nil {
			return err
		}
		h.Name = p
		if err := tw.WriteHeader(h); err != nil {
			return err
		}
		if i.IsDir() {
			return nil
		}
		f, err := in.Open(p)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}
		return nil
	}
}
