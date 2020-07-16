package hellomock

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_hellomock "hellomock/mock"
)

func TestCompany_Meeting(t *testing.T) {
	person:=NewPersobn("王你妹")
	company:=NewCompany(person)
	t.Log(company.Meeting("王尼玛"))
}
func TestCompany_Meeting2(t *testing.T) {
	ctl:=gomock.NewController(t)
	mock_talker:=mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("string")
	company:=NewCompany(mock_talker)
	t.Log(company.Meeting("王玛"))
}