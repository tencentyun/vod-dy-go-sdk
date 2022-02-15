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

package v20180717

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-17"

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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDeleteMediaForDYRequest() (request *DeleteMediaForDYRequest) {
    request = &DeleteMediaForDYRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DeleteMediaForDY")
    
    
    return
}

func NewDeleteMediaForDYResponse() (response *DeleteMediaForDYResponse) {
    response = &DeleteMediaForDYResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMediaForDY
// 删除指定文件，如果是 m3u8 文件，会同时删除 ts 文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMediaForDY(request *DeleteMediaForDYRequest) (response *DeleteMediaForDYResponse, err error) {
    if request == nil {
        request = NewDeleteMediaForDYRequest()
    }
    
    response = NewDeleteMediaForDYResponse()
    err = c.Send(request, response)
    return
}

// DeleteMediaForDY
// 删除指定文件，如果是 m3u8 文件，会同时删除 ts 文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMediaForDYWithContext(ctx context.Context, request *DeleteMediaForDYRequest) (response *DeleteMediaForDYResponse, err error) {
    if request == nil {
        request = NewDeleteMediaForDYRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteMediaForDYResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaMetaDataForDYRequest() (request *DescribeMediaMetaDataForDYRequest) {
    request = &DescribeMediaMetaDataForDYRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeMediaMetaDataForDY")
    
    
    return
}

func NewDescribeMediaMetaDataForDYResponse() (response *DescribeMediaMetaDataForDYResponse) {
    response = &DescribeMediaMetaDataForDYResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMediaMetaDataForDY
// 获取媒体的元信息，包括视频画面宽、高、编码格式、时长、帧率等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) DescribeMediaMetaDataForDY(request *DescribeMediaMetaDataForDYRequest) (response *DescribeMediaMetaDataForDYResponse, err error) {
    if request == nil {
        request = NewDescribeMediaMetaDataForDYRequest()
    }
    
    response = NewDescribeMediaMetaDataForDYResponse()
    err = c.Send(request, response)
    return
}

// DescribeMediaMetaDataForDY
// 获取媒体的元信息，包括视频画面宽、高、编码格式、时长、帧率等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"
func (c *Client) DescribeMediaMetaDataForDYWithContext(ctx context.Context, request *DescribeMediaMetaDataForDYRequest) (response *DescribeMediaMetaDataForDYResponse, err error) {
    if request == nil {
        request = NewDescribeMediaMetaDataForDYRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMediaMetaDataForDYResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailForDYRequest() (request *DescribeTaskDetailForDYRequest) {
    request = &DescribeTaskDetailForDYRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "DescribeTaskDetailForDY")
    
    
    return
}

func NewDescribeTaskDetailForDYResponse() (response *DescribeTaskDetailForDYResponse) {
    response = &DescribeTaskDetailForDYResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetailForDY
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询3天之内提交的任务）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetailForDY(request *DescribeTaskDetailForDYRequest) (response *DescribeTaskDetailForDYResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailForDYRequest()
    }
    
    response = NewDescribeTaskDetailForDYResponse()
    err = c.Send(request, response)
    return
}

// DescribeTaskDetailForDY
// 通过任务 ID 查询任务的执行状态和结果的详细信息（最多可以查询3天之内提交的任务）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskDetailForDYWithContext(ctx context.Context, request *DescribeTaskDetailForDYRequest) (response *DescribeTaskDetailForDYResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailForDYRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskDetailForDYResponse()
    err = c.Send(request, response)
    return
}

func NewEditMediaForDYRequest() (request *EditMediaForDYRequest) {
    request = &EditMediaForDYRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "EditMediaForDY")
    
    
    return
}

func NewEditMediaForDYResponse() (response *EditMediaForDYResponse) {
    response = &EditMediaForDYResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditMediaForDY
// 对视频进行编辑（剪辑、拼接等），生成一个新的点播视频。编辑的功能包括：
//
// 
//
// 1. 对一个文件进行剪辑，生成一个新的视频；
//
// 2. 对多个文件进行拼接，生成一个新的视频；
//
// 3. 对多个文件进行剪辑，然后再拼接，生成一个新的视频。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMediaForDY(request *EditMediaForDYRequest) (response *EditMediaForDYResponse, err error) {
    if request == nil {
        request = NewEditMediaForDYRequest()
    }
    
    response = NewEditMediaForDYResponse()
    err = c.Send(request, response)
    return
}

