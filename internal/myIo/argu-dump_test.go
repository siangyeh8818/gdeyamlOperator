package myIo

import (
	"testing"
)

func TestDumpArguments(t *testing.T) {

	DumpArguments("../test/deploy.yml", "../test/environment.yml", "../temp/Test_DumpArguments.log")

}
