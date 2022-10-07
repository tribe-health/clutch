package shortlink

import (
	"testing"

	"github.com/lyft/clutch/backend/mock/service/shortlinkmock"
	"github.com/lyft/clutch/backend/module/moduletest"
	"github.com/lyft/clutch/backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap/zaptest"
)

func TestModule(t *testing.T) {
	service.Registry["clutch.service.shortlink"] = shortlinkmock.New()

	log := zaptest.NewLogger(t)
	scope := tally.NewTestScope("", nil)

	m, err := New(nil, log, scope)
	assert.NoError(t, err)

	r := moduletest.NewRegisterChecker()
	assert.NoError(t, m.Register(r))
	assert.True(t, r.JSONRegistered())
}
