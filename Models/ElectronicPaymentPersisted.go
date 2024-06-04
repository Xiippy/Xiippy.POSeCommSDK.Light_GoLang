/*
******************************************************************************************
Copyright © 2019 Xiippy.ai. All rights reserved. Australian patents awarded. PCT patent pending.

NOTES:

- No payment gateway SDK function is consumed directly. Interfaces are defined out of such interactions and then the interface is implemented for payment gateways. Design the interface with the most common members and data structures between different gateways.
A proper factory or provider must instantiate an instance of the interface that is interacted with.
- Any major change made to SDKs should begin with the c# SDK with the mindset to keep the high-level syntax, structures and class names the same to minimise porting efforts to other languages. Do not use language specific features that don't exist in other languages. We are not in the business of doing the same thing from scratch multiple times in different forms.
- Pascal Case for naming conventions should be used for all languages
- No secret or passwords or keys must exist in the code when checked in

******************************************************************************************
*/
package Models

type ElectronicPaymentPersisted struct {
	ElectronicPaymentID            int64             `json:"ElectronicPaymentID"`
	RandomStatementID              string            `json:"RandomStatementID"`
	Bank                           string            `json:"Bank"`
	MerchantAccountOwnerDetail     string            `json:"MerchantAccountOwnerDetail"`
	Terminal                       string            `json:"Terminal"`
	Reference                      string            `json:"Reference"`
	CardNO                         string            `json:"CardNO"`
	AccountType                    string            `json:"AccountType"`
	CardExpiry                     string            `json:"CardExpiry"`
	Aid                            string            `json:"Aid"`
	Atc                            string            `json:"Atc"`
	Tvr                            string            `json:"Tvr"`
	Csn                            string            `json:"Csn"`
	AuthNo                         string            `json:"AuthNo"`
	PosRefNo                       string            `json:"PosRefNo"`
	MAccountNumber                 string            `json:"MAccountNumber"`
	Mrrn                           string            `json:"Mrrn"`
	Mauth                          string            `json:"Mauth"`
	PaymentType                    string            `json:"PaymentType"`
	MLocationCode                  string            `json:"MLocationCode"`
	MAccountType                   string            `json:"MAccountType"`
	Apsn                           string            `json:"Apsn"`
	Arqc                           string            `json:"Arqc"`
	CurrencyCode                   string            `json:"CurrencyCode"`
	ExtraInfo1                     string            `json:"ExtraInfo1"`
	ExtraInfo2                     string            `json:"ExtraInfo2"`
	ExtraInfo3                     string            `json:"ExtraInfo3"`
	ExtraInfo4                     string            `json:"ExtraInfo4"`
	ExtraInfo5                     string            `json:"ExtraInfo5"`
	ExtraInfo6                     string            `json:"ExtraInfo6"`
	ExtraInfo7                     string            `json:"ExtraInfo7"`
	ExtraInfo8                     string            `json:"ExtraInfo8"`
	ExtraInfo9                     string            `json:"ExtraInfo9"`
	ExtraInfo10                    string            `json:"ExtraInfo10"`
	Purchase                       float32           `json:"Purchase"`
	Total                          float32           `json:"Total"`
	TransactionType                string            `json:"TransactionType"`
	StatusId                       string            `json:"StatusId"`
	TxnStatusId                    string            `json:"TxnStatusId"`
	Complete                       string            `json:"Complete"`
	StatementText                  string            `json:"StatementText"`
	ApprovedFlag                   bool              `json:"ApprovedFlag"`
	ExpectedSettlementDate         string            `json:"ExpectedSettlementDate"`
	ExpectedSettlementDateTimeZone string            `json:"ExpectedSettlementDateTimeZone"`
	DateOfTransaction              string            `json:"DateOfTransaction"`
	DateOfTransactionTimeZone      string            `json:"DateOfTransactionTimeZone"`
	Stan                           string            `json:"Stan"`
	DpsBillingId                   string            `json:"DpsBillingId"`
	ResponseCode                   string            `json:"ResponseCode"`
	ResponseText                   string            `json:"ResponseText"`
	AmtSurcgarge                   float32           `json:"AmtSurcgarge"`
	AmtTip                         float32           `json:"AmtTip"`
	AmtCashOut                     float32           `json:"AmtCashOut"`
	CardType                       string            `json:"CardType"`
	MetaDataExtras                 map[string]string `json:"MetaDataExtras"`
	Tsi                            string            `json:"Tsi"`
	DedicatedFileName              string            `json:"DedicatedFileName"`
	Cvm                            string            `json:"Cvm"`
	AuthorizationCode              string            `json:"AuthorizationCode"`
	ApplicationPreferredName       string            `json:"ApplicationPreferredName"`
}
