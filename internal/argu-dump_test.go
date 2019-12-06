package gdeyamloperator

import (
	"testing"

	. "github.com/siangyeh8818/gdeyamlOperator/internal"
)

func TestDumpArguments(t *testing.T) {

	DumpArguments("../test/deploy.yml", "../test/environment.yml", "../temp/Test_DumpArguments.log")

}
