package database_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	junodb "github.com/forbole/juno/v2/database"
	"github.com/forbole/juno/v2/logging"

	dbcfg "github.com/forbole/juno/v2/database/config"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/disperze/wasmx/database"
	"github.com/disperze/wasmx/types/config"

	_ "github.com/proullon/ramsql/driver"
)

type DbTestSuite struct {
	suite.Suite

	database *database.Db
	testData TestData
}

type TestData struct {
	contract wasmtypes.ContractInfo
}

func (suite *DbTestSuite) SetupTest() {
	// Setup test data
	suite.setupTestData()

	// Build the database
	encodingConfig := config.MakeEncodingConfig()
	databaseConfig := dbcfg.NewDatabaseConfig(
		"wasmx",
		"localhost",
		5433,
		"wasmx",
		"password",
		"",
		"public",
		10,
		10,
	)

	db, err := database.Builder(junodb.NewContext(databaseConfig, &encodingConfig, logging.DefaultLogger()))
	suite.Require().NoError(err)

	desmosDb, ok := (db).(*database.Db)
	suite.Require().True(ok)

	// Delete the public schema
	_, err = desmosDb.Sql.Exec(fmt.Sprintf(`DROP SCHEMA %s CASCADE;`, databaseConfig.GetSchema()))
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = desmosDb.Sql.Exec(fmt.Sprintf(`CREATE SCHEMA %s;`, databaseConfig.GetSchema()))
	suite.Require().NoError(err)

	dirPath := "schema"
	dir, err := ioutil.ReadDir(dirPath)
	for _, fileInfo := range dir {
		if !strings.HasSuffix(fileInfo.Name(), ".sql") {
			continue
		}

		file, err := ioutil.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		suite.Require().NoError(err)

		commentsRegExp := regexp.MustCompile(`/\*.*\*/`)
		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err := desmosDb.Sql.Exec(commentsRegExp.ReplaceAllString(request, ""))
			suite.Require().NoError(err)
		}
	}

	suite.database = desmosDb
}

func (suite *DbTestSuite) setupTestData() {
	sender, _ := sdk.AccAddressFromBech32("cosmos1qpzgtwec63yhxz9hesj8ve0j3ytzhhqaqxrc5d")
	suite.testData = TestData{
		contract: wasmtypes.NewContractInfo(
			2,
			sender,
			sender,
			"First Contract",
			&wasmtypes.AbsoluteTxPosition{
				BlockHeight: 1,
				TxIndex:     1,
			},
		),
	}
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}
