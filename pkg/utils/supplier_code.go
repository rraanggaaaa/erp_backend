package utils

import "fmt"

func GenerateSupplierCode(lastNumber int) string {

	return fmt.Sprintf("61%06d", lastNumber+1)

}
