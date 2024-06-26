package upload

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/openshift-pipelines/tekton-caches/internal/provider/oci"
)

func Upload(ctx context.Context, hash, target, folder string, insecure bool) error {
	u, err := url.Parse(target)
	if err != nil {
		return err
	}
	newTarget := strings.TrimPrefix(target, u.Scheme+"://")
	switch u.Scheme {
	case "oci":
		return oci.Upload(ctx, hash, newTarget, folder, insecure)
	case "s3":
		return fmt.Errorf("s3 schema not (yet) supported: %s", target)
	case "gcs":
		return fmt.Errorf("gcs schema not (yet) supported: %s", target)
	default:
		return fmt.Errorf("unknown schema: %s", target)
	}
}
