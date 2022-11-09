package jobapi

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
	"github.com/Azure/azure-sdk-for-go/services/datalake/analytics/2016-11-01/job"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
)

// PipelineClientAPI contains the set of methods on the PipelineClient type.
type PipelineClientAPI interface {
	Get(ctx context.Context, accountName string, pipelineIdentity uuid.UUID, startDateTime *date.Time, endDateTime *date.Time) (result job.PipelineInformation, err error)
	List(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (result job.PipelineInformationListResultPage, err error)
	ListComplete(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (result job.PipelineInformationListResultIterator, err error)
}

var _ PipelineClientAPI = (*job.PipelineClient)(nil)

// RecurrenceClientAPI contains the set of methods on the RecurrenceClient type.
type RecurrenceClientAPI interface {
	Get(ctx context.Context, accountName string, recurrenceIdentity uuid.UUID, startDateTime *date.Time, endDateTime *date.Time) (result job.RecurrenceInformation, err error)
	List(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (result job.RecurrenceInformationListResultPage, err error)
	ListComplete(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (result job.RecurrenceInformationListResultIterator, err error)
}

var _ RecurrenceClientAPI = (*job.RecurrenceClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Build(ctx context.Context, accountName string, parameters job.BuildJobParameters) (result job.Information, err error)
	Cancel(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result autorest.Response, err error)
	Create(ctx context.Context, accountName string, jobIdentity uuid.UUID, parameters job.CreateJobParameters) (result job.Information, err error)
	Get(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result job.Information, err error)
	GetDebugDataPath(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result job.DataPath, err error)
	GetStatistics(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result job.Statistics, err error)
	List(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result job.InfoListResultPage, err error)
	ListComplete(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result job.InfoListResultIterator, err error)
}

var _ ClientAPI = (*job.Client)(nil)