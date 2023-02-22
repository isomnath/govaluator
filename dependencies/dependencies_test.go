package dependencies

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DependenciesTestSuite struct {
	suite.Suite
}

func (suite *DependenciesTestSuite) TestInitializeDependencies() {
	suite.NotNil(Initialize())
}

func TestDependenciesTestSuite(t *testing.T) {
	suite.Run(t, new(DependenciesTestSuite))
}
