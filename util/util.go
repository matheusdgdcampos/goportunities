package util

import "fmt"

func ErroParamIsMissing(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
