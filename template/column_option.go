package template

// Model used as a variable because it cannot load template file after packed, params still can pass file
var ColumnOptionTemplate = map[string]string{"common": `package {{Package}}

import (
    "database/sql/driver"
)

type {{OptionStructName}} struct {

}

func (c {{OptionStructName}}) Value() (driver.Value, error) {
	return c, nil
}

func (c *{{OptionStructName}}) Scan(value interface{}) error {
	return nil
}

`,
	"json": `package {{Package}}

import (
    "database/sql/driver"
	"encoding/json"
	"fmt"
)

type {{OptionStructName}} struct {

}

func (c {{OptionStructName}}) Value() (driver.Value, error) {
	if &c == nil {
		return "", nil
	}

	return json.Marshal(c)
}

func (c *{{OptionStructName}}) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}
	if len(b) == 0 {
		return nil
	}

	return json.Unmarshal(b, &c)
}
`,
}