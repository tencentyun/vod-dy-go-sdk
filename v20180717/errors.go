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

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 没有开通点播业务。
	FAILEDOPERATION_INVALIDVODUSER = "FailedOperation.InvalidVodUser"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 参数错误：文件不支持解冻。
	INVALIDPARAMETERVALUE_NOTRESTORABLE = "InvalidParameterValue.NotRestorable"

	// 参数错误：解冻天数错误。
	INVALIDPARAMETERVALUE_RESTOREDAY = "InvalidParameterValue.RestoreDay"

	// SessionContext 过长。
	INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"

	// 去重识别码重复，请求被去重。
	INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"

	// SessionId 过长。
	INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不存在：文件不存在。
	RESOURCENOTFOUND_FILENOTEXIST = "ResourceNotFound.FileNotExist"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
