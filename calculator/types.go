package calculator

type ModuleType int8

const (
	ModuleType1 ModuleType = iota + 1
	ModuleType2
	ModuleType3
)

type Modules struct {
	Split   float32  `json:"split"`
	Modules []Module `json:"modules"`
	grade   float32
	credit  int8
}

type Module struct {
	Name   string     `json:"name"`
	Type   ModuleType `json:"type"`
	Coff   int8       `json:"coff"`
	Credit int8       `json:"credit"`
	Grades Grade      `json:"grades"`
	grade  float32
}

type Grade struct {
	Exam float32 `json:"exam"`
	Td   float32 `json:"td"`
	Tp   float32 `json:"tp"`
}
