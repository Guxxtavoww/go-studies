package main

import (
	"context"
	"fmt"
	"time"
)

type Data struct {
	Id string `json:"id"`
}

// fetch_data simulates fetching a Data record by id from an external
// source (e.g. a database), while respecting the cancellation/timeout
// signal carried by ctx.
//
// It blocks for 3 seconds to simulate a slow operation. If ctx is
// canceled or its deadline is exceeded before those 3 seconds elapse,
// fetch_data returns immediately with a zero-value Data and the error
// from ctx.Err() (either context.Canceled or context.DeadlineExceeded).
//
// Parameters:
//   - ctx: carries the cancellation signal and/or deadline for this call.
//   - id:  the identifier of the record to fetch.
//
// Returns:
//   - Data:  the fetched record (only valid if err is nil).
//   - error: non-nil if the context was canceled or timed out before
//     the operation completed.
func fetch_data(ctx context.Context, id string) (Data, error) {
	select {
	case <-time.After(3 * time.Second):
		// Simulated work finished successfully within the allowed time.
		return Data{Id: id}, nil
	case <-ctx.Done():
		// The context was canceled or its deadline was reached first.
		// ctx.Err() reports which of the two happened.
		return Data{}, ctx.Err()
	}
}

// main creates a context with a 2-second timeout and uses it to call
// fetch_data, which itself takes 3 seconds to "complete". Since the
// timeout (2s) is shorter than the simulated work (3s), fetch_data is
// expected to fail with context.DeadlineExceeded.
func main() {
	// WithTimeout derives a new context from context.Background() that
	// will automatically be canceled after 2 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// cancel must always be called once the context is no longer needed,
	// even if the timeout already fired, to release resources associated
	// with the context (avoids a context leak).
	defer cancel()

	data, err := fetch_data(ctx, "user_id")

	if err != nil {
		// In this example, fetch_data is expected to time out before
		// completing, so this branch will run in practice.
		panic("erro ou canceled")
	}

	fmt.Println(data.Id)
}
