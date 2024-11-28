// Package release retrieves latest release version.
package release

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/zctx"
)

// Stable version.
func Stable(ctx context.Context) (string, error) {
	b := backoff.NewConstantBackOff(time.Millisecond * 100)
	res, err := backoff.RetryNotifyWithData(func() (*http.Response, error) {
		return http.Get("https://dl.k8s.io/release/stable.txt")
	}, backoff.WithContext(b, ctx), func(err error, duration time.Duration) {
		zctx.From(ctx)
	})
	if err != nil {
		return "", errors.Wrap(err, "fetch")
	}
	defer func() {
		_ = res.Body.Close()
	}()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "read")
	}

	return strings.TrimSpace(string(data)), nil
}
