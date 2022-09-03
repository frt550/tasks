//go:build integration

package test

import (
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/repository/postgres"
	pb "tasks/pkg/api/task"
	testPoolPkg "tasks/test/pool"
	"testing"

	"github.com/stretchr/testify/suite"
	apiPkg "tasks/internal/api/task"
)

type ApiSuite struct {
	suite.Suite
	core     taskPkg.Interface
	api      pb.AdminServer
	_cleanup func()
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(ApiSuite))
}

func (s *ApiSuite) SetupTest() {
	pool, cleanup := testPoolPkg.GetInstance()
	// SetupTest is also called before suite (not sure), causing unexpected behavior
	// so call cleanup of previous setup
	if s._cleanup != nil {
		s._cleanup()
	}
	s.core = taskPkg.New(postgres.New(pool))
	s.api = apiPkg.New(s.core)
	s._cleanup = cleanup
}

func (s *ApiSuite) TearDownTest() {
	s._cleanup()
	s._cleanup = nil
}
