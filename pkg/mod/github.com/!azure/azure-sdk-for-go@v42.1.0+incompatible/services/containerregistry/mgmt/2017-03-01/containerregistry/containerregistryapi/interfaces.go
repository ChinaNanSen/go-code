package containerregistryapi

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
	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2017-03-01/containerregistry"
	"github.com/Azure/go-autorest/autorest"
)

// RegistriesClientAPI contains the set of methods on the RegistriesClient type.
type RegistriesClientAPI interface {
	CheckNameAvailability(ctx context.Context, registryNameCheckRequest containerregistry.RegistryNameCheckRequest) (result containerregistry.RegistryNameStatus, err error)
	Create(ctx context.Context, resourceGroupName string, registryName string, registryCreateParameters containerregistry.RegistryCreateParameters) (result containerregistry.RegistriesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.Registry, err error)
	List(ctx context.Context) (result containerregistry.RegistryListResultPage, err error)
	ListComplete(ctx context.Context) (result containerregistry.RegistryListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerregistry.RegistryListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerregistry.RegistryListResultIterator, err error)
	ListCredentials(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistryListCredentialsResult, err error)
	RegenerateCredential(ctx context.Context, resourceGroupName string, registryName string, regenerateCredentialParameters containerregistry.RegenerateCredentialParameters) (result containerregistry.RegistryListCredentialsResult, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, registryUpdateParameters containerregistry.RegistryUpdateParameters) (result containerregistry.Registry, err error)
}

var _ RegistriesClientAPI = (*containerregistry.RegistriesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result containerregistry.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result containerregistry.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*containerregistry.OperationsClient)(nil)