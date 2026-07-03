package main

import (
	"context"
	"fmt"
	"time"

	workflow "github.com/IvalisEXE/workflow-engine-go/workflow/step"
)

func main() {
	engine := workflow.NewEngine()

	flow := workflow.NewWorkflow("Debit Money")

	flow.
		Do(Debit{}).
		Retry(1).
		Timeout(5 * time.Second).
		Compensate(
			RefundDebit{},
		)

	err := engine.Run(
		context.Background(),
		flow,
	)
	if err != nil {
		fmt.Printf("Workflow failed: %v\n", err)
	}
}

type CheckBalance struct{}

func (CheckBalance) Name() string {
	return "Check Balance"
}

func (CheckBalance) Execute(
	ctx context.Context,
) error {

	fmt.Println("Checking balance...")

	return nil
}

type Debit struct{}

func (Debit) Name() string {
	return "Debit"
}

func (Debit) Execute(
	ctx context.Context,
) error {

	fmt.Println("Debit success")

	return nil
}

type RefundDebit struct{}

func (RefundDebit) Name() string {
	return "Refund Debit"
}

func (RefundDebit) Execute(
	ctx context.Context,
) error {

	fmt.Println("Refund Debit success")

	return nil
}

type Credit struct{}

func (Credit) Name() string {
	return "Credit"
}

func (Credit) Execute(
	ctx context.Context,
) error {

	fmt.Println("Credit success")

	return nil
}

type SendNotification struct{}

func (SendNotification) Name() string {
	return "Send Notification"
}

func (SendNotification) Execute(
	ctx context.Context,
) error {

	fmt.Println("Send Notification")

	return nil
}
