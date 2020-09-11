/*
 * Paygate API
 *
 * PayGate is a RESTful API enabling first-party Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transfers to be created without a deep understanding of a full NACHA file specification. First-party transfers initiate at an Originating Depository Financial Institution (ODFI) and are sent off to other Financial Institutions.  A namespace is a value used to isolate models from each other. This can be set to a \"user ID\" from your authentication service or any value your system has to identify.  There are also [admin endpoints](https://moov-io.github.io/paygate/admin/) for back-office operations.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Destination Customer that is receiving a Transfer
type Destination struct {
	// A customerID from the Customers service used as source for this Transfer
	CustomerID string `json:"customerID"`
	// A accountID from the Customers service under the specified Customer used for this Transfer. If the Customer only has one account this value can be left empty.
	AccountID string `json:"accountID"`
}
