/*******************************************************************************************
Copyright © 2019 Xiippy.ai. All rights reserved. Australian patents awarded. PCT patent pending.

NOTES:

- No payment gateway SDK function is consumed directly. Interfaces are defined out of such interactions and then the interface is implemented for payment gateways. Design the interface with the most common members and data structures between different gateways.
A proper factory or provider must instantiate an instance of the interface that is interacted with.
- Any major change made to SDKs should begin with the c# SDK with the mindset to keep the high-level syntax, structures and class names the same to minimise porting efforts to other languages. Do not use language specific features that don't exist in other languages. We are not in the business of doing the same thing from scratch multiple times in different forms.
- Pascal Case for naming conventions should be used for all languages
- No secret or passwords or keys must exist in the code when checked in

*******************************************************************************************/
//using SQLite;
package models

import "time"

type StatementItem struct {
	StatementItemID                  int       `json:"StatementItemID"`
	RandomStatementID                string    `json:"RandomStatementID"`
	Description                      string    `json:"Description"`
	Identifier                       string    `json:"Identifier"`
	Url                              string    `json:"Url"`
	Quantity                         float64   `json:"Quantity"`
	UnitPrice                        float64   `json:"UnitPrice"`
	TotalPrice                       float64   `json:"TotalPrice"`
	Tax                              float64   `json:"Tax"`
	ExtraInfo1                       string    `json:"ExtraInfo1"`
	ExtraInfo2                       string    `json:"ExtraInfo2"`
	ExtraInfo3                       string    `json:"ExtraInfo3"`
	ExtraInfo4                       string    `json:"ExtraInfo4"`
	ExtraInfo5                       string    `json:"ExtraInfo5"`
	ExtraInfo6                       string    `json:"ExtraInfo6"`
	UnitTitle                        string    `json:"UnitTitle"`
	UnitLoyaltyPoint                 float64   `json:"UnitLoyaltyPoint"`
	LoyaltyPoint                     float64   `json:"LoyaltyPoint"`
	ItemClassification               string    `json:"ItemClsassification"`
	ItemCategoryID                   string    `json:"ItemCategoryID"`
	ItemCategoryTitle                string    `json:"ItemCategoryTitle"`
	AddedMoment                      time.Time `json:"AddedMoment"`
	WarrantyExpiryMomentISO8601      string    `json:"WarrantyExpiryMomentISO8601"`
	LoyaltyPointsExpiryMomentISO8601 string    `json:"LoyaltyPointsExpiryMomentISO8601"`
}
