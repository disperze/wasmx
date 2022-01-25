package database_test

import (
	"github.com/disperze/wasmx/types"
)

func (suite *DbTestSuite) TestSaveCode() {
	code := types.NewCode(
		uint64(1),
		"cosmos15c66kjz44zm58xqlcqjwftan4tnaeq7rtmhn4f",
		"ffffff",
		uint64(10000),
		"2022-01-12T00:02:56Z",
		2000,
	)

	// Save the data
	err := suite.database.SaveCode(code)
	suite.Require().NoError(err)

	// Verify code data
	stored, err := suite.database.GetCodeData(code.CodeID)
	suite.Require().NoError(err)
	suite.Require().Nil(stored.Version)
	suite.Require().Nil(stored.CW20)
	suite.Require().Nil(stored.IBC)

	// Update the data
	codeData := types.NewCodeData(code.CodeID, "1.0.0", true, true)
	err = suite.database.SetCodeData(codeData)
	suite.Require().NoError(err)

	// Verify code data
	stored, err = suite.database.GetCodeData(code.CodeID)
	suite.Require().NoError(err)
	suite.Require().NotNil(stored.Version)
	suite.Require().True(*stored.CW20)
	suite.Require().True(*stored.IBC)
}
