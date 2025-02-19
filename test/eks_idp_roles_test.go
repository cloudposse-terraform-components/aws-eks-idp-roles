package test

import (
	"testing"

	helper "github.com/cloudposse/test-helpers/pkg/atmos/component-helper"
	"github.com/cloudposse/test-helpers/pkg/atmos"
)


type ComponentSuite struct {
	helper.TestSuite
}


func (s *ComponentSuite) TestDisabled() {
	const component = "eks/idp-roles/disabled"
	const stack = "default-test"
	opts := s.GetAtmosOptions(component, stack, nil)

	atmos.VendorPull(s)

	atmos.RunAtmosCommand(s, opts, "init")
}

func TestRunEksIdpRolesSuite(t *testing.T) {
	suite := new(ComponentSuite)
	helper.Run(t, suite)
}