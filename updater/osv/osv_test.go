package osv

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/quay/zlog"

	"github.com/quay/claircore"
	"github.com/quay/claircore/libvuln/driver"
)

func TestFetch(t *testing.T) {
	srv := httptest.NewServer(&apiStub{t, ""})
	defer srv.Close()
	ctx := zlog.Test(context.Background(), t)

	f := Factory{}
	cfgFunc := func(v interface{}) error {
		cfg := v.(*FactoryConfig)
		cfg.URL = srv.URL
		return nil
	}
	if err := f.Configure(ctx, cfgFunc, srv.Client()); err != nil {
		t.Error(err)
	}

	s, err := f.UpdaterSet(ctx)
	if err != nil {
		t.Error(err)
	}
	if len(s.Updaters()) == 0 {
		t.Errorf("expected more than 0 updaters")
	}

	for _, u := range s.Updaters() {
		rc, fp, err := u.Fetch(ctx, driver.Fingerprint(""))
		if err != nil {
			t.Error(err)
		}
		_ = fp
		if rc != nil {
			rc.Close()
		}

	}
}

type apiStub struct {
	*testing.T
	path string
}

func (a *apiStub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Logf("req: %s", r.RequestURI)
	sys := os.DirFS(filepath.Join("testdata", a.path))
	p := r.URL.Path
	switch {
	case p == "/ecosystems.txt":
		out := bufio.NewWriter(w)
		defer out.Flush()
		fmt.Fprintln(out, "testing_ecosystem")
		ms, err := fs.Glob(sys, "*.zip")
		if err != nil {
			panic(err) // can only ever be ErrBadPatern
		}
		for _, m := range ms {
			fmt.Fprintln(out, strings.TrimSuffix(m, ".zip"))
		}
	case strings.HasSuffix(p, "all.zip"):
		w.WriteHeader(http.StatusOK)
		n := strings.ToLower(path.Dir(p)[1:]) + ".zip"
		a.Logf("serving %q", n)
		if f, err := sys.Open(n); errors.Is(err, nil) {
			defer f.Close()
			if _, err := io.Copy(w, f); err != nil {
				a.Error(err)
			}
			return
		}
		z := zip.NewWriter(w)
		if err := z.SetComment("empty zip"); err != nil {
			a.Error(err)
		}
		if err := z.Close(); err != nil {
			a.Error(err)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func TestParse(t *testing.T) {
	srv := httptest.NewServer(&apiStub{t, ""})
	defer srv.Close()
	ctx := zlog.Test(context.Background(), t)

	f := Factory{}
	cfgFunc := func(v interface{}) error {
		cfg := v.(*FactoryConfig)
		cfg.URL = srv.URL
		return nil
	}
	if err := f.Configure(ctx, cfgFunc, srv.Client()); err != nil {
		t.Error(err)
	}
	s, err := f.UpdaterSet(ctx)
	if err != nil {
		t.Error(err)
	}
	if len(s.Updaters()) == 0 {
		t.Errorf("expected more than 0 updaters")
	}

	for _, u := range s.Updaters() {
		rc, _, err := u.Fetch(ctx, driver.Fingerprint(""))
		if err != nil {
			t.Error(err)
		}
		defer rc.Close()
		vs, err := u.Parse(ctx, rc)
		if err != nil {
			t.Error(err)
		}
		t.Logf("parsed %d vulnerabilities", len(vs))
		if len(vs) != 0 {
			t.Log("first one:")
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			enc.SetIndent("", "\t")
			if err := enc.Encode(vs[0]); err != nil {
				t.Error(err)
			}
			t.Log(buf.String())
		}
	}
}

var severityTestCases = []struct {
	name                       string
	a                          *advisory
	expectedNormalizedSeverity claircore.Severity
	expectedSeverity           string
}{
	{
		name: "CVSS V3 HIGH",
		a: &advisory{
			ID: "test1",
			Severity: []severity{
				{
					Type:  "CVSS_V3",
					Score: "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N",
				},
			},
			Affected: []affected{
				{
					Package: _package{
						Ecosystem: "go",
						Name:      "something",
					},
					Ranges: []_range{
						{
							Type: "ECOSYSTEM",
							Events: []rangeEvent{
								{
									Introduced: "0.1",
									Fixed:      "0.4",
								},
							},
						},
					},
				},
			},
		},
		expectedNormalizedSeverity: claircore.High,
		expectedSeverity:           "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N",
	},
	{
		name: "CVSS V2 MEDIUM",
		a: &advisory{
			ID: "test2",
			Severity: []severity{
				{
					Type:  "CVSS_V2",
					Score: "AV:L/AC:H/Au:N/C:C/I:C/A:C",
				},
			},
			Affected: []affected{
				{
					Package: _package{
						Ecosystem: "go",
						Name:      "something",
					},
					Ranges: []_range{
						{
							Type: "ECOSYSTEM",
							Events: []rangeEvent{
								{
									Introduced: "0.1",
									Fixed:      "0.4",
								},
							},
						},
					},
				},
			},
		},
		expectedNormalizedSeverity: claircore.Medium,
		expectedSeverity:           "AV:L/AC:H/Au:N/C:C/I:C/A:C",
	},
	{
		name: "database_specific moderate",
		a: &advisory{
			ID: "test2",
			Affected: []affected{
				{
					Package: _package{
						Ecosystem: "go",
						Name:      "something",
					},
					Ranges: []_range{
						{
							Type: "ECOSYSTEM",
							Events: []rangeEvent{
								{
									Introduced: "0.1",
									Fixed:      "0.4",
								},
							},
						},
					},
				},
			},
			Database: json.RawMessage([]byte(`{"severity":"moderate"}`)),
		},
		expectedNormalizedSeverity: claircore.Medium,
		expectedSeverity:           "moderate",
	},
	{
		name: "CVSS V3 HIGH and database_specific moderate",
		a: &advisory{
			ID: "test2",
			Severity: []severity{
				{
					Type:  "CVSS_V3",
					Score: "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N",
				},
			},
			Affected: []affected{
				{
					Package: _package{
						Ecosystem: "go",
						Name:      "something",
					},
					Ranges: []_range{
						{
							Type: "ECOSYSTEM",
							Events: []rangeEvent{
								{
									Introduced: "0.1",
									Fixed:      "0.4",
								},
							},
						},
					},
				},
			},
			Database: json.RawMessage([]byte(`{"severity":"moderate"}`)),
		},
		expectedNormalizedSeverity: claircore.High,
		expectedSeverity:           "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N",
	},
}

func TestSeverityParsing(t *testing.T) {
	ctx := zlog.Test(context.Background(), t)

	for _, tt := range severityTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ecs := newECS("test")

			err := ecs.Insert(ctx, nil, "", tt.a)
			if err != nil {
				t.Error("got error Inserting advisory", err)
			}
			if len(ecs.Vulnerability) != 1 {
				t.Errorf("should have one vulnerability but got %d", len(ecs.Vulnerability))
			}
			v := ecs.Vulnerability[0]
			if v.NormalizedSeverity != tt.expectedNormalizedSeverity {
				t.Errorf("expected severity %q but got %q", tt.expectedNormalizedSeverity, v.NormalizedSeverity)
			}
			if v.Severity != tt.expectedSeverity {
				t.Errorf("expected severity %q but got %q", tt.expectedSeverity, v.Severity)
			}

		})
	}
}
