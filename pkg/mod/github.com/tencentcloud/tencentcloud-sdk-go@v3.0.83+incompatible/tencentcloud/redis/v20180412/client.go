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

package v20180412

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-12"

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


func NewCleanUpInstanceRequest() (request *CleanUpInstanceRequest) {
    request = &CleanUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "CleanUpInstance")
    return
}

func NewCleanUpInstanceResponse() (response *CleanUpInstanceResponse) {
    response = &CleanUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 回收站实例立即下线
func (c *Client) CleanUpInstance(request *CleanUpInstanceRequest) (response *CleanUpInstanceResponse, err error) {
    if request == nil {
        request = NewCleanUpInstanceRequest()
    }
    response = NewCleanUpInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewClearInstanceRequest() (request *ClearInstanceRequest) {
    request = &ClearInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ClearInstance")
    return
}

func NewClearInstanceResponse() (response *ClearInstanceResponse) {
    response = &ClearInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 清空Redis实例的实例数据。
func (c *Client) ClearInstance(request *ClearInstanceRequest) (response *ClearInstanceResponse, err error) {
    if request == nil {
        request = NewClearInstanceRequest()
    }
    response = NewClearInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "CreateInstances")
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建redis实例
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoBackupConfigRequest() (request *DescribeAutoBackupConfigRequest) {
    request = &DescribeAutoBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeAutoBackupConfig")
    return
}

func NewDescribeAutoBackupConfigResponse() (response *DescribeAutoBackupConfigResponse) {
    response = &DescribeAutoBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取备份配置
func (c *Client) DescribeAutoBackupConfig(request *DescribeAutoBackupConfigRequest) (response *DescribeAutoBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAutoBackupConfigRequest()
    }
    response = NewDescribeAutoBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupUrlRequest() (request *DescribeBackupUrlRequest) {
    request = &DescribeBackupUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeBackupUrl")
    return
}

func NewDescribeBackupUrlResponse() (response *DescribeBackupUrlResponse) {
    response = &DescribeBackupUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询备份Rdb下载地址(接口灰度中，需要加白名单使用)
func (c *Client) DescribeBackupUrl(request *DescribeBackupUrlRequest) (response *DescribeBackupUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBackupUrlRequest()
    }
    response = NewDescribeBackupUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceBackupsRequest() (request *DescribeInstanceBackupsRequest) {
    request = &DescribeInstanceBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceBackups")
    return
}

func NewDescribeInstanceBackupsResponse() (response *DescribeInstanceBackupsResponse) {
    response = &DescribeInstanceBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询 CRS 实例备份列表
func (c *Client) DescribeInstanceBackups(request *DescribeInstanceBackupsRequest) (response *DescribeInstanceBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceBackupsRequest()
    }
    response = NewDescribeInstanceBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDealDetailRequest() (request *DescribeInstanceDealDetailRequest) {
    request = &DescribeInstanceDealDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceDealDetail")
    return
}

func NewDescribeInstanceDealDetailResponse() (response *DescribeInstanceDealDetailResponse) {
    response = &DescribeInstanceDealDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询订单信息
func (c *Client) DescribeInstanceDealDetail(request *DescribeInstanceDealDetailRequest) (response *DescribeInstanceDealDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDealDetailRequest()
    }
    response = NewDescribeInstanceDealDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamRecordsRequest() (request *DescribeInstanceParamRecordsRequest) {
    request = &DescribeInstanceParamRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceParamRecords")
    return
}

func NewDescribeInstanceParamRecordsResponse() (response *DescribeInstanceParamRecordsResponse) {
    response = &DescribeInstanceParamRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询参数修改历史列表
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamRecordsRequest()
    }
    response = NewDescribeInstanceParamRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
    request = &DescribeInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceParams")
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例参数列表
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSecurityGroupRequest() (request *DescribeInstanceSecurityGroupRequest) {
    request = &DescribeInstanceSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceSecurityGroup")
    return
}

func NewDescribeInstanceSecurityGroupResponse() (response *DescribeInstanceSecurityGroupResponse) {
    response = &DescribeInstanceSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例安全组信息
func (c *Client) DescribeInstanceSecurityGroup(request *DescribeInstanceSecurityGroupRequest) (response *DescribeInstanceSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSecurityGroupRequest()
    }
    response = NewDescribeInstanceSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceShardsRequest() (request *DescribeInstanceShardsRequest) {
    request = &DescribeInstanceShardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceShards")
    return
}

func NewDescribeInstanceShardsResponse() (response *DescribeInstanceShardsResponse) {
    response = &DescribeInstanceShardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群版实例分片信息
func (c *Client) DescribeInstanceShards(request *DescribeInstanceShardsRequest) (response *DescribeInstanceShardsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceShardsRequest()
    }
    response = NewDescribeInstanceShardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询Redis实例列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductInfoRequest() (request *DescribeProductInfoRequest) {
    request = &DescribeProductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProductInfo")
    return
}

func NewDescribeProductInfoResponse() (response *DescribeProductInfoResponse) {
    response = &DescribeProductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口查询指定可用区和实例类型下 Redis 的售卖规格， 如果用户不在购买白名单中，将不能查询该可用区或该类型的售卖规格详情。申请购买某地域白名单可以提交工单
func (c *Client) DescribeProductInfo(request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProductInfoRequest()
    }
    response = NewDescribeProductInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupRequest() (request *DescribeProjectSecurityGroupRequest) {
    request = &DescribeProjectSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProjectSecurityGroup")
    return
}

func NewDescribeProjectSecurityGroupResponse() (response *DescribeProjectSecurityGroupResponse) {
    response = &DescribeProjectSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询项目安全组信息
func (c *Client) DescribeProjectSecurityGroup(request *DescribeProjectSecurityGroupRequest) (response *DescribeProjectSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupRequest()
    }
    response = NewDescribeProjectSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
    request = &DescribeTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeTaskInfo")
    return
}

func NewDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
    response = &DescribeTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询任务结果
func (c *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInfoRequest()
    }
    response = NewDescribeTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPostpaidInstanceRequest() (request *DestroyPostpaidInstanceRequest) {
    request = &DestroyPostpaidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DestroyPostpaidInstance")
    return
}

func NewDestroyPostpaidInstanceResponse() (response *DestroyPostpaidInstanceResponse) {
    response = &DestroyPostpaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按量计费实例销毁
func (c *Client) DestroyPostpaidInstance(request *DestroyPostpaidInstanceRequest) (response *DestroyPostpaidInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyPostpaidInstanceRequest()
    }
    response = NewDestroyPostpaidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPrepaidInstanceRequest() (request *DestroyPrepaidInstanceRequest) {
    request = &DestroyPrepaidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DestroyPrepaidInstance")
    return
}

func NewDestroyPrepaidInstanceResponse() (response *DestroyPrepaidInstanceResponse) {
    response = &DestroyPrepaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 包年包月实例退还
func (c *Client) DestroyPrepaidInstance(request *DestroyPrepaidInstanceRequest) (response *DestroyPrepaidInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyPrepaidInstanceRequest()
    }
    response = NewDestroyPrepaidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDisableReplicaReadonlyRequest() (request *DisableReplicaReadonlyRequest) {
    request = &DisableReplicaReadonlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DisableReplicaReadonly")
    return
}

func NewDisableReplicaReadonlyResponse() (response *DisableReplicaReadonlyResponse) {
    response = &DisableReplicaReadonlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 禁用读写分离
func (c *Client) DisableReplicaReadonly(request *DisableReplicaReadonlyRequest) (response *DisableReplicaReadonlyResponse, err error) {
    if request == nil {
        request = NewDisableReplicaReadonlyRequest()
    }
    response = NewDisableReplicaReadonlyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableReplicaReadonlyRequest() (request *EnableReplicaReadonlyRequest) {
    request = &EnableReplicaReadonlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "EnableReplicaReadonly")
    return
}

func NewEnableReplicaReadonlyResponse() (response *EnableReplicaReadonlyResponse) {
    response = &EnableReplicaReadonlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用读写分离
func (c *Client) EnableReplicaReadonly(request *EnableReplicaReadonlyRequest) (response *EnableReplicaReadonlyResponse, err error) {
    if request == nil {
        request = NewEnableReplicaReadonlyRequest()
    }
    response = NewEnableReplicaReadonlyResponse()
    err = c.Send(request, response)
    return
}

func NewManualBackupInstanceRequest() (request *ManualBackupInstanceRequest) {
    request = &ManualBackupInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ManualBackupInstance")
    return
}

func NewManualBackupInstanceResponse() (response *ManualBackupInstanceResponse) {
    response = &ManualBackupInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 手动备份Redis实例
func (c *Client) ManualBackupInstance(request *ManualBackupInstanceRequest) (response *ManualBackupInstanceResponse, err error) {
    if request == nil {
        request = NewManualBackupInstanceRequest()
    }
    response = NewManualBackupInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModfiyInstancePasswordRequest() (request *ModfiyInstancePasswordRequest) {
    request = &ModfiyInstancePasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ModfiyInstancePassword")
    return
}

func NewModfiyInstancePasswordResponse() (response *ModfiyInstancePasswordResponse) {
    response = &ModfiyInstancePasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改redis密码
func (c *Client) ModfiyInstancePassword(request *ModfiyInstancePasswordRequest) (response *ModfiyInstancePasswordResponse, err error) {
    if request == nil {
        request = NewModfiyInstancePasswordRequest()
    }
    response = NewModfiyInstancePasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoBackupConfigRequest() (request *ModifyAutoBackupConfigRequest) {
    request = &ModifyAutoBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ModifyAutoBackupConfig")
    return
}

func NewModifyAutoBackupConfigResponse() (response *ModifyAutoBackupConfigResponse) {
    response = &ModifyAutoBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置自动备份时间
func (c *Client) ModifyAutoBackupConfig(request *ModifyAutoBackupConfigRequest) (response *ModifyAutoBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoBackupConfigRequest()
    }
    response = NewModifyAutoBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstance")
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改实例相关信息（目前支持：实例重命名）
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamsRequest() (request *ModifyInstanceParamsRequest) {
    request = &ModifyInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceParams")
    return
}

func NewModifyInstanceParamsResponse() (response *ModifyInstanceParamsResponse) {
    response = &ModifyInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改实例参数
func (c *Client) ModifyInstanceParams(request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamsRequest()
    }
    response = NewModifyInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkConfigRequest() (request *ModifyNetworkConfigRequest) {
    request = &ModifyNetworkConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ModifyNetworkConfig")
    return
}

func NewModifyNetworkConfigResponse() (response *ModifyNetworkConfigResponse) {
    response = &ModifyNetworkConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改实例网络配置
func (c *Client) ModifyNetworkConfig(request *ModifyNetworkConfigRequest) (response *ModifyNetworkConfigResponse, err error) {
    if request == nil {
        request = NewModifyNetworkConfigRequest()
    }
    response = NewModifyNetworkConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "RenewInstance")
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 续费实例
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    response = NewRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetPasswordRequest() (request *ResetPasswordRequest) {
    request = &ResetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ResetPassword")
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重置密码
func (c *Client) ResetPassword(request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    if request == nil {
        request = NewResetPasswordRequest()
    }
    response = NewResetPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreInstanceRequest() (request *RestoreInstanceRequest) {
    request = &RestoreInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "RestoreInstance")
    return
}

func NewRestoreInstanceResponse() (response *RestoreInstanceResponse) {
    response = &RestoreInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 恢复 CRS 实例
func (c *Client) RestoreInstance(request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    if request == nil {
        request = NewRestoreInstanceRequest()
    }
    response = NewRestoreInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchInstanceVipRequest() (request *SwitchInstanceVipRequest) {
    request = &SwitchInstanceVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "SwitchInstanceVip")
    return
}

func NewSwitchInstanceVipResponse() (response *SwitchInstanceVipResponse) {
    response = &SwitchInstanceVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在通过DTS支持跨可用区灾备的场景中，通过该接口交换实例VIP完成实例灾备切换。交换VIP后目标实例可写，源和目标实例VIP互换，同时源与目标实例间DTS同步任务断开
func (c *Client) SwitchInstanceVip(request *SwitchInstanceVipRequest) (response *SwitchInstanceVipResponse, err error) {
    if request == nil {
        request = NewSwitchInstanceVipRequest()
    }
    response = NewSwitchInstanceVipResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeInstance")
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 升级实例
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}
