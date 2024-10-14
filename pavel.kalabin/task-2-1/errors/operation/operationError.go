package operationError

import "fmt"

type OperationError struct {}

func (err OperationError) Error() string {
    return fmt.Sprintf("Unknown operation: '<=' or '>=' required")
}
