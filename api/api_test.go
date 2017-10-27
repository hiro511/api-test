package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hiro511/api-test/config"
)

func BenchmarkTimetable(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(getTimetable))
	defer ts.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := http.Get(ts.URL + "/timetable")
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			log.Fatalf("unexpected status code code=%d", res.StatusCode)
		}
	}
}

func TestGetTimetable(t *testing.T) {
	config.SetDBBasePath("../")
	ts := httptest.NewServer(http.HandlerFunc(getTimetable))
	defer ts.Close()

	res, err := httpGet(ts.URL+"/timetable", http.Header{})
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code, code=%d, body=%s", res.StatusCode, string(body))
	}

	log.Printf("content-length = %d, content-encoding = %s",
		len(body), res.Header.Get("Content-Encoding"))
}

func httpGet(path string, header http.Header) (*http.Response, error) {
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: header,
	}
	return http.DefaultClient.Do(req)
}
