package azuredataapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/azuredata/mgmt/2017-03-01-preview/azuredata"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result azuredata.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result azuredata.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*azuredata.OperationsClient)(nil)

// SQLServerRegistrationsClientAPI contains the set of methods on the SQLServerRegistrationsClient type.
type SQLServerRegistrationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, parameters azuredata.SQLServerRegistration) (result azuredata.SQLServerRegistration, err error)
	Delete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string) (result azuredata.SQLServerRegistration, err error)
	List(ctx context.Context) (result azuredata.SQLServerRegistrationListResultPage, err error)
	ListComplete(ctx context.Context) (result azuredata.SQLServerRegistrationListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result azuredata.SQLServerRegistrationListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result azuredata.SQLServerRegistrationListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, parameters azuredata.SQLServerRegistrationUpdate) (result azuredata.SQLServerRegistration, err error)
}

var _ SQLServerRegistrationsClientAPI = (*azuredata.SQLServerRegistrationsClient)(nil)

// SQLServersClientAPI contains the set of methods on the SQLServersClient type.
type SQLServersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string, parameters azuredata.SQLServer) (result azuredata.SQLServer, err error)
	Delete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string, expand string) (result azuredata.SQLServer, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, expand string) (result azuredata.SQLServerListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, expand string) (result azuredata.SQLServerListResultIterator, err error)
}

var _ SQLServersClientAPI = (*azuredata.SQLServersClient)(nil)