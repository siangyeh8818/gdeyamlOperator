package action

import (
	myConfig "github.com/siangyeh8818/gdeyamlOperator/internal/config"
)

type ACTION interface {
	DefineAction(myConfig.BINARYCONFIG)
}
