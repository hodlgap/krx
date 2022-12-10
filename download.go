package krx

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

const (
	dateStringFormat = "20060102"
)

func toDateStr(t time.Time) string {
	return t.Format(dateStringFormat)
}

func closeResponseBody(resp *http.Response) {
	if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
		log.Error(err)
	}
	if err := resp.Body.Close(); err != nil {
		log.Error(err)
	}
}

func GetStockFile(downloadCode string) ([]byte, error) {
	resp, err := http.PostForm("http://data.krx.co.kr/comm/fileDn/download_csv/download.cmd", url.Values{
		"code": {downloadCode},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer closeResponseBody(resp)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("invalid status code: %d", resp.StatusCode)
	}

	file, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return file, nil
}

func GetDownloadCode(when time.Time) (string, error) {
	resp, err := http.PostForm("http://data.krx.co.kr/comm/fileDn/GenerateOTP/generate.cmd", url.Values{
		"locale":      {"ko_KR"},
		"mktId":       {"ALL"},
		"trdDd":       {toDateStr(when)},
		"share":       {"1"},
		"money":       {"1"},
		"csvxls_isNo": {"false"},
		"name":        {"fileDown"},
		"url":         {"dbms/MDC/STAT/standard/MDCSTAT01501"},
	})
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer closeResponseBody(resp)
	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("invalid status code: %d", resp.StatusCode)
	}

	code, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(code), nil
}
