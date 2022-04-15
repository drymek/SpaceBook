package service

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BookingSuite struct {
	suite.Suite
}

func TestDayDateSuite(t *testing.T) {
	s := new(BookingSuite)

	suite.Run(t, s)
}

func (s *BookingSuite) TestXXX() {

}
