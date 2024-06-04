// *******************************************************************************************
// Copyright © 2019 Xiippy.ai. All rights reserved. Australian patents awarded. PCT patent pending.
//
// NOTES:
//
// - No payment gateway SDK function is consumed directly. Interfaces are defined out of such interactions and then the interface is implemented for payment gateways. Design the interface with the most common members and data structures between different gateways.
// - A proper factory or provider must instantiate an instance of the interface that is interacted with.
// - Any major change made to SDKs should begin with the c sharp SDK with the mindset to keep the high-level syntax, structures and class names the same to minimise porting efforts to other languages. Do not use language specific features that do not exist in other languages. We are not in the business of doing the same thing from scratch multiple times in different forms.
// - Pascal Case for naming conventions should be used for all languages
// - No secret or passwords or keys must exist in the code when checked in
//
// *******************************************************************************************

package XiippySDKBridgeApiClient

import (
	"github.com/Xiippy/Xiippy.POSeCommSDK.Light_GoLang/Models"
	"github.com/Xiippy/Xiippy.POSeCommSDK.Light_GoLang/Utils"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	ApplicationJson = "application/json"
)

type XiippySDKBridgeApiClient struct {
	BridgeBaseUrl   string
	client          *http.Client
	IsTest          bool
	BridgeAPIKey    string
	MerchantID      string
	MerchantGroupID string
}

func NewXiippySDKBridgeApiClient(isTest bool, bridgeAPIKey, bridgeBaseUrl, merchantID, merchantGroupID string) *XiippySDKBridgeApiClient {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	if isTest {
		// Bypass SSL certificate validation in test mode
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second,
		}
	}

	return &XiippySDKBridgeApiClient{
		BridgeBaseUrl:   bridgeBaseUrl,
		client:          client,
		IsTest:          isTest,
		BridgeAPIKey:    bridgeAPIKey,
		MerchantID:      merchantID,
		MerchantGroupID: merchantGroupID,
	}
}

func (x *XiippySDKBridgeApiClient) InitiateXiippyPayment(req *Models.PaymentProcessingRequest) (*Models.PaymentProcessingResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(jsonData)

	httpReq, err := http.NewRequest(http.MethodPost, x.BridgeBaseUrl+InitiateXiippyPaymentPath, reqBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", ApplicationJson)
	httpReq.Header.Set("Accept", ApplicationJson)

	// Add the authentication signature headers
	Utils.AddXiippyV1RequestSignatureToClient(string(jsonData), httpReq, x.BridgeAPIKey)

	resp, err := x.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Response Code: %d Body: %s", resp.StatusCode, reqBody))
	}

	var responseObj Models.PaymentProcessingResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseObj); err != nil {
		return nil, err
	}

	return &responseObj, nil
}
