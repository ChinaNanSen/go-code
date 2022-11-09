// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180321

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-21"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAgentPayDealsRequest() (request *AgentPayDealsRequest) {
    request = &AgentPayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "AgentPayDeals")
    return
}

func NewAgentPayDealsResponse() (response *AgentPayDealsResponse) {
    response = &AgentPayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 代理商支付订单接口，支持自付/代付
func (c *Client) AgentPayDeals(request *AgentPayDealsRequest) (response *AgentPayDealsResponse, err error) {
    if request == nil {
        request = NewAgentPayDealsRequest()
    }
    response = NewAgentPayDealsResponse()
    err = c.Send(request, response)
    return
}

func NewAgentTransferMoneyRequest() (request *AgentTransferMoneyRequest) {
    request = &AgentTransferMoneyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "AgentTransferMoney")
    return
}

func NewAgentTransferMoneyResponse() (response *AgentTransferMoneyResponse) {
    response = &AgentTransferMoneyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为合作伙伴提供转账给客户能力。仅支持合作伙伴为自己名下客户转账。
func (c *Client) AgentTransferMoney(request *AgentTransferMoneyRequest) (response *AgentTransferMoneyResponse, err error) {
    if request == nil {
        request = NewAgentTransferMoneyRequest()
    }
    response = NewAgentTransferMoneyResponse()
    err = c.Send(request, response)
    return
}

func NewAuditApplyClientRequest() (request *AuditApplyClientRequest) {
    request = &AuditApplyClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "AuditApplyClient")
    return
}

