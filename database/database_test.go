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

	"github.com/stretchr/testify/suite"

	"github.com/disperze/wasmx/database"
	"github.com/disperze/wasmx/types/config"

	_ "github.com/proullon/ramsql/driver"
)

type DbTestSuite struct {
	suite.Suite

	database *database.Db
}

func (suite *DbTestSuite) SetupTest() {
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

	wasmDb, ok := (db).(*database.Db)
	suite.Require().True(ok)

	// Delete the public schema
	_, err = wasmDb.Sql.Exec(fmt.Sprintf(`DROP SCHEMA %s CASCADE;`, databaseConfig.Schema))
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = wasmDb.Sql.Exec(fmt.Sprintf(`CREATE SCHEMA %s;`, databaseConfig.Schema))
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
			_, err := wasmDb.Sql.Exec(commentsRegExp.ReplaceAllString(request, ""))
			suite.Require().NoError(err)
		}
	}

	suite.database = wasmDb
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}
