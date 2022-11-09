package billing

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PaymentMethodsClient is the billing client provides access to billing resources for Azure subscriptions.
type PaymentMethodsClient struct {
	BaseClient
}

// NewPaymentMethodsClient creates an instance of the PaymentMethodsClient client.
func NewPaymentMethodsClient(subscriptionID string) PaymentMethodsClient {
	return NewPaymentMethodsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPaymentMethodsClientWithBaseURI creates an instance of the PaymentMethodsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPaymentMethodsClientWithBaseURI(baseURI string, subscriptionID string) PaymentMethodsClient {
	return PaymentMethodsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByBillingAccountName lists the Payment Methods by billing account Id.
// Parameters:
// billingAccountName - billing Account Id.
func (client PaymentMethodsClient) ListByBillingAccountName(ctx context.Context, billingAccountName string) (result PaymentMethodsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PaymentMethodsClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.pmlr.Response.Response != nil {
				sc = result.pmlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNameNextResults
	req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "ListByBillingAccountName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.pmlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "ListByBillingAccountName", resp, "Failure sending request")
		return
	}

	result.pmlr, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "ListByBillingAccountName", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
func (client PaymentMethodsClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/paymentMethods", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
// http.Response Body if it receives an error.
func (client PaymentMethodsClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client PaymentMethodsClient) ListByBillingAccountNameResponder(resp *http.Response) (result PaymentMethodsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNameNextResults retrieves the next set of results, if any.
func (client PaymentMethodsClient) listByBillingAccountNameNextResults(ctx context.Context, lastResults PaymentMethodsListResult) (result PaymentMethodsListResult, err error) {
	req, err := lastResults.paymentMethodsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "listByBillingAccountNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "listByBillingAccountNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "listByBillingAccountNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client PaymentMethodsClient) ListByBillingAccountNameComplete(ctx context.Context, billingAccountName string) (result PaymentMethodsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PaymentMethodsClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccountName(ctx, billingAccountName)
	return
}

// ListByBillingProfileName lists the Payment Methods by billing profile Id.
// Parameters:
// billingAccountName - billing Account Id.
// billingProfileName - billing Profile Id.
func (client PaymentMethodsClient) ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result PaymentMethodsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PaymentMethodsClient.ListByBillingProfileName")
		defer func() {
			sc := -1
			if result.pmlr.Response.Response != nil {
				sc = result.pmlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNameNextResults
	req, err := client.ListByBillingProfileNamePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "ListByBillingProfileName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileNameSender(req)
	if err != nil {
		result.pmlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "ListByBillingProfileName", resp, "Failure sending request")
		return
	}

	result.pmlr, err = client.ListByBillingProfileNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "ListByBillingProfileName", resp, "Failure responding to request")
	}

	return
}

// ListByBillingProfileNamePreparer prepares the ListByBillingProfileName request.
func (client PaymentMethodsClient) ListByBillingProfileNamePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/paymentMethods", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileNameSender sends the ListByBillingProfileName request. The method will close the
// http.Response Body if it receives an error.
func (client PaymentMethodsClient) ListByBillingProfileNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileNameResponder handles the response to the ListByBillingProfileName request. The method always
// closes the http.Response Body.
func (client PaymentMethodsClient) ListByBillingProfileNameResponder(resp *http.Response) (result PaymentMethodsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingProfileNameNextResults retrieves the next set of results, if any.
func (client PaymentMethodsClient) listByBillingProfileNameNextResults(ctx context.Context, lastResults PaymentMethodsListResult) (result PaymentMethodsListResult, err error) {
	req, err := lastResults.paymentMethodsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "listByBillingProfileNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "listByBillingProfileNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PaymentMethodsClient", "listByBillingProfileNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client PaymentMethodsClient) ListByBillingProfileNameComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result PaymentMethodsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PaymentMethodsClient.ListByBillingProfileName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfileName(ctx, billingAccountName, billingProfileName)
	return
}