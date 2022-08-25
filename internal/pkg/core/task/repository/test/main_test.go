package test

import (
	repositoryPkg "tasks/internal/pkg/core/task/repository"
	"tasks/internal/pkg/core/task/repository/postgres"
	testPoolPkg "tasks/test/pool"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	repository repositoryPkg.Interface
	_cleanup   func()
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) SetupTest() {
	pool, cleanup := testPoolPkg.GetInstance()
	// SetupTest is also called before suite (not sure), causing unexpected behavior
	// so call cleanup of previous setup
	if s._cleanup != nil {
		s._cleanup()
	}
	s.repository = postgres.New(pool)
	s._cleanup = cleanup
}

func (s *RepositorySuite) TearDownTest() {
	s._cleanup()
	s._cleanup = nil
}
