package httpsupport

import (
	"testing"

	"net/http"

	"github.com/fabric8-services/fabric8-common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAbsoluteURLOK(t *testing.T) {
	resource.Require(t, resource.UnitTest)
	t.Parallel()

	req := &http.Request{Host: "api.service.domain.org"}
	// HTTP
	urlStr := AbsoluteURL(req, "/testpath")
	assert.Equal(t, "http://api.service.domain.org/testpath", urlStr)

	// HTTPS
	r, err := http.NewRequest("", "https://api.service.domain.org", nil)
	require.NoError(t, err)
	urlStr = AbsoluteURL(r, "/testpath2")
	assert.Equal(t, "https://api.service.domain.org/testpath2", urlStr)
}

func TestAbsoluteURLOKWithProxyForward(t *testing.T) {
	resource.Require(t, resource.UnitTest)
	t.Parallel()

	// HTTPS
	r, err := http.NewRequest("", "http://api.service.domain.org", nil)
	require.NoError(t, err)
	r.Header.Set("X-Forwarded-Proto", "https")
	urlStr := AbsoluteURL(r, "/testpath2")
	assert.Equal(t, "https://api.service.domain.org/testpath2", urlStr)
}

func TestReplaceDomainPrefixOK(t *testing.T) {
	resource.Require(t, resource.UnitTest)
	t.Parallel()

	host, err := ReplaceDomainPrefix("api.service.domain.org", "sso")
	require.NoError(t, err)
	assert.Equal(t, "sso.service.domain.org", host)
}

func TestReplaceDomainPrefixInTooShortHostFails(t *testing.T) {
	resource.Require(t, resource.UnitTest)
	t.Parallel()

	_, err := ReplaceDomainPrefix("org", "sso")
	assert.NotNil(t, err)
}