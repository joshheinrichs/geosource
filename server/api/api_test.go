package api

import (
	"log"
	"os"
	"testing"

	"github.com/joshheinrichs/geosource/server/config"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	testConfig, err := config.ReadFile("../config_test.gcfg")
	testConfig.Website.Directory = "../" + testConfig.Website.Directory
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = transactions.Init(testConfig)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	Init(testConfig)
	os.Exit(m.Run())

}
