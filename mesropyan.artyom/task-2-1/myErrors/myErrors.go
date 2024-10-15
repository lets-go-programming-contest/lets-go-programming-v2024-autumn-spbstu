package myErrors

type DepCountError struct{}

func (e *DepCountError) Error() string {
	return "The number of departments is incorrect"
}

type EmplCountError struct{}

func (e *EmplCountError) Error() string {
	return "The number of employees is incorrect"
}
