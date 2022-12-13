package krx

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

func ParseStockBuf(buf []byte) ([]*Stock, error) {
	decoded, err := decodeEUCKR(buf)
	if err != nil {
		return nil, err
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ','

		return r // Allows use pipe as delimiter
	})

	var rows []*Stock
	if err := gocsv.UnmarshalBytes(decoded, &rows); err != nil {
		return nil, errors.WithStack(err)
	}

	return rows, nil
}

func decodeEUCKR(from []byte) ([]byte, error) {
	// NOTE: 한국거래소 파일은 기본적으로 EUC-KR 포맷을 사용함.

	r := transform.NewReader(bytes.NewReader(from), korean.EUCKR.NewDecoder())
	d, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return d, nil
}
