/*
 * Paygate Admin API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// IatEntryDetail struct for IatEntryDetail
type IatEntryDetail struct {
	// Entry Detail ID
	ID string `json:"ID,omitempty"`
	// TransactionCode if the receivers account is Credit (deposit) to checking account '22' Prenote for credit to checking account '23' Debit (withdrawal) to checking account '27' Prenote for debit to checking account '28' Credit to savings account '32' Prenote for credit to savings account '33' Debit to savings account '37' Prenote for debit to savings account '38'
	TransactionCode int32 `json:"transactionCode,omitempty"`
	// RDFI's routing number without the last digit.
	RDFIIdentification string `json:"RDFIIdentification,omitempty"`
	// Last digit in RDFI routing number.
	CheckDigit string `json:"checkDigit,omitempty"`
	// Number of Addenda Records
	AddendaRecords float32 `json:"AddendaRecords,omitempty"`
	// Number of cents you are debiting/crediting this account
	Amount int32 `json:"amount,omitempty"`
	// The receiver's bank account number you are crediting/debiting. It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DFIAccountNumber string `json:"DFIAccountNumber,omitempty"`
	// Signifies if the record has been screened against OFAC records
	OFACScreeningIndicator string `json:"OFACScreeningIndicator,omitempty"`
	// Signifies if the record has been screened against OFAC records by a secondary entry
	SecondaryOFACScreeningIndicator string `json:"SecondaryOFACScreeningIndicator,omitempty"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record. A value of \"1\" indicates that one ore more addenda records follow, and \"0\" means no such record is present.
	AddendaRecordIndicator int32 `json:"addendaRecordIndicator,omitempty"`
	// Matches the Entry Detail Trace Number of the entry being returned.
	TraceNumber string    `json:"traceNumber,omitempty"`
	Addenda10   Addenda10 `json:"addenda10,omitempty"`
	Addenda11   Addenda11 `json:"addenda11,omitempty"`
	Addenda12   Addenda12 `json:"addenda12,omitempty"`
	Addenda13   Addenda13 `json:"addenda13,omitempty"`
	Addenda14   Addenda14 `json:"addenda14,omitempty"`
	Addenda15   Addenda15 `json:"addenda15,omitempty"`
	Addenda16   Addenda16 `json:"addenda16,omitempty"`
	Addenda17   Addenda17 `json:"addenda17,omitempty"`
	Addenda18   Addenda18 `json:"addenda18,omitempty"`
	Addenda98   Addenda98 `json:"addenda98,omitempty"`
	Addenda99   Addenda99 `json:"addenda99,omitempty"`
	// Category defines if the entry is a Forward, Return, or NOC
	Category string `json:"category,omitempty"`
}
