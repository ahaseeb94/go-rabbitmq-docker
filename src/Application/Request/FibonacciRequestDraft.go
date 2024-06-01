package Request

type FibonacciRequestDraft struct {
	Number int `form:"number" binding:"required"`
}
