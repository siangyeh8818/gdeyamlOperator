package gdeyamloperator

type OPERATORPLAYBOOK struct {
	STAGE      []STAGE     `yaml:"stage"`
}

type STAGE struct {
	Name    string `yaml:"name"`
	Action ACTION `yaml:"action"`
}

type ACTION struct {
	BINARYCONFIG BINARYCONFIG  `yaml:"argv"`
}