//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Optional fields for `us_bank_account`.
type TestHelpersTreasuryReceivedDebitInitiatingPaymentMethodDetailsUSBankAccountParams struct {
	// The bank account holder's name.
	AccountHolderName *string `form:"account_holder_name"`
	// The bank account number.
	AccountNumber *string `form:"account_number"`
	// The bank account's routing number.
	RoutingNumber *string `form:"routing_number"`
}

// Initiating payment method details for the object.
type TestHelpersTreasuryReceivedDebitInitiatingPaymentMethodDetailsParams struct {
	// The source type.
	Type *string `form:"type"`
	// Optional fields for `us_bank_account`.
	USBankAccount *TestHelpersTreasuryReceivedDebitInitiatingPaymentMethodDetailsUSBankAccountParams `form:"us_bank_account"`
}

// Optional fields for `ach`.
type TestHelpersTreasuryReceivedDebitNetworkDetailsACHParams struct {
	// Addenda record data associated with this ReceivedDebit.
	Addenda *string `form:"addenda"`
}

// Details about the network used for the ReceivedDebit.
type TestHelpersTreasuryReceivedDebitNetworkDetailsParams struct {
	// Optional fields for `ach`.
	ACH *TestHelpersTreasuryReceivedDebitNetworkDetailsACHParams `form:"ach"`
	// The type of flow that originated the ReceivedDebit.
	Type *string `form:"type"`
}

// Use this endpoint to simulate a test mode ReceivedDebit initiated by a third party. In live mode, you can't directly create ReceivedDebits initiated by third parties.
type TestHelpersTreasuryReceivedDebitParams struct {
	Params `form:"*"`
	// Amount (in cents) to be transferred.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The FinancialAccount to pull funds from.
	FinancialAccount *string `form:"financial_account"`
	// Initiating payment method details for the object.
	InitiatingPaymentMethodDetails *TestHelpersTreasuryReceivedDebitInitiatingPaymentMethodDetailsParams `form:"initiating_payment_method_details"`
	// The rails used for the object.
	Network *string `form:"network"`
	// Details about the network used for the ReceivedDebit.
	NetworkDetails *TestHelpersTreasuryReceivedDebitNetworkDetailsParams `form:"network_details"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryReceivedDebitParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
