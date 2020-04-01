/*
 * Paygate Admin API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// FileHeader struct for FileHeader
type FileHeader struct {
	// contains the Routing Number of the ACH Operator or sending point that is sending the file.
	ImmediateOrigin string `json:"immediateOrigin"`
	// The name of the ACH operator or sending point that is sending the file.
	ImmediateOriginName string `json:"immediateOriginName"`
	// contains the Routing Number of the ACH Operator or receiving point to which the file is being sent
	ImmediateDestination string `json:"immediateDestination"`
	// The name of the ACH or receiving point for which that file is destined.
	ImmediateDestinationName string `json:"immediateDestinationName"`
	// The File Creation Date is the date when the file was prepared by an ODFI. (Format HHmm - H=Hour, m=Minute)
	FileCreationTime string `json:"fileCreationTime,omitempty"`
	// The File Creation Time is the time when the file was prepared by an ODFI. (Format YYMMDD - Y=Year, M=Month, D=Day)
	FileCreationDate string `json:"fileCreationDate,omitempty"`
	// Incremented value for each file for RDFI's.
	FileIDModifier string `json:"fileIDModifier,omitempty"`
}
