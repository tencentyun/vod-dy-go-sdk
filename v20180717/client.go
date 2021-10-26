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