func NewAuditApplyClientResponse() (response *AuditApplyClientResponse) {
    response = &AuditApplyClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 代理商可以审核其名下申请中代客
func (c *Client) AuditApplyClient(request *AuditApplyClientRequest) (response *AuditApplyClientResponse, err error) {
    if request == nil {
        request = NewAuditApplyClientRequest()
    }
    response = NewAuditApplyClientResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePayRelationForClientRequest() (request *CreatePayRelationForClientRequest) {
    request = &CreatePayRelationForClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "CreatePayRelationForClient")
    return
}

func NewCreatePayRelationForClientResponse() (response *CreatePayRelationForClientResponse) {
    response = &CreatePayRelationForClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 合作伙伴为客户创建强代付关系
func (c *Client) CreatePayRelationForClient(request *CreatePayRelationForClientRequest) (response *CreatePayRelationForClientResponse, err error) {
    if request == nil {
        request = NewCreatePayRelationForClientRequest()
    }
    response = NewCreatePayRelationForClientResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentAuditedClientsRequest() (request *DescribeAgentAuditedClientsRequest) {
    request = &DescribeAgentAuditedClientsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentAuditedClients")
    return
}

func NewDescribeAgentAuditedClientsResponse() (response *DescribeAgentAuditedClientsResponse) {
    response = &DescribeAgentAuditedClientsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询已审核客户列表
func (c *Client) DescribeAgentAuditedClients(request *DescribeAgentAuditedClientsRequest) (response *DescribeAgentAuditedClientsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentAuditedClientsRequest()
    }
    response = NewDescribeAgentAuditedClientsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentBillsRequest() (request *DescribeAgentBillsRequest) {
    request = &DescribeAgentBillsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentBills")
    return
}

func NewDescribeAgentBillsResponse() (response *DescribeAgentBillsResponse) {
    response = &DescribeAgentBillsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 代理商可查询自己及名下代客所有业务明细
func (c *Client) DescribeAgentBills(request *DescribeAgentBillsRequest) (response *DescribeAgentBillsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentBillsRequest()
    }
    response = NewDescribeAgentBillsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentClientsRequest() (request *DescribeAgentClientsRequest) {
    request = &DescribeAgentClientsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentClients")
    return
}

func NewDescribeAgentClientsResponse() (response *DescribeAgentClientsResponse) {
    response = &DescribeAgentClientsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 代理商可查询自己名下待审核客户列表
func (c *Client) DescribeAgentClients(request *DescribeAgentClientsRequest) (response *DescribeAgentClientsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentClientsRequest()
    }
    response = NewDescribeAgentClientsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentDealsCacheRequest() (request *DescribeAgentDealsCacheRequest) {
    request = &DescribeAgentDealsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentDealsCache")
    return
}

func NewDescribeAgentDealsCacheResponse() (response *DescribeAgentDealsCacheResponse) {
    response = &DescribeAgentDealsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 供超大型代理商（代客数量>=3000 ）拉取缓存的全量客户订单。
func (c *Client) DescribeAgentDealsCache(request *DescribeAgentDealsCacheRequest) (response *DescribeAgentDealsCacheResponse, err error) {
    if request == nil {
        request = NewDescribeAgentDealsCacheRequest()
    }
    response = NewDescribeAgentDealsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentPayDealsRequest() (request *DescribeAgentPayDealsRequest) {
    request = &DescribeAgentPayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "DescribeAgentPayDeals")
    return
}

func NewDescribeAgentPayDealsResponse() (response *DescribeAgentPayDealsResponse) {
    response = &DescribeAgentPayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 可以查询代理商代付的所有订单
func (c *Client) DescribeAgentPayDeals(request *DescribeAgentPayDealsRequest) (response *DescribeAgentPayDealsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentPayDealsRequest()
    }
    response = NewDescribeAgentPayDealsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientBalanceRequest() (request *DescribeClientBalanceRequest) {
    request = &DescribeClientBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "DescribeClientBalance")
    return
}

func NewDescribeClientBalanceResponse() (response *DescribeClientBalanceResponse) {
    response = &DescribeClientBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为合作伙伴提供查询客户余额能力。调用者必须是合作伙伴，只能查询自己名下客户余额
func (c *Client) DescribeClientBalance(request *DescribeClientBalanceRequest) (response *DescribeClientBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeClientBalanceRequest()
    }
    response = NewDescribeClientBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRebateInfosRequest() (request *DescribeRebateInfosRequest) {
    request = &DescribeRebateInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "DescribeRebateInfos")
    return
}

func NewDescribeRebateInfosResponse() (response *DescribeRebateInfosResponse) {
    response = &DescribeRebateInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 代理商可查询自己名下全部返佣信息
func (c *Client) DescribeRebateInfos(request *DescribeRebateInfosRequest) (response *DescribeRebateInfosResponse, err error) {
    if request == nil {
        request = NewDescribeRebateInfosRequest()
    }
    response = NewDescribeRebateInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSalesmansRequest() (request *DescribeSalesmansRequest) {
    request = &DescribeSalesmansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "DescribeSalesmans")
    return
}

func NewDescribeSalesmansResponse() (response *DescribeSalesmansResponse) {
    response = &DescribeSalesmansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 代理商查询名下业务员列表信息
func (c *Client) DescribeSalesmans(request *DescribeSalesmansRequest) (response *DescribeSalesmansResponse, err error) {
    if request == nil {
        request = NewDescribeSalesmansRequest()
    }
    response = NewDescribeSalesmansResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClientRemarkRequest() (request *ModifyClientRemarkRequest) {
    request = &ModifyClientRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "ModifyClientRemark")
    return
}

func NewModifyClientRemarkResponse() (response *ModifyClientRemarkResponse) {
    response = &ModifyClientRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 代理商可以对名下客户添加备注、修改备注
func (c *Client) ModifyClientRemark(request *ModifyClientRemarkRequest) (response *ModifyClientRemarkResponse, err error) {
    if request == nil {
        request = NewModifyClientRemarkRequest()
    }
    response = NewModifyClientRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewRemovePayRelationForClientRequest() (request *RemovePayRelationForClientRequest) {
    request = &RemovePayRelationForClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("partners", APIVersion, "RemovePayRelationForClient")
    return
}

func NewRemovePayRelationForClientResponse() (response *RemovePayRelationForClientResponse) {
    response = &RemovePayRelationForClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 合作伙伴为客户消除强代付关系
func (c *Client) RemovePayRelationForClient(request *RemovePayRelationForClientRequest) (response *RemovePayRelationForClientResponse, err error) {
    if request == nil {
        request = NewRemovePayRelationForClientRequest()
    }
    response = NewRemovePayRelationForClientResponse()
    err = c.Send(request, response)
    return
}
