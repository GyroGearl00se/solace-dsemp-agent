# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#CreateMsgVpnTopicEndpointTemplate) | **Post** /msgVpns/{msgVpnName}/topicEndpointTemplates | Create a Topic Endpoint Template object.
[**DeleteMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#DeleteMsgVpnTopicEndpointTemplate) | **Delete** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Delete a Topic Endpoint Template object.
[**GetMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#GetMsgVpnTopicEndpointTemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
[**GetMsgVpnTopicEndpointTemplates**](TopicEndpointTemplateApi.md#GetMsgVpnTopicEndpointTemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
[**ReplaceMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#ReplaceMsgVpnTopicEndpointTemplate) | **Put** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Replace a Topic Endpoint Template object.
[**UpdateMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#UpdateMsgVpnTopicEndpointTemplate) | **Patch** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Update a Topic Endpoint Template object.

# **CreateMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse CreateMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, optional)
Create a Topic Endpoint Template object.

Create a Topic Endpoint Template object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Const|Required|Read-Only :---|:---:|:---:|:---:|:---: msgVpnName|x|||x topicEndpointTemplateName|x|x|x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnTopicEndpointTemplateEventBindCountThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventBindCountThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventBindCountThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventBindCountThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|setValue|clearValue|clearPercent, setPercent    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***TopicEndpointTemplateApiCreateMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiCreateMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnTopicEndpointTemplate**
> SempMetaOnlyResponse DeleteMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName)
Delete a Topic Endpoint Template object.

Delete a Topic Endpoint Template object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse GetMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName, optional)
Get a Topic Endpoint Template object.

Get a Topic Endpoint Template object.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying :---|:---: msgVpnName|x topicEndpointTemplateName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTemplates**
> MsgVpnTopicEndpointTemplatesResponse GetMsgVpnTopicEndpointTemplates(ctx, msgVpnName, optional)
Get a list of Topic Endpoint Template objects.

Get a list of Topic Endpoint Template objects.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying :---|:---: msgVpnName|x topicEndpointTemplateName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 500.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplatesResponse**](MsgVpnTopicEndpointTemplatesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse ReplaceMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, topicEndpointTemplateName, optional)
Replace a Topic Endpoint Template object.

Replace a Topic Endpoint Template object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Const|Read-Only :---|:---:|:---:|:---: msgVpnName|x||x topicEndpointTemplateName|x|x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnTopicEndpointTemplateEventBindCountThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventBindCountThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventBindCountThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventBindCountThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|setValue|clearValue|clearPercent, setPercent    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"vpn/read-write\" is required if this operation creates an object.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***TopicEndpointTemplateApiReplaceMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiReplaceMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse UpdateMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, topicEndpointTemplateName, optional)
Update a Topic Endpoint Template object.

Update a Topic Endpoint Template object. Any attribute missing from the request will be left unchanged.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Const|Read-Only :---|:---:|:---:|:---: msgVpnName|x||x topicEndpointTemplateName|x|x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnTopicEndpointTemplateEventBindCountThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventBindCountThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventBindCountThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventBindCountThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventMsgSpoolUsageThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnTopicEndpointTemplateEventRejectLowPriorityMsgLimitThreshold|setValue|clearValue|clearPercent, setPercent    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***TopicEndpointTemplateApiUpdateMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiUpdateMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

