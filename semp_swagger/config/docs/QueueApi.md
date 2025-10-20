# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnQueue**](QueueApi.md#CreateMsgVpnQueue) | **Post** /msgVpns/{msgVpnName}/queues | Create a Queue object.
[**CreateMsgVpnQueueSubscription**](QueueApi.md#CreateMsgVpnQueueSubscription) | **Post** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Create a Queue Subscription object.
[**DeleteMsgVpnQueue**](QueueApi.md#DeleteMsgVpnQueue) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName} | Delete a Queue object.
[**DeleteMsgVpnQueueSubscription**](QueueApi.md#DeleteMsgVpnQueueSubscription) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Delete a Queue Subscription object.
[**GetMsgVpnQueue**](QueueApi.md#GetMsgVpnQueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
[**GetMsgVpnQueueSubscription**](QueueApi.md#GetMsgVpnQueueSubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Get a Queue Subscription object.
[**GetMsgVpnQueueSubscriptions**](QueueApi.md#GetMsgVpnQueueSubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Get a list of Queue Subscription objects.
[**GetMsgVpnQueues**](QueueApi.md#GetMsgVpnQueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
[**ReplaceMsgVpnQueue**](QueueApi.md#ReplaceMsgVpnQueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName} | Replace a Queue object.
[**UpdateMsgVpnQueue**](QueueApi.md#UpdateMsgVpnQueue) | **Patch** /msgVpns/{msgVpnName}/queues/{queueName} | Update a Queue object.

# **CreateMsgVpnQueue**
> MsgVpnQueueResponse CreateMsgVpnQueue(ctx, body, msgVpnName, optional)
Create a Queue object.

Create a Queue object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Const|Required|Read-Only :---|:---:|:---:|:---:|:---: msgVpnName|x|||x queueName|x|x|x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnQueueEventBindCountThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventBindCountThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventBindCountThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventBindCountThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnQueueEventMsgSpoolUsageThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventMsgSpoolUsageThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventMsgSpoolUsageThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventMsgSpoolUsageThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|setValue|clearValue|clearPercent, setPercent    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueue**](MsgVpnQueue.md)| The Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***QueueApiCreateMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiCreateMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnQueueSubscription**
> MsgVpnQueueSubscriptionResponse CreateMsgVpnQueueSubscription(ctx, body, msgVpnName, queueName, optional)
Create a Queue Subscription object.

Create a Queue Subscription object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.   Attribute|Identifying|Const|Required|Read-Only :---|:---:|:---:|:---:|:---: msgVpnName|x|||x queueName|x|||x subscriptionTopic|x|x|x|    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueSubscription**](MsgVpnQueueSubscription.md)| The Queue Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***QueueApiCreateMsgVpnQueueSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiCreateMsgVpnQueueSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionResponse**](MsgVpnQueueSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnQueue**
> SempMetaOnlyResponse DeleteMsgVpnQueue(ctx, msgVpnName, queueName)
Delete a Queue object.

Delete a Queue object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnQueueSubscription**
> SempMetaOnlyResponse DeleteMsgVpnQueueSubscription(ctx, msgVpnName, queueName, subscriptionTopic)
Delete a Queue Subscription object.

Delete a Queue Subscription object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **subscriptionTopic** | **string**| The topic of the Subscription. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueue**
> MsgVpnQueueResponse GetMsgVpnQueue(ctx, msgVpnName, queueName, optional)
Get a Queue object.

Get a Queue object.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying :---|:---: msgVpnName|x queueName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***QueueApiGetMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiGetMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueSubscription**
> MsgVpnQueueSubscriptionResponse GetMsgVpnQueueSubscription(ctx, msgVpnName, queueName, subscriptionTopic, optional)
Get a Queue Subscription object.

Get a Queue Subscription object.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.   Attribute|Identifying :---|:---: msgVpnName|x queueName|x subscriptionTopic|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **subscriptionTopic** | **string**| The topic of the Subscription. | 
 **optional** | ***QueueApiGetMsgVpnQueueSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiGetMsgVpnQueueSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionResponse**](MsgVpnQueueSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueSubscriptions**
> MsgVpnQueueSubscriptionsResponse GetMsgVpnQueueSubscriptions(ctx, msgVpnName, queueName, optional)
Get a list of Queue Subscription objects.

Get a list of Queue Subscription objects.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.   Attribute|Identifying :---|:---: msgVpnName|x queueName|x subscriptionTopic|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 100.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***QueueApiGetMsgVpnQueueSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiGetMsgVpnQueueSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionsResponse**](MsgVpnQueueSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueues**
> MsgVpnQueuesResponse GetMsgVpnQueues(ctx, msgVpnName, optional)
Get a list of Queue objects.

Get a list of Queue objects.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying :---|:---: msgVpnName|x queueName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 100.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***QueueApiGetMsgVpnQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiGetMsgVpnQueuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueuesResponse**](MsgVpnQueuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnQueue**
> MsgVpnQueueResponse ReplaceMsgVpnQueue(ctx, body, msgVpnName, queueName, optional)
Replace a Queue object.

Replace a Queue object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Const|Read-Only|Auto-Disable :---|:---:|:---:|:---:|:---: accessType||||x msgVpnName|x||x| owner||||x permission||||x queueName|x|x|| redeliveryDelayEnabled||||x redeliveryDelayInitialInterval||||x redeliveryDelayMaxInterval||||x redeliveryDelayMultiplier||||x rejectMsgToSenderOnDiscardBehavior||||x respectMsgPriorityEnabled||||x    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnQueueEventBindCountThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventBindCountThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventBindCountThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventBindCountThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnQueueEventMsgSpoolUsageThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventMsgSpoolUsageThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventMsgSpoolUsageThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventMsgSpoolUsageThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|setValue|clearValue|clearPercent, setPercent    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"vpn/read-write\" is required if this operation creates an object.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueue**](MsgVpnQueue.md)| The Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***QueueApiReplaceMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiReplaceMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnQueue**
> MsgVpnQueueResponse UpdateMsgVpnQueue(ctx, body, msgVpnName, queueName, optional)
Update a Queue object.

Update a Queue object. Any attribute missing from the request will be left unchanged.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Const|Read-Only|Auto-Disable :---|:---:|:---:|:---:|:---: accessType||||x msgVpnName|x||x| owner||||x permission||||x queueName|x|x|| redeliveryDelayEnabled||||x redeliveryDelayInitialInterval||||x redeliveryDelayMaxInterval||||x redeliveryDelayMultiplier||||x rejectMsgToSenderOnDiscardBehavior||||x respectMsgPriorityEnabled||||x    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnQueueEventBindCountThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventBindCountThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventBindCountThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventBindCountThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnQueueEventMsgSpoolUsageThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventMsgSpoolUsageThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventMsgSpoolUsageThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventMsgSpoolUsageThreshold|setValue|clearValue|clearPercent, setPercent MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|clearPercent|setPercent|clearValue, setValue MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|clearValue|setValue|clearPercent, setPercent MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|setPercent|clearPercent|clearValue, setValue MsgVpnQueueEventRejectLowPriorityMsgLimitThreshold|setValue|clearValue|clearPercent, setPercent    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueue**](MsgVpnQueue.md)| The Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***QueueApiUpdateMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiUpdateMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

