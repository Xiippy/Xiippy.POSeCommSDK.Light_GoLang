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
package Utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HashHMAC(key, message []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return h.Sum(nil)
}

const XiippyReqSignatureHeader = "XIIPPY-API-SIG-V1"
const XiippyReqMomentHeader = "XIIPPY-MOMENT-V1"

func AddXiippyV1RequestSignatureToClient(content string, requestObj *http.Request, apiKey string) {
	if apiKey != "" {

		key, _ := base64.StdEncoding.DecodeString(apiKey)

		now := time.Now().UnixNano() / int64(time.Millisecond)

		stringToSign := content + "_" + strconv.FormatInt(now, 10)

		sigBytes := HashHMAC(key, []byte(stringToSign))
		signatureHeaderValue := base64.StdEncoding.EncodeToString(sigBytes)

		// Add the moment header

		requestObj.Header.Set(XiippyReqMomentHeader, strconv.FormatInt(now, 10))

		// Add the signature header
		requestObj.Header.Set(XiippyReqSignatureHeader, signatureHeaderValue)
	}
}

func VerifyXiippyV1RequestSignature(content string, moment int64, signature, apiKey string) bool {
	if apiKey != "" {
		key, _ := base64.StdEncoding.DecodeString(apiKey)

		now := time.Now().UnixNano() / int64(time.Millisecond)

		if now-moment > 20000 { // Allow a 20-second window of validity
			return false
		}

		stringToSign := content + "_" + strconv.FormatInt(moment, 10)

		sigBytes := HashHMAC(key, []byte(stringToSign))
		sigBytesPassed, _ := base64.StdEncoding.DecodeString(signature)
		return strings.EqualFold(string(sigBytes), string(sigBytesPassed))
	}
	return false
}