// EditMediaForDY
// 对视频进行编辑（剪辑、拼接等），生成一个新的点播视频。编辑的功能包括：
//
// 
//
// 1. 对一个文件进行剪辑，生成一个新的视频；
//
// 2. 对多个文件进行拼接，生成一个新的视频；
//
// 3. 对多个文件进行剪辑，然后再拼接，生成一个新的视频。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EditMediaForDYWithContext(ctx context.Context, request *EditMediaForDYRequest) (response *EditMediaForDYResponse, err error) {
    if request == nil {
        request = NewEditMediaForDYRequest()
    }
    request.SetContext(ctx)
    
    response = NewEditMediaForDYResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMediaForDYRequest() (request *ProcessMediaForDYRequest) {
    request = &ProcessMediaForDYRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "ProcessMediaForDY")
    
    
    return
}

func NewProcessMediaForDYResponse() (response *ProcessMediaForDYResponse) {
    response = &ProcessMediaForDYResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ProcessMediaForDY
// 对 COS 中的媒体文件发起处理任务，功能包括：
//
// 1. 视频转码（带水印）；
//
// 2. 视频转动图；
//
// 3. 对视频按指定时间点截图；
//
// 4. 对视频采样截图；
//
// 5. 对视频截图雪碧图；
//
// 6. 对视频转自适应码流；
//
// 7. 智能内容审核（鉴黄、鉴恐、鉴政）；
//
// 8. 智能内容分析（标签、分类、封面、按帧标签）；
//
// 9. 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessMediaForDY(request *ProcessMediaForDYRequest) (response *ProcessMediaForDYResponse, err error) {
    if request == nil {
        request = NewProcessMediaForDYRequest()
    }
    
    response = NewProcessMediaForDYResponse()
    err = c.Send(request, response)
    return
}

// ProcessMediaForDY
// 对 COS 中的媒体文件发起处理任务，功能包括：
//
// 1. 视频转码（带水印）；
//
// 2. 视频转动图；
//
// 3. 对视频按指定时间点截图；
//
// 4. 对视频采样截图；
//
// 5. 对视频截图雪碧图；
//
// 6. 对视频转自适应码流；
//
// 7. 智能内容审核（鉴黄、鉴恐、鉴政）；
//
// 8. 智能内容分析（标签、分类、封面、按帧标签）；
//
// 9. 智能内容识别（人脸、文本全文、文本关键词、语音全文、语音关键词）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"
//  INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"
//  INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"
func (c *Client) ProcessMediaForDYWithContext(ctx context.Context, request *ProcessMediaForDYRequest) (response *ProcessMediaForDYResponse, err error) {
    if request == nil {
        request = NewProcessMediaForDYRequest()
    }
    request.SetContext(ctx)
    
    response = NewProcessMediaForDYResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreMediaForDYRequest() (request *RestoreMediaForDYRequest) {
    request = &RestoreMediaForDYRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vod", APIVersion, "RestoreMediaForDY")
    
    
    return
}

func NewRestoreMediaForDYResponse() (response *RestoreMediaForDYResponse) {
    response = &RestoreMediaForDYResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestoreMediaForDY
// 当媒体文件的存储类型是归档存储或深度归档存储时，是不可访问的。如需访问，则需要调用本接口进行解冻，解冻后可访问的媒体文件是临时的，在有效期过后，则不可访问。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTRESTORABLE = "InvalidParameterValue.NotRestorable"
//  INVALIDPARAMETERVALUE_RESTOREDAY = "InvalidParameterValue.RestoreDay"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestoreMediaForDY(request *RestoreMediaForDYRequest) (response *RestoreMediaForDYResponse, err error) {
    if request == nil {
        request = NewRestoreMediaForDYRequest()
    }
    
    response = NewRestoreMediaForDYResponse()
    err = c.Send(request, response)
    return
}

// RestoreMediaForDY
// 当媒体文件的存储类型是归档存储或深度归档存储时，是不可访问的。如需访问，则需要调用本接口进行解冻，解冻后可访问的媒体文件是临时的，在有效期过后，则不可访问。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTRESTORABLE = "InvalidParameterValue.NotRestorable"
//  INVALIDPARAMETERVALUE_RESTOREDAY = "InvalidParameterValue.RestoreDay"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestoreMediaForDYWithContext(ctx context.Context, request *RestoreMediaForDYRequest) (response *RestoreMediaForDYResponse, err error) {
    if request == nil {
        request = NewRestoreMediaForDYRequest()
    }
    request.SetContext(ctx)
    
    response = NewRestoreMediaForDYResponse()
    err = c.Send(request, response)
    return
}
