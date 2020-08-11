package routing

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spidernest-go/logger"
)

var errNotWebP = errors.New("Image did not contain a WebP")

func decodeJacket(b64jacket string) (*os.File, error) {
	// decode jacket to a byte slice
	jacket, err := base64.StdEncoding.DecodeString(b64jacket)
	if err != nil {
		logger.Warn().
			Err(err).
			Msg("Jacket was not a valid base64 string.")

		return nil, err
	}

	// ensure the jacket is an actual webp
	ftype := http.DetectContentType(jacket)
	logger.Info().Msg(ftype)
	if ftype != "image/webp" {
		logger.Warn().
			Msg("Jacket was not a image/webp MIME type.")
		return nil, errNotWebP
	}

	// create a temporary file and write to it
	f, err := ioutil.TempFile(filepath.Join(os.TempDir(), "orchestrafm"),
		"img-*.webp")

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Temporary file couldn't be created.")

		return nil, err
	}

	if _, err := f.Write(jacket); err != nil {
		logger.Error().
			Err(err).
			Msg("Temporary file couldn't be written to.")

		return nil, err
	}
	if err := f.Sync(); err != nil {
		logger.Error().
			Err(err).
			Msg("Temporary file contents were not commited to storage (fsync).")

		return nil, err
	}

	return f, nil
}
