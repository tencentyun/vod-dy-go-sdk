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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AdaptiveDynamicStreamingInfoItemForDY struct {
	// 转自适应码流规格。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 打包格式，可能为 HLS和 MPEG-DASH 两种。
	Package *string `json:"Package,omitnil,omitempty" name:"Package"`

	// 播放路径。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 自适应码流文件的存储位置。
	Storage *TaskOutputStorageForDY `json:"Storage,omitnil,omitempty" name:"Storage"`
}

type AdaptiveDynamicStreamingTaskInputForDY struct {
	// 转自适应码流模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInputForDY `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// 转自适应码流后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 转自适应码流后，manifest 文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// 转自适应码流后，子流文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}.{format}`。
	SubStreamObjectName *string `json:"SubStreamObjectName,omitnil,omitempty" name:"SubStreamObjectName"`

	// 转自适应码流（仅 HLS）后，分片文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}_{segmentNumber}.{format}`。
	SegmentObjectName *string `json:"SegmentObjectName,omitnil,omitempty" name:"SegmentObjectName"`
}

type AnimatedGraphicTaskInputForDY struct {
	// 视频转动图模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 动图在视频中的开始时间，单位为秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间，单位为秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 转动图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 转动图后文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_animatedGraphic_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`
}

type CosInputInfoForDY struct {
	// 视频处理对象文件所在的 COS Bucket 名，如 TopRankVideo-125xxx88。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 视频处理对象文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 视频处理对象文件的输入路径，如`/movie/201907/WildAnimal.mov`。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`
}

