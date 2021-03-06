package frontend

import (
	"net/http"

	"github.com/pydio/cells/common/config"
)

type RequestStatus struct {
	Config        *config.Config
	AclParameters *config.Map
	AclActions    *config.Map
	WsScopes      []string

	User     *User
	NoClaims bool
	Lang     string

	Request *http.Request
}
