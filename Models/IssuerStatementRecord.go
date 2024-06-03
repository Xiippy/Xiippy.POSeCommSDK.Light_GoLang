/*******************************************************************************************
Copyright © 2019 Xiippy.ai. All rights reserved. Australian patents awarded. PCT patent pending.

NOTES:

- No payment gateway SDK function is consumed directly. Interfaces are defined out of such interactions and then the interface is implemented for payment gateways. Design the interface with the most common members and data structures between different gateways.
A proper factory or provider must instantiate an instance of the interface that is interacted with.
- Any major change made to SDKs should begin with the c# SDK with the mindset to keep the high-level syntax, structures and class names the same to minimise porting efforts to other languages. Do not use language specific features that don't exist in other languages. We are not in the business of doing the same thing from scratch multiple times in different forms.
- Pascal Case for naming conventions should be used for all languages
- No secret or passwords or keys must exist in the code when checked in

*******************************************************************************************/

package models

type IssuerStatementRecord struct {
	RandomStatementID               string                       `json:"RandomStatementID"`
	StatementItems                  []StatementItem              `json:"StatementItems"`
	ElectronicPayments              []ElectronicPaymentPersisted `json:"ElectronicPayments"`
	CashPayments                    []CashPaymentPersisted       `json:"CashPayments"`
	TotalBillVariations             []TotalBillVariation         `json:"TotalBillVariations"`
	ShiftID                         string                       `json:"ShiftID"`
	StatementTimeStamp              string                       `json:"StatementTimeStamp"`
	IssuersPrivateMetadata          map[string]string            `json:"IssuersPrivateMetadata"`
	ShortStatementID                string                       `json:"ShortStatementID"`
	StatementCreationDate           string                       `json:"StatementCreationDate"`
	StatementText                   string                       `json:"StatementText"`
	Description                     string                       `json:"Description"`
	PersonResponsible               string                       `json:"PersonResponsible"`
	IssuerLogoUrl                   string                       `json:"IssuerLogoUrl"`
	DetailsInHeader                 string                       `json:"DetailsInHeader"`
	TotalAmount                     float32                      `json:"TotalAmount"`
	TotalTaxAmount                  float32                      `json:"TotalTaxAmount"`
	DetailsInBodyBeforeItems        string                       `json:"DetailsInBodyBeforeItems"`
	DetailsInBodyAfterItems         string                       `json:"DetailsInBodyAfterItems"`
	DetailsInFooter                 string                       `json:"DetailsInFooter"`
	StatementIdentifier             string                       `json:"StatementIdentifier"`
	StatementIdentifier2            string                       `json:"StatementIdentifier2"`
	StatementIdentifier3            string                       `json:"StatementIdentifier3"`
	ExtraInfo1                      string                       `json:"ExtraInfo1"`
	ExtraInfo2                      string                       `json:"ExtraInfo2"`
	ExtraInfo3                      string                       `json:"ExtraInfo3"`
	ExtraInfo4                      string                       `json:"ExtraInfo4"`
	ExtraInfo5                      string                       `json:"ExtraInfo5"`
	ExtraInfo6                      string                       `json:"ExtraInfo6"`
	ExtraInfo7                      string                       `json:"ExtraInfo7"`
	ExtraInfo8                      string                       `json:"ExtraInfo8"`
	ExtraInfo9                      string                       `json:"ExtraInfo9"`
	ExtraInfo10                     string                       `json:"ExtraInfo10"`
	AttachmentPageTop               []byte                       `json:"AttachmentPageTop"`
	AttachmentPageTopMimeType       string                       `json:"AttachmentPageTopMimeType"`
	AttachmentItemsTop              []byte                       `json:"AttachmentItemsTop"`
	AttachmentItemsTopMimeType      string                       `json:"AttachmentItemsTopMimeType"`
	AttachmentItemsBottom           []byte                       `json:"AttachmentItemsBottom"`
	AttachmentItemsBottomMimeType   string                       `json:"AttachmentItemsBottomMimeType"`
	AttachmentPageBottom            []byte                       `json:"AttachmentPageBottom"`
	AttachmentPageBottomMimeType    string                       `json:"AttachmentPageBottomMimeType"`
	TotalLoyaltyPoints              float32                      `json:"TotalLoyaltyPoints"`
	PaymentProcessingResult         uint                         `json:"PaymentProcessingResult"`
	PaymentProcessingMsg            string                       `json:"PaymentProcessingMsg"`
	UniqueStatementID               string                       `json:"UniqueStatementID"`
	PaymentNameAddressRequestObject PaymentNameAddressRequest    `json:"PaymentNameAddressRequestObject"`
	RetailerGroupID                 string                       `json:"RetailerGroupID"`
	MetaDataExtras                  map[string]string            `json:"MetaDataExtras"`
}
