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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AdaptiveDynamicStreamingTaskInputForDY struct {

	// 转自适应码流模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInputForDY `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// 转自适应码流后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 转自适应码流后，manifest 文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 转自适应码流后，子流文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}.{format}`。
	SubStreamObjectName *string `json:"SubStreamObjectName,omitempty" name:"SubStreamObjectName"`

	// 转自适应码流（仅 HLS）后，分片文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}_{segmentNumber}.{format}`。
	SegmentObjectName *string `json:"SegmentObjectName,omitempty" name:"SegmentObjectName"`
}

type AnimatedGraphicTaskInputForDY struct {

	// 视频转动图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 动图在视频中的开始时间，单位为秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间，单位为秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 转动图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 转动图后文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_animatedGraphic_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`
}

type CosInputInfoForDY struct {

	// 视频处理对象文件所在的 COS Bucket 名，如 TopRankVideo-125xxx88。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 视频处理对象文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 视频处理对象文件的输入路径，如`/movie/201907/WildAnimal.mov`。
	Object *string `json:"Object,omitempty" name:"Object"`
}

type CosOutputStorageForDY struct {

	// 视频处理生成的文件输出的目标 Bucket 名，如 TopRankVideo-125xxx88。如果不填，表示继承上层。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 视频处理生成的文件输出的目标 Bucket 的园区，如 ap-chongqing。如果不填，表示继承上层。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DeleteMediaForDYRequest struct {
	*tchttp.BaseRequest

	// 文件所在的 COS Bucket 名，如 wsd-tx-ugc-pub。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// 文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 文件的 COS 完整路径。
	ObjectSet []*string `json:"ObjectSet,omitempty" name:"ObjectSet"`

	// 来源上下文，用于透传用户请求信息，删除回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
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

type DeleteMediaForDYResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type EditMediaFileInfoForDY struct {

	// 视频的输入信息。
	InputInfo *MediaInputInfoForDY `json:"InputInfo,omitempty" name:"InputInfo"`

	// 视频剪辑的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 视频剪辑的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type EditMediaForDYRequest struct {
	*tchttp.BaseRequest

	// 输入的视频文件信息。
	FileInfos []*EditMediaFileInfoForDY `json:"FileInfos,omitempty" name:"FileInfos"`

	// 视频处理输出文件的目标存储。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 视频处理输出文件的目标路径。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 编辑后生成的文件配置。
	OutputConfig *EditMediaOutputConfigForDY `json:"OutputConfig,omitempty" name:"OutputConfig"`

	// 任务优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
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
	delete(f, "SessionId")
	delete(f, "SessionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditMediaForDYRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EditMediaForDYResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 编辑视频的任务 ID，可以通过该 ID 查询编辑任务的状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Container *string `json:"Container,omitempty" name:"Container"`
}

type ImageSpriteTaskInputForDY struct {

	// 雪碧图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 截取雪碧图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 截取雪碧图后，雪碧图图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_imageSprite_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 截取雪碧图后，Web VTT 文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：`{inputName}_imageSprite_{definition}.{format}`。
	WebVttObjectName *string `json:"WebVttObjectName,omitempty" name:"WebVttObjectName"`

	// 截取雪碧图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormatForDY `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type MediaInputInfoForDY struct {

	// 输入来源对象的类型，支持 COS 和 URL 两种。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当 Type 为 COS 时有效，则该项为必填，表示视频处理 COS 对象信息。
	CosInputInfo *CosInputInfoForDY `json:"CosInputInfo,omitempty" name:"CosInputInfo"`
}

type MediaProcessTaskInputForDY struct {

	// 视频转码任务列表。
	TranscodeTaskSet []*TranscodeTaskInputForDY `json:"TranscodeTaskSet,omitempty" name:"TranscodeTaskSet"`

	// 视频转动图任务列表。
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInputForDY `json:"AnimatedGraphicTaskSet,omitempty" name:"AnimatedGraphicTaskSet"`

	// 对视频按时间点截图任务列表。
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInputForDY `json:"SnapshotByTimeOffsetTaskSet,omitempty" name:"SnapshotByTimeOffsetTaskSet"`

	// 对视频采样截图任务列表。
	SampleSnapshotTaskSet []*SampleSnapshotTaskInputForDY `json:"SampleSnapshotTaskSet,omitempty" name:"SampleSnapshotTaskSet"`

	// 对视频截雪碧图任务列表。
	ImageSpriteTaskSet []*ImageSpriteTaskInputForDY `json:"ImageSpriteTaskSet,omitempty" name:"ImageSpriteTaskSet"`

	// 转自适应码流任务列表。
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInputForDY `json:"AdaptiveDynamicStreamingTaskSet,omitempty" name:"AdaptiveDynamicStreamingTaskSet"`
}

type NumberFormatForDY struct {

	// `{number}`变量的起始值，默认为0。
	InitialValue *uint64 `json:"InitialValue,omitempty" name:"InitialValue"`

	// `{number}`变量的增长步长，默认为1。
	Increment *uint64 `json:"Increment,omitempty" name:"Increment"`

	// `{number}`变量的最小长度，不足时补占位符。默认为1。
	MinLength *uint64 `json:"MinLength,omitempty" name:"MinLength"`

	// `{number}`变量的长度不足时，补充的占位符。默认为"0"。
	PlaceHolder *string `json:"PlaceHolder,omitempty" name:"PlaceHolder"`
}

type ProcessMediaForDYRequest struct {
	*tchttp.BaseRequest

	// 视频处理的文件输入信息。
	InputInfo *MediaInputInfoForDY `json:"InputInfo,omitempty" name:"InputInfo"`

	// 视频处理输出文件的目标存储。
	// 当 InputInfo 的输入类型是 COS 时可以不填，不填代表继承 InputInfo 中的存储位置。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 视频处理生成的文件输出的目标目录，如`/movie/201907/`。如果不填，表示与 InputInfo 中文件所在的目录一致。
	OutputDir *string `json:"OutputDir,omitempty" name:"OutputDir"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInputForDY `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到 10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`
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

type ProcessMediaForDYResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	ImageContent *MediaInputInfoForDY `json:"ImageContent,omitempty" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：10%。
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素。</li>
	// 默认值：0px，表示 Height 按照原始水印图片的宽高比缩放。
	Height *string `json:"Height,omitempty" name:"Height"`

	// 印重复类型。使用场景：水印为动态图像。取值范围：
	// <li>once：动态水印播放完后，不再出现；</li>
	// <li>repeat_last_frame：水印播放完后，停留在最后一帧；</li>
	// <li>repeat：水印循环播放，直到视频结束（默认值）。</li>
	RepeatType *string `json:"RepeatType,omitempty" name:"RepeatType"`
}

type RawWatermarkParameterForDY struct {

	// 水印类型，可选值：
	// <li>image：图片水印。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 原点位置，目前仅支持：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角。</li>
	// 默认值：TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 图片水印模板，当 Type 为 image，该字段必填。当 Type 为 text，该字段无效。
	ImageTemplate *RawImageWatermarkInputForDY `json:"ImageTemplate,omitempty" name:"ImageTemplate"`
}

type RestoreMediaForDYRequest struct {
	*tchttp.BaseRequest

	// 文件所在的 COS Bucket 名，如 bucket-xxx。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// 文件所在的 COS Bucket 所属园区，如 ap-chongqing。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 文件的 COS 完整路径。
	Object *string `json:"Object,omitempty" name:"Object"`

	// 解冻出的临时媒体文件的可访问持续时长，单位为“天”。默认为1天。
	RestoreDay *uint64 `json:"RestoreDay,omitempty" name:"RestoreDay"`

	// 来源上下文，用于透传用户请求信息，最长 1000 个字符。
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
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

type RestoreMediaForDYResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInputForDY `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// 采样截图后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 采样截图后图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_sampleSnapshot_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 采样截图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormatForDY `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type SnapshotByTimeOffsetTaskInputForDY struct {

	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 截图时间点列表，时间点支持 s、% 两种格式：
	// <li>当字符串以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// <li>当字符串以 % 结尾，表示时间点为视频时长的百分比大小，如10%表示时间点为视频前第10%的时间。</li>
	ExtTimeOffsetSet []*string `json:"ExtTimeOffsetSet,omitempty" name:"ExtTimeOffsetSet"`

	// 截图时间点列表，单位为<font color=red>秒</font>。此参数已不再建议使用，建议您使用 ExtTimeOffsetSet 参数。
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitempty" name:"TimeOffsetSet"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInputForDY `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// 时间点截图后文件的目标存储，不填则继承上层的 OutputStorage 值。<li>当字符串以 s 结尾，表示时间点单位为秒，如 3.5s 表示时间点为第3.5秒；</li>
	// <li>当字符串以 % 结尾，表示时间点为视频时长的百分比大小，如10%表示时间点为视频前第10%的时间。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 时间点截图后图片文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_snapshotByTimeOffset_{definition}_{number}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 时间点截图后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormatForDY `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type TaskOutputStorageForDY struct {

	// 视频处理输出对象存储位置的类型，现在仅支持 COS。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 当 Type 为 COS 时有效，则该项为必填，表示视频处理 COS 输出位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosOutputStorage *CosOutputStorageForDY `json:"CosOutputStorage,omitempty" name:"CosOutputStorage"`
}

type TranscodeTaskInputForDY struct {

	// 视频转码模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInputForDY `json:"WatermarkSet,omitempty" name:"WatermarkSet"`

	// 转码后的视频的起始时间偏移，单位：秒。
	// <li>不填或填0，表示转码后的视频从原始视频的起始位置开始；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频从原始视频的第 n 秒位置开始；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频从原始视频结束 n 秒前的位置开始。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 转码后视频的终止时间偏移，单位：秒。
	// <li>不填或填0，表示转码后的视频持续到原始视频的末尾终止；</li>
	// <li>当数值大于0时（假设为 n），表示转码后的视频持续到原始视频第 n 秒时终止；</li>
	// <li>当数值小于0时（假设为 -n），表示转码后的视频持续到原始视频结束 n 秒前终止。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 转码后文件的目标存储，不填则继承上层的 OutputStorage 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputStorage *TaskOutputStorageForDY `json:"OutputStorage,omitempty" name:"OutputStorage"`

	// 转码后主文件的输出路径，可以为相对路径或者绝对路径。如果不填，则默认为相对路径：`{inputName}_transcode_{definition}.{format}`。
	OutputObjectPath *string `json:"OutputObjectPath,omitempty" name:"OutputObjectPath"`

	// 码后分片文件的输出路径（转码 HLS 时 ts 的路径），只能为相对路径。如果不填，则默认为：`{inputName}_transcode_{definition}_{number}.{format}`。
	SegmentObjectName *string `json:"SegmentObjectName,omitempty" name:"SegmentObjectName"`

	// 转码后输出路径中的`{number}`变量的规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectNumberFormat *NumberFormatForDY `json:"ObjectNumberFormat,omitempty" name:"ObjectNumberFormat"`
}

type WatermarkInputForDY struct {

	// 水印模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印自定义参数，当 Definition 填 0 时有效。
	// 该参数用于高度定制场景，建议您优先使用 Definition 指定水印参数。
	// 水印自定义参数不支持截图打水印。
	RawParameter *RawWatermarkParameterForDY `json:"RawParameter,omitempty" name:"RawParameter"`

	// 文字内容，长度不超过100个字符。仅当水印类型为文字水印时填写。
	// 文字水印不支持截图打水印。
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`

	// SVG 内容。长度不超过 2000000 个字符。仅当水印类型为 SVG 水印时填写。
	// SVG 水印不支持截图打水印。
	SvgContent *string `json:"SvgContent,omitempty" name:"SvgContent"`

	// 水印的起始时间偏移，单位：秒。不填或填0，表示水印从画面出现时开始显现。
	// <li>不填或填0，表示水印从画面开始就出现；</li>
	// <li>当数值大于0时（假设为 n），表示水印从画面开始的第 n 秒出现；</li>
	// <li>当数值小于0时（假设为 -n），表示水印从离画面结束 n 秒前开始出现。</li>
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 水印的结束时间偏移，单位：秒。
	// <li>不填或填0，表示水印持续到画面结束；</li>
	// <li>当数值大于0时（假设为 n），表示水印持续到第 n 秒时消失；</li>
	// <li>当数值小于0时（假设为 -n），表示水印持续到离画面结束 n 秒前消失。</li>
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}