type CosOutputStorageForDY struct {
	// 视频处理生成的文件输出的目标 Bucket 名，如 TopRankVideo-125xxx88。如果不填，表示继承上层。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 视频处理生成的文件输出的目标 Bucket 的园区，如 ap-chongqing。如果不填，表示继承上层。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type DeleteMediaForDYRequestParams struct {
	// 文件所在的 COS Bucket 名，如 wsd-tx-ugc-pub。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// 文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 文件的 COS 完整路径。
	ObjectSet []*string `json:"ObjectSet,omitnil,omitempty" name:"ObjectSet"`

	// 来源上下文，用于透传用户请求信息，删除回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type DeleteMediaForDYRequest struct {
	*tchttp.BaseRequest
	
	// 文件所在的 COS Bucket 名，如 wsd-tx-ugc-pub。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// 文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 文件的 COS 完整路径。
	ObjectSet []*string `json:"ObjectSet,omitnil,omitempty" name:"ObjectSet"`

	// 来源上下文，用于透传用户请求信息，删除回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

func (r *DeleteMediaForDYRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaForDYRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosBucket")
	delete(f, "CosRegion")
	delete(f, "ObjectSet")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMediaForDYRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMediaForDYResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMediaForDYResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMediaForDYResponseParams `json:"Response"`
}

func (r *DeleteMediaForDYResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaForDYResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaMetaDataForDYRequestParams struct {
	// 需要获取元信息的文件输入信息。
	InputInfo *MediaInputInfoForDY `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
}

type DescribeMediaMetaDataForDYRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取元信息的文件输入信息。
	InputInfo *MediaInputInfoForDY `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
}

func (r *DescribeMediaMetaDataForDYRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaMetaDataForDYRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaMetaDataForDYRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaMetaDataForDYResponseParams struct {
	// 媒体元信息。
	MetaData *MediaMetaDataForDY `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMediaMetaDataForDYResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaMetaDataForDYResponseParams `json:"Response"`
}

func (r *DescribeMediaMetaDataForDYResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaMetaDataForDYResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailForDYRequestParams struct {
	// 视频处理任务的任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskDetailForDYRequest struct {
	*tchttp.BaseRequest
	
	// 视频处理任务的任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskDetailForDYRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailForDYRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailForDYRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailForDYResponseParams struct {
	// 任务类型，目前取值有：
	// <li>WorkflowTask：视频工作流处理任务。</li>
	// <li>EditMediaTask：视频编辑任务。</li>
	// <li>LiveStreamProcessTask：直播流处理任务。</li>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态，取值：
	// <li>WAITING：等待中；</li>
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务的创建时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	BeginProcessTime *string `json:"BeginProcessTime,omitnil,omitempty" name:"BeginProcessTime"`

	// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/862/37710#52)。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 视频处理任务信息，仅当 TaskType 为 WorkflowTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTask *WorkflowTaskForDY `json:"WorkflowTask,omitnil,omitempty" name:"WorkflowTask"`

	// 视频编辑任务信息，仅当 TaskType 为 EditMediaTask，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditMediaTask *EditMediaTaskForDY `json:"EditMediaTask,omitnil,omitempty" name:"EditMediaTask"`

	// 任务流的优先级，取值范围为 [-10, 10]。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长50个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长1000个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 扩展信息字段，仅用于特定场景。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskDetailForDYResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskDetailForDYResponseParams `json:"Response"`
}

func (r *DescribeTaskDetailForDYResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailForDYResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditMediaFileInfoForDY struct {
	// 视频的输入信息。
	InputInfo *MediaInputInfoForDY `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 视频剪辑的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 视频剪辑的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

// Predefined struct for user
type EditMediaForDYRequestParams struct {
	// 输入的视频文件信息。
	FileInfos []*EditMediaFileInfoForDY `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 视频处理输出文件的目标存储。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 视频处理输出文件的目标路径。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// 编辑后生成的文件配置。
	OutputConfig *EditMediaOutputConfigForDY `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 任务优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type EditMediaForDYRequest struct {
	*tchttp.BaseRequest
	
	// 输入的视频文件信息。
	FileInfos []*EditMediaFileInfoForDY `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 视频处理输出文件的目标存储。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 视频处理输出文件的目标路径。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// 编辑后生成的文件配置。
	OutputConfig *EditMediaOutputConfigForDY `json:"OutputConfig,omitnil,omitempty" name:"OutputConfig"`

	// 任务优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 保留字段，特殊用途时使用。
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

func (r *EditMediaForDYRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditMediaForDYRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileInfos")
	delete(f, "OutputStorage")
	delete(f, "OutputObjectPath")
	delete(f, "OutputConfig")
	delete(f, "TasksPriority")
	delete(f, "ExtInfo")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditMediaForDYRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditMediaForDYResponseParams struct {
	// 编辑视频的任务 ID，可以通过该 ID 查询编辑任务的状态。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EditMediaForDYResponse struct {
	*tchttp.BaseResponse
	Response *EditMediaForDYResponseParams `json:"Response"`
}

func (r *EditMediaForDYResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditMediaForDYResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditMediaOutputConfigForDY struct {
	// 封装格式，可选值：mp4、hls、mov、flv、avi。默认是 mp4。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`
}

type EditMediaTaskForDY struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 视频编辑任务的输入。
	Input *EditMediaTaskInputForDY `json:"Input,omitnil,omitempty" name:"Input"`

	// 视频编辑任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *EditMediaTaskOutputForDY `json:"Output,omitnil,omitempty" name:"Output"`
}

type EditMediaTaskInputForDY struct {
	// 输入的视频文件信息。
	FileInfoSet []*EditMediaFileInfoForDY `json:"FileInfoSet,omitnil,omitempty" name:"FileInfoSet"`
}

type EditMediaTaskOutputForDY struct {
	// 编辑后文件的目标存储。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 编辑后的视频文件路径。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 云点播对编辑后的视频文件发起 ProcessMediaForDY 任务的任务 ID。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitnil,omitempty" name:"ProcedureTaskId"`
}

type ImageSpriteTaskInputForDY struct {
	// 雪碧图模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 截取雪碧图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 截取雪碧图后，雪碧图图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_imageSprite_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// 截取雪碧图后，Web VTT 文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_imageSprite_{definition}.{format}`。
	WebVttObjectName *string `json:"WebVttObjectName,omitnil,omitempty" name:"WebVttObjectName"`

	// 截取雪碧图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormatForDY `json:"ObjectNumberFormat,omitnil,omitempty" name:"ObjectNumberFormat"`
}

type MediaAnimatedGraphicsItemForDY struct {
	// 转动图文件的存储位置。
	Storage *TaskOutputStorageForDY `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 转动图的文件路径。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 转动图模板 ID，参见[转动图参数模板](https://cloud.tencent.com/document/product/862/37042#.E9.A2.84.E7.BD.AE.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 动图格式，如 gif。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 动图的高度，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 动图的宽度，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 动图码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 动图大小，单位：字节。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 动图的md5值。
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 动图在视频中的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type MediaAudioStreamItemForDY struct {
	// 音频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 音频流的采样率，单位：hz。
	SamplingRate *int64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// 音频流的编码格式，例如 aac。
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`
}

type MediaImageSpriteItemForDY struct {
	// 雪碧图规格，参见[雪碧图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 雪碧图小图的高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 雪碧图小图的宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 每一张雪碧图大图里小图的数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 每一张雪碧图大图的路径。
	ImagePathSet []*string `json:"ImagePathSet,omitnil,omitempty" name:"ImagePathSet"`

	// 雪碧图子图位置与时间关系的 WebVtt 文件路径。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	WebVttPath *string `json:"WebVttPath,omitnil,omitempty" name:"WebVttPath"`

	// 雪碧图文件的存储位置。
	Storage *TaskOutputStorageForDY `json:"Storage,omitnil,omitempty" name:"Storage"`
}

type MediaInputInfoForDY struct {
	// 输入来源对象的类型，支持 COS 和 URL 两种。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当 Type 为 COS 时有效，则该项为必填，表示视频处理 COS 对象信息。
	CosInputInfo *CosInputInfoForDY `json:"CosInputInfo,omitnil,omitempty" name:"CosInputInfo"`
}

type MediaMetaDataForDY struct {
	// 上传的媒体文件大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度。
	Rotate *int64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItemForDY `json:"VideoStreamSet,omitnil,omitempty" name:"VideoStreamSet"`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItemForDY `json:"AudioStreamSet,omitnil,omitempty" name:"AudioStreamSet"`

	// 视频时长，单位：秒。
	VideoDuration *float64 `json:"VideoDuration,omitnil,omitempty" name:"VideoDuration"`

	// 音频时长，单位：秒。
	AudioDuration *float64 `json:"AudioDuration,omitnil,omitempty" name:"AudioDuration"`
}

type MediaProcessTaskAdaptiveDynamicStreamingResultForDY struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频转自适应码流任务的输入。
	Input *AdaptiveDynamicStreamingTaskInputForDY `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频转自适应码流任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AdaptiveDynamicStreamingInfoItemForDY `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaProcessTaskAnimatedGraphicResultForDY struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 转动图任务的输入。
	Input *AnimatedGraphicTaskInputForDY `json:"Input,omitnil,omitempty" name:"Input"`

	// 转动图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaAnimatedGraphicsItemForDY `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaProcessTaskImageSpriteResultForDY struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频截雪碧图任务的输入。
	Input *ImageSpriteTaskInputForDY `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频截雪碧图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaImageSpriteItemForDY `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaProcessTaskInputForDY struct {
	// 视频转码任务列表。
	TranscodeTaskSet []*TranscodeTaskInputForDY `json:"TranscodeTaskSet,omitnil,omitempty" name:"TranscodeTaskSet"`

	// 视频转动图任务列表。
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInputForDY `json:"AnimatedGraphicTaskSet,omitnil,omitempty" name:"AnimatedGraphicTaskSet"`

	// 对视频按时间点截图任务列表。
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInputForDY `json:"SnapshotByTimeOffsetTaskSet,omitnil,omitempty" name:"SnapshotByTimeOffsetTaskSet"`

	// 对视频采样截图任务列表。
	SampleSnapshotTaskSet []*SampleSnapshotTaskInputForDY `json:"SampleSnapshotTaskSet,omitnil,omitempty" name:"SampleSnapshotTaskSet"`

	// 对视频截雪碧图任务列表。
	ImageSpriteTaskSet []*ImageSpriteTaskInputForDY `json:"ImageSpriteTaskSet,omitnil,omitempty" name:"ImageSpriteTaskSet"`

	// 转自适应码流任务列表。
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInputForDY `json:"AdaptiveDynamicStreamingTaskSet,omitnil,omitempty" name:"AdaptiveDynamicStreamingTaskSet"`
}

type MediaProcessTaskResultForDY struct {
	// 任务的类型，可以取的值有：
	// <li>Transcode：转码</li>
	// <li>AnimatedGraphics：转动图</li>
	// <li>SnapshotByTimeOffset：时间点截图</li>
	// <li>SampleSnapshot：采样截图</li>
	// <li>ImageSprites：雪碧图</li>
	// <li>CoverBySnapshot：截图做封面</li>
	// <li>AdaptiveDynamicStreaming：自适应码流</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 视频转码任务的查询结果，当任务类型为 Transcode 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *MediaProcessTaskTranscodeResultForDY `json:"TranscodeTask,omitnil,omitempty" name:"TranscodeTask"`

	// 视频转动图任务的查询结果，当任务类型为 AnimatedGraphics 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResultForDY `json:"AnimatedGraphicTask,omitnil,omitempty" name:"AnimatedGraphicTask"`

	// 对视频按时间点截图任务的查询结果，当任务类型为 SnapshotByTimeOffset 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResultForDY `json:"SnapshotByTimeOffsetTask,omitnil,omitempty" name:"SnapshotByTimeOffsetTask"`

	// 对视频采样截图任务的查询结果，当任务类型为 SampleSnapshot 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResultForDY `json:"SampleSnapshotTask,omitnil,omitempty" name:"SampleSnapshotTask"`

	// 对视频截雪碧图任务的查询结果，当任务类型为 ImageSprite 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTask *MediaProcessTaskImageSpriteResultForDY `json:"ImageSpriteTask,omitnil,omitempty" name:"ImageSpriteTask"`

	// 转自适应码流任务查询结果，当任务类型为 AdaptiveDynamicStreaming 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResultForDY `json:"AdaptiveDynamicStreamingTask,omitnil,omitempty" name:"AdaptiveDynamicStreamingTask"`
}

type MediaProcessTaskSampleSnapshotResultForDY struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频做采样截图任务输入。
	Input *SampleSnapshotTaskInputForDY `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频做采样截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSampleSnapshotItemForDY `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaProcessTaskSnapshotByTimeOffsetResultForDY struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 对视频按指定时间点截图任务输入。
	Input *SnapshotByTimeOffsetTaskInputForDY `json:"Input,omitnil,omitempty" name:"Input"`

	// 对视频按指定时间点截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSnapshotByTimeOffsetItemForDY `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaProcessTaskTranscodeResultForDY struct {
	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码，空字符串表示成功，其他值表示失败，取值请参考 [视频处理类错误码](https://cloud.tencent.com/document/product/862/50369#.E8.A7.86.E9.A2.91.E5.A4.84.E7.90.86.E7.B1.BB.E9.94.99.E8.AF.AF.E7.A0.81) 列表。
	ErrCodeExt *string `json:"ErrCodeExt,omitnil,omitempty" name:"ErrCodeExt"`

	// 错误码，0 表示成功，其他值表示失败（该字段已不推荐使用，建议使用新的错误码字段 ErrCodeExt）。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 转码任务的输入。
	Input *TranscodeTaskInputForDY `json:"Input,omitnil,omitempty" name:"Input"`

	// 转码任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaTranscodeItemForDY `json:"Output,omitnil,omitempty" name:"Output"`
}

type MediaSampleSnapshotItemForDY struct {
	// 采样截图规格 ID，参见[采样截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 采样方式，取值范围：
	// <li>Percent：根据百分比间隔采样。</li>
	// <li>Time：根据时间间隔采样。</li>
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 采样间隔
	// <li>当 SampleType 为 Percent 时，该值表示多少百分比一张图。</li>
	// <li>当 SampleType 为 Time 时，该值表示多少时间间隔一张图，单位秒， 第一张图均为视频首帧。</li>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 截图后文件的存储位置。
	Storage *TaskOutputStorageForDY `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 生成的截图 path 列表。
	ImagePathSet []*string `json:"ImagePathSet,omitnil,omitempty" name:"ImagePathSet"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitnil,omitempty" name:"WaterMarkDefinition"`
}

type MediaSnapshotByTimeOffsetItemForDY struct {
	// 指定时间点截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 同一规格的截图信息集合，每个元素代表一张截图。
	PicInfoSet []*MediaSnapshotByTimePicInfoItemForDY `json:"PicInfoSet,omitnil,omitempty" name:"PicInfoSet"`

	// 指定时间点截图文件的存储位置。
	Storage *TaskOutputStorageForDY `json:"Storage,omitnil,omitempty" name:"Storage"`
}

type MediaSnapshotByTimePicInfoItemForDY struct {
	// 该张截图对应视频文件中的时间偏移，单位为<font color=red>毫秒</font>。
	TimeOffset *float64 `json:"TimeOffset,omitnil,omitempty" name:"TimeOffset"`

	// 该张截图的路径。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitnil,omitempty" name:"WaterMarkDefinition"`
}

type MediaTranscodeItemForDY struct {
	// 转码后文件的目标存储。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 转码后的视频文件路径。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/862/37042)。
	Definition *int64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 媒体文件总大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 视频的 md5 值。
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItemForDY `json:"AudioStreamSet,omitnil,omitempty" name:"AudioStreamSet"`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItemForDY `json:"VideoStreamSet,omitnil,omitempty" name:"VideoStreamSet"`
}

type MediaVideoStreamItemForDY struct {
	// 视频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 视频流的高度，单位：px。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频流的宽度，单位：px。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频流的编码格式，例如 h264。
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// 帧率，单位：hz。
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`
}

type NumberFormatForDY struct {
	// `{number}`变量的起始值，默认为0。
	InitialValue *uint64 `json:"InitialValue,omitnil,omitempty" name:"InitialValue"`

	// `{number}`变量的增长步长，默认为1。
	Increment *uint64 `json:"Increment,omitnil,omitempty" name:"Increment"`

	// `{number}`变量的最小长度，不足时补占位符。默认为1。
	MinLength *uint64 `json:"MinLength,omitnil,omitempty" name:"MinLength"`

	// `{number}`变量的长度不足时，补充的占位符。默认为"0"。
	PlaceHolder *string `json:"PlaceHolder,omitnil,omitempty" name:"PlaceHolder"`
}

// Predefined struct for user
type ProcessMediaForDYRequestParams struct {
	// 视频处理的文件输入信息。
	InputInfo *MediaInputInfoForDY `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 视频处理输出文件的目标存储。
	// 当 InputInfo 的输入类型是 COS 时可以不填，不填代表继承 InputInfo 中的存储位置。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。如果不填，表示与 InputInfo 中文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInputForDY `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

type ProcessMediaForDYRequest struct {
	*tchttp.BaseRequest
	
	// 视频处理的文件输入信息。
	InputInfo *MediaInputInfoForDY `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 视频处理输出文件的目标存储。
	// 当 InputInfo 的输入类型是 COS 时可以不填，不填代表继承 InputInfo 中的存储位置。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。如果不填，表示与 InputInfo 中文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInputForDY `json:"MediaProcessTask,omitnil,omitempty" name:"MediaProcessTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitnil,omitempty" name:"TasksPriority"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`
}

func (r *ProcessMediaForDYRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaForDYRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputInfo")
	delete(f, "OutputStorage")
	delete(f, "OutputDir")
	delete(f, "MediaProcessTask")
	delete(f, "TasksPriority")
	delete(f, "SessionId")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProcessMediaForDYRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProcessMediaForDYResponseParams struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProcessMediaForDYResponse struct {
	*tchttp.BaseResponse
	Response *ProcessMediaForDYResponseParams `json:"Response"`
}

func (r *ProcessMediaForDYResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProcessMediaForDYResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RawImageWatermarkInputForDY struct {
	// 水印图片的输入内容。支持 jpeg、png 图片格式。
	ImageContent *MediaInputInfoForDY `json:"ImageContent,omitnil,omitempty" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：10%。
	Width *string `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素。</li>
	// 默认值：0px，表示 Height 按照原始水印图片的宽高比缩放。
	Height *string `json:"Height,omitnil,omitempty" name:"Height"`

	// 印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束（默认值）。</li>
	RepeatType *string `json:"RepeatType,omitnil,omitempty" name:"RepeatType"`
}

type RawWatermarkParameterForDY struct {
	// 水印类型，可选值：
	// <li>image：图片水印。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 原点位置，目前仅支持：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitnil,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitnil,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitnil,omitempty" name:"YPos"`

	// 图片水印模板，当 Type 为 image，该字段必填。当 Type 为 text，该字段无效。
	ImageTemplate *RawImageWatermarkInputForDY `json:"ImageTemplate,omitnil,omitempty" name:"ImageTemplate"`
}

// Predefined struct for user
type RestoreMediaForDYRequestParams struct {
	// 文件所在的 COS Bucket 名，如 bucket-xxx。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// 文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 文件的 COS 完整路径。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 解冻出的临时媒体文件的可访问持续时长，单位为“天”。默认为1天。
	RestoreDay *uint64 `json:"RestoreDay,omitnil,omitempty" name:"RestoreDay"`

	// 来源上下文，用于透传用户请求信息，最长 1000 个字符。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`
}

type RestoreMediaForDYRequest struct {
	*tchttp.BaseRequest
	
	// 文件所在的 COS Bucket 名，如 bucket-xxx。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// 文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 文件的 COS 完整路径。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 解冻出的临时媒体文件的可访问持续时长，单位为“天”。默认为1天。
	RestoreDay *uint64 `json:"RestoreDay,omitnil,omitempty" name:"RestoreDay"`

	// 来源上下文，用于透传用户请求信息，最长 1000 个字符。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`
}

func (r *RestoreMediaForDYRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreMediaForDYRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosBucket")
	delete(f, "CosRegion")
	delete(f, "Object")
	delete(f, "RestoreDay")
	delete(f, "SourceContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreMediaForDYRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreMediaForDYResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreMediaForDYResponse struct {
	*tchttp.BaseResponse
	Response *RestoreMediaForDYResponseParams `json:"Response"`
}

func (r *RestoreMediaForDYResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreMediaForDYResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SampleSnapshotTaskInputForDY struct {
	// 采样截图模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInputForDY `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// 采样截图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 采样截图后图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_sampleSnapshot_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// 采样截图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormatForDY `json:"ObjectNumberFormat,omitnil,omitempty" name:"ObjectNumberFormat"`
}

type SnapshotByTimeOffsetTaskInputForDY struct {
	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 截图时间点列表，时间点支持 s、% 两种格式：
	// <li>当字符串以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// <li>当字符串以 % 结尾，表示时间点为视频时长的百分比大小，如10%表示时间点为视频前第10%的时间。</li>
	ExtTimeOffsetSet []*string `json:"ExtTimeOffsetSet,omitnil,omitempty" name:"ExtTimeOffsetSet"`

	// 截图时间点列表，单位为<font color=red>秒</font>。此参数已不再建议使用，建议您使用 ExtTimeOffsetSet 参数。
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitnil,omitempty" name:"TimeOffsetSet"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInputForDY `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// 时间点截图后文件的目标存储，不填则继承上层的 OutputStorage 值。<li>当字符串以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// <li>当字符串以 % 结尾，表示时间点为视频时长的百分比大小，如10%表示时间点为视频前第10%的时间。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 时间点截图后图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_snapshotByTimeOffset_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// 时间点截图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormatForDY `json:"ObjectNumberFormat,omitnil,omitempty" name:"ObjectNumberFormat"`
}

type TaskOutputStorageForDY struct {
	// 视频处理输出对象存储位置的类型，现在仅支持 COS。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当 Type 为 COS 时有效，则该项为必填，表示视频处理 COS 输出位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosOutputStorage *CosOutputStorageForDY `json:"CosOutputStorage,omitnil,omitempty" name:"CosOutputStorage"`
}

type TranscodeTaskInputForDY struct {
	// 视频转码模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInputForDY `json:"WatermarkSet,omitnil,omitempty" name:"WatermarkSet"`

	// 转码后的视频的起始时间偏移，单位：秒。
	// <li>不填或填0，表示转码后的视频从原始视频的起始位置开始；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频从原始视频的第 n 秒位置开始；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频从原始视频结束 n 秒前的位置开始。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 转码后视频的终止时间偏移，单位：秒。
	// <li>不填或填0，表示转码后的视频持续到原始视频的末尾终止；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频持续到原始视频第 n 秒时终止；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频持续到原始视频结束 n 秒前终止。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 转码后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitnil,omitempty" name:"OutputStorage"`

	// 转码后主文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_transcode_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitnil,omitempty" name:"OutputObjectPath"`

	// 码后分片文件的输出路径（转码 HLS 时 ts 的路径），只能为相对路径。如果不填，则默认为：`{inputName}_transcode_{definition}_{number}.{format}`。
	SegmentObjectName *string `json:"SegmentObjectName,omitnil,omitempty" name:"SegmentObjectName"`

	// 转码后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormatForDY `json:"ObjectNumberFormat,omitnil,omitempty" name:"ObjectNumberFormat"`
}

type WatermarkInputForDY struct {
	// 水印模板 ID。
	Definition *uint64 `json:"Definition,omitnil,omitempty" name:"Definition"`

	// 水印自定义参数，当 Definition 填 0 时有效。
	// 该参数用于高度定制场景，建议您优先使用 Definition 指定水印参数。
	// 水印自定义参数不支持截图打水印。
	RawParameter *RawWatermarkParameterForDY `json:"RawParameter,omitnil,omitempty" name:"RawParameter"`

	// 文字内容，长度不超过100个字符。仅当水印类型为文字水印时填写。
	// 文字水印不支持截图打水印。
	TextContent *string `json:"TextContent,omitnil,omitempty" name:"TextContent"`

	// SVG 内容。长度不超过 2000000 个字符。仅当水印类型为 SVG 水印时填写。
	// SVG 水印不支持截图打水印。
	SvgContent *string `json:"SvgContent,omitnil,omitempty" name:"SvgContent"`

	// 水印的起始时间偏移，单位：秒。不填或填0，表示水印从画面出现时开始显现。
	// <li>不填或填0，表示水印从画面开始就出现；</li>
	// <li>当数值大于0时（假设为 n），表示水印从画面开始的第 n 秒出现；</li>
	// <li>当数值小于0时（假设为 -n），表示水印从离画面结束 n 秒前开始出现。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 水印的结束时间偏移，单位：秒。
	// <li>不填或填0，表示水印持续到画面结束；</li>
	// <li>当数值大于0时（假设为 n），表示水印持续到第 n 秒时消失；</li>
	// <li>当数值小于0时（假设为 -n），表示水印持续到离画面结束 n 秒前消失。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`
}

type WorkflowTaskForDY struct {
	// 视频处理任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 已弃用，请使用各个具体任务的 ErrCode。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 已弃用，请使用各个具体任务的 Message。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 视频处理的目标文件信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputInfo *MediaInputInfoForDY `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`

	// 原始视频的元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaDataForDY `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// 视频处理任务的执行状态与结果。
	MediaProcessResultSet []*MediaProcessTaskResultForDY `json:"MediaProcessResultSet,omitnil,omitempty" name:"MediaProcessResultSet"`
}