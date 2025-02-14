package test

import (
	"testing"

	helper "github.com/cloudposse/test-helpers/pkg/atmos/component-helper"
)


type ComponentSuite struct {
	helper.TestSuite
}


func (s *ComponentSuite) TestDisabled() {
	const component = "eks/idp-roles/disabled"
	const stack = "default-test"

	s.VendorPull()

	s.RunAtmosCommand("init", component, "-s", stack)
}



func TestRunEksIdpRolesSuite(t *testing.T) {
	suite := new(ComponentSuite)
	helper.Run(t, suite)
}