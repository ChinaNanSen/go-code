package streamanalytics

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
	"net/http"
)

// FunctionsClient is the stream Analytics Client
type FunctionsClient struct {
	BaseClient
}

// NewFunctionsClient creates an instance of the FunctionsClient client.
func NewFunctionsClient(subscriptionID string) FunctionsClient {
	return NewFunctionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFunctionsClientWithBaseURI creates an instance of the FunctionsClient client.
func NewFunctionsClientWithBaseURI(baseURI string, subscriptionID string) FunctionsClient {
	return FunctionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrReplace creates a function or replaces an already existing function under an existing streaming job.
// Parameters:
// function - the definition of the function that will be used to create a new function or replace the existing
// one under the streaming job.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// functionName - the name of the function.
// ifMatch - the ETag of the function. Omit this value to always overwrite the current function. Specify the
// last-seen ETag value to prevent accidentally overwritting concurrent changes.
// ifNoneMatch - set to '*' to allow a new function to be created, but to prevent updating an existing
// function. Other values will result in a 412 Pre-condition Failed response.
func (client FunctionsClient) CreateOrReplace(ctx context.Context, function Function, resourceGroupName string, jobName string, functionName string, ifMatch string, ifNoneMatch string) (result Function, err error) {
	req, err := client.CreateOrReplacePreparer(ctx, function, resourceGroupName, jobName, functionName, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrReplaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrReplaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", resp, "Failure responding to request")
	}

	return
}

// CreateOrReplacePreparer prepares the CreateOrReplace request.
func (client FunctionsClient) CreateOrReplacePreparer(ctx context.Context, function Function, resourceGroupName string, jobName string, functionName string, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithJSON(function),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrReplaceSender sends the CreateOrReplace request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) CreateOrReplaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrReplaceResponder handles the response to the CreateOrReplace request. The method always
// closes the http.Response Body.
func (client FunctionsClient) CreateOrReplaceResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a function from the streaming job.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// functionName - the name of the function.
func (client FunctionsClient) Delete(ctx context.Context, resourceGroupName string, jobName string, functionName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, jobName, functionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FunctionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, jobName string, functionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FunctionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details about the specified function.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// functionName - the name of the function.
func (client FunctionsClient) Get(ctx context.Context, resourceGroupName string, jobName string, functionName string) (result Function, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, jobName, functionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FunctionsClient) GetPreparer(ctx context.Context, resourceGroupName string, jobName string, functionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FunctionsClient) GetResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByStreamingJob lists all of the functions under the specified streaming job.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// selectParameter - the $select OData query parameter. This is a comma-separated list of structural properties
// to include in the response, or ???*??? to include all properties. By default, all properties are returned except
// diagnostics. Currently only accepts '*' as a valid value.
func (client FunctionsClient) ListByStreamingJob(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result FunctionListResultPage, err error) {
	result.fn = client.listByStreamingJobNextResults
	req, err := client.ListByStreamingJobPreparer(ctx, resourceGroupName, jobName, selectParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByStreamingJobSender(req)
	if err != nil {
		result.flr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", resp, "Failure sending request")
		return
	}

	result.flr, err = client.ListByStreamingJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", resp, "Failure responding to request")
	}

	return
}

// ListByStreamingJobPreparer prepares the ListByStreamingJob request.
func (client FunctionsClient) ListByStreamingJobPreparer(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByStreamingJobSender sends the ListByStreamingJob request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) ListByStreamingJobSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByStreamingJobResponder handles the response to the ListByStreamingJob request. The method always
// closes the http.Response Body.
func (client FunctionsClient) ListByStreamingJobResponder(resp *http.Response) (result FunctionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByStreamingJobNextResults retrieves the next set of results, if any.
func (client FunctionsClient) listByStreamingJobNextResults(lastResults FunctionListResult) (result FunctionListResult, err error) {
	req, err := lastResults.functionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "listByStreamingJobNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByStreamingJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "listByStreamingJobNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByStreamingJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "listByStreamingJobNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByStreamingJobComplete enumerates all values, automatically crossing page boundaries as required.
func (client FunctionsClient) ListByStreamingJobComplete(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result FunctionListResultIterator, err error) {
	result.page, err = client.ListByStreamingJob(ctx, resourceGroupName, jobName, selectParameter)
	return
}

// RetrieveDefaultDefinition retrieves the default definition of a function based on the parameters specified.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// functionName - the name of the function.
// functionRetrieveDefaultDefinitionParameters - parameters used to specify the type of function to retrieve
// the default definition for.
func (client FunctionsClient) RetrieveDefaultDefinition(ctx context.Context, resourceGroupName string, jobName string, functionName string, functionRetrieveDefaultDefinitionParameters *BasicFunctionRetrieveDefaultDefinitionParameters) (result Function, err error) {
	req, err := client.RetrieveDefaultDefinitionPreparer(ctx, resourceGroupName, jobName, functionName, functionRetrieveDefaultDefinitionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", nil, "Failure preparing request")
		return
	}

	resp, err := client.RetrieveDefaultDefinitionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", resp, "Failure sending request")
		return
	}

	result, err = client.RetrieveDefaultDefinitionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", resp, "Failure responding to request")
	}

	return
}

// RetrieveDefaultDefinitionPreparer prepares the RetrieveDefaultDefinition request.
func (client FunctionsClient) RetrieveDefaultDefinitionPreparer(ctx context.Context, resourceGroupName string, jobName string, functionName string, functionRetrieveDefaultDefinitionParameters *BasicFunctionRetrieveDefaultDefinitionParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}/RetrieveDefaultDefinition", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if functionRetrieveDefaultDefinitionParameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(functionRetrieveDefaultDefinitionParameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RetrieveDefaultDefinitionSender sends the RetrieveDefaultDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) RetrieveDefaultDefinitionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// RetrieveDefaultDefinitionResponder handles the response to the RetrieveDefaultDefinition request. The method always
// closes the http.Response Body.
func (client FunctionsClient) RetrieveDefaultDefinitionResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Test tests if the information provided for a function is valid. This can range from testing the connection to the
// underlying web service behind the function or making sure the function code provided is syntactically correct.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// functionName - the name of the function.
// function - if the function specified does not already exist, this parameter must contain the full function
// definition intended to be tested. If the function specified already exists, this parameter can be left null
// to test the existing function as is or if specified, the properties specified will overwrite the
// corresponding properties in the existing function (exactly like a PATCH operation) and the resulting
// function will be tested.
func (client FunctionsClient) Test(ctx context.Context, resourceGroupName string, jobName string, functionName string, function *Function) (result FunctionsTestFuture, err error) {
	req, err := client.TestPreparer(ctx, resourceGroupName, jobName, functionName, function)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Test", nil, "Failure preparing request")
		return
	}

	result, err = client.TestSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Test", result.Response(), "Failure sending request")
		return
	}

	return
}

// TestPreparer prepares the Test request.
func (client FunctionsClient) TestPreparer(ctx context.Context, resourceGroupName string, jobName string, functionName string, function *Function) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}/test", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if function != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(function))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TestSender sends the Test request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) TestSender(req *http.Request) (future FunctionsTestFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// TestResponder handles the response to the Test request. The method always
// closes the http.Response Body.
func (client FunctionsClient) TestResponder(resp *http.Response) (result ResourceTestStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing function under an existing streaming job. This can be used to partially update (ie.
// update one or two properties) a function without affecting the rest the job or function definition.
// Parameters:
// function - a function object. The properties specified here will overwrite the corresponding properties in
// the existing function (ie. Those properties will be updated). Any properties that are set to null here will
// mean that the corresponding property in the existing function will remain the same and not change as a
// result of this PATCH operation.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// functionName - the name of the function.
// ifMatch - the ETag of the function. Omit this value to always overwrite the current function. Specify the
// last-seen ETag value to prevent accidentally overwritting concurrent changes.
func (client FunctionsClient) Update(ctx context.Context, function Function, resourceGroupName string, jobName string, functionName string, ifMatch string) (result Function, err error) {
	req, err := client.UpdatePreparer(ctx, function, resourceGroupName, jobName, functionName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client FunctionsClient) UpdatePreparer(ctx context.Context, function Function, resourceGroupName string, jobName string, functionName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithJSON(function),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FunctionsClient) UpdateResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
