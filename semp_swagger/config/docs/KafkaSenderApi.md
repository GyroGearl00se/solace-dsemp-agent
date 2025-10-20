# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnKafkaSender**](KafkaSenderApi.md#CreateMsgVpnKafkaSender) | **Post** /msgVpns/{msgVpnName}/kafkaSenders | Create a Kafka Sender object.
[**CreateMsgVpnKafkaSenderQueueBinding**](KafkaSenderApi.md#CreateMsgVpnKafkaSenderQueueBinding) | **Post** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName}/queueBindings | Create a Queue Binding object.
[**DeleteMsgVpnKafkaSender**](KafkaSenderApi.md#DeleteMsgVpnKafkaSender) | **Delete** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName} | Delete a Kafka Sender object.
[**DeleteMsgVpnKafkaSenderQueueBinding**](KafkaSenderApi.md#DeleteMsgVpnKafkaSenderQueueBinding) | **Delete** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName}/queueBindings/{queueName} | Delete a Queue Binding object.
[**GetMsgVpnKafkaSender**](KafkaSenderApi.md#GetMsgVpnKafkaSender) | **Get** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName} | Get a Kafka Sender object.
[**GetMsgVpnKafkaSenderQueueBinding**](KafkaSenderApi.md#GetMsgVpnKafkaSenderQueueBinding) | **Get** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName}/queueBindings/{queueName} | Get a Queue Binding object.
[**GetMsgVpnKafkaSenderQueueBindings**](KafkaSenderApi.md#GetMsgVpnKafkaSenderQueueBindings) | **Get** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName}/queueBindings | Get a list of Queue Binding objects.
[**GetMsgVpnKafkaSenders**](KafkaSenderApi.md#GetMsgVpnKafkaSenders) | **Get** /msgVpns/{msgVpnName}/kafkaSenders | Get a list of Kafka Sender objects.
[**ReplaceMsgVpnKafkaSender**](KafkaSenderApi.md#ReplaceMsgVpnKafkaSender) | **Put** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName} | Replace a Kafka Sender object.
[**ReplaceMsgVpnKafkaSenderQueueBinding**](KafkaSenderApi.md#ReplaceMsgVpnKafkaSenderQueueBinding) | **Put** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName}/queueBindings/{queueName} | Replace a Queue Binding object.
[**UpdateMsgVpnKafkaSender**](KafkaSenderApi.md#UpdateMsgVpnKafkaSender) | **Patch** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName} | Update a Kafka Sender object.
[**UpdateMsgVpnKafkaSenderQueueBinding**](KafkaSenderApi.md#UpdateMsgVpnKafkaSenderQueueBinding) | **Patch** /msgVpns/{msgVpnName}/kafkaSenders/{kafkaSenderName}/queueBindings/{queueName} | Update a Queue Binding object.

# **CreateMsgVpnKafkaSender**
> MsgVpnKafkaSenderResponse CreateMsgVpnKafkaSender(ctx, body, msgVpnName, optional)
Create a Kafka Sender object.

Create a Kafka Sender object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Kafka Sender sends messages to a Kafka Cluster.   Attribute|Identifying|Const|Required|Read-Only|Write-Only|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationAwsMskIamSecretAccessKey|||||x| authenticationBasicPassword|||||x|x authenticationClientCertContent|||||x|x authenticationClientCertPassword|||||x| authenticationKerberosKeytabContent|||||x| authenticationOauthClientSecret|||||x|x authenticationScramPassword|||||x|x kafkaSenderName|x|x|x||| msgVpnName|x|||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires :---|:---|:--- MsgVpnKafkaSender|authenticationBasicPassword|authenticationBasicUsername MsgVpnKafkaSender|authenticationClientCertPassword|authenticationClientCertContent MsgVpnKafkaSender|authenticationKerberosKeytabContent|authenticationKerberosKeytabFileName, authenticationKerberosUserPrincipalName MsgVpnKafkaSender|authenticationKerberosKeytabFileName|authenticationKerberosKeytabContent, authenticationKerberosUserPrincipalName MsgVpnKafkaSender|authenticationKerberosUserPrincipalName|authenticationKerberosKeytabContent, authenticationKerberosKeytabFileName MsgVpnKafkaSender|authenticationScramPassword|authenticationScramUsername    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaSender**](MsgVpnKafkaSender.md)| The Kafka Sender object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***KafkaSenderApiCreateMsgVpnKafkaSenderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiCreateMsgVpnKafkaSenderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderResponse**](MsgVpnKafkaSenderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnKafkaSenderQueueBinding**
> MsgVpnKafkaSenderQueueBindingResponse CreateMsgVpnKafkaSenderQueueBinding(ctx, body, msgVpnName, kafkaSenderName, optional)
Create a Queue Binding object.

Create a Queue Binding object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue Binding sends messages from a local Solace Queue to a remote Kafka topic.   Attribute|Identifying|Const|Required|Read-Only :---|:---:|:---:|:---:|:---: kafkaSenderName|x|||x msgVpnName|x|||x queueName|x|x|x|    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaSenderQueueBinding**](MsgVpnKafkaSenderQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
 **optional** | ***KafkaSenderApiCreateMsgVpnKafkaSenderQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiCreateMsgVpnKafkaSenderQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderQueueBindingResponse**](MsgVpnKafkaSenderQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnKafkaSender**
> SempMetaOnlyResponse DeleteMsgVpnKafkaSender(ctx, msgVpnName, kafkaSenderName)
Delete a Kafka Sender object.

Delete a Kafka Sender object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Kafka Sender sends messages to a Kafka Cluster.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnKafkaSenderQueueBinding**
> SempMetaOnlyResponse DeleteMsgVpnKafkaSenderQueueBinding(ctx, msgVpnName, kafkaSenderName, queueName)
Delete a Queue Binding object.

Delete a Queue Binding object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue Binding sends messages from a local Solace Queue to a remote Kafka topic.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnKafkaSender**
> MsgVpnKafkaSenderResponse GetMsgVpnKafkaSender(ctx, msgVpnName, kafkaSenderName, optional)
Get a Kafka Sender object.

Get a Kafka Sender object.  A Kafka Sender sends messages to a Kafka Cluster.   Attribute|Identifying|Write-Only|Opaque :---|:---:|:---:|:---: authenticationAwsMskIamSecretAccessKey||x| authenticationBasicPassword||x|x authenticationClientCertContent||x|x authenticationClientCertPassword||x| authenticationKerberosKeytabContent||x| authenticationOauthClientSecret||x|x authenticationScramPassword||x|x kafkaSenderName|x|| msgVpnName|x||    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
 **optional** | ***KafkaSenderApiGetMsgVpnKafkaSenderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiGetMsgVpnKafkaSenderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderResponse**](MsgVpnKafkaSenderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnKafkaSenderQueueBinding**
> MsgVpnKafkaSenderQueueBindingResponse GetMsgVpnKafkaSenderQueueBinding(ctx, msgVpnName, kafkaSenderName, queueName, optional)
Get a Queue Binding object.

Get a Queue Binding object.  A Queue Binding sends messages from a local Solace Queue to a remote Kafka topic.   Attribute|Identifying :---|:---: kafkaSenderName|x msgVpnName|x queueName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***KafkaSenderApiGetMsgVpnKafkaSenderQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiGetMsgVpnKafkaSenderQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderQueueBindingResponse**](MsgVpnKafkaSenderQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnKafkaSenderQueueBindings**
> MsgVpnKafkaSenderQueueBindingsResponse GetMsgVpnKafkaSenderQueueBindings(ctx, msgVpnName, kafkaSenderName, optional)
Get a list of Queue Binding objects.

Get a list of Queue Binding objects.  A Queue Binding sends messages from a local Solace Queue to a remote Kafka topic.   Attribute|Identifying :---|:---: kafkaSenderName|x msgVpnName|x queueName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 100.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
 **optional** | ***KafkaSenderApiGetMsgVpnKafkaSenderQueueBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiGetMsgVpnKafkaSenderQueueBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderQueueBindingsResponse**](MsgVpnKafkaSenderQueueBindingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnKafkaSenders**
> MsgVpnKafkaSendersResponse GetMsgVpnKafkaSenders(ctx, msgVpnName, optional)
Get a list of Kafka Sender objects.

Get a list of Kafka Sender objects.  A Kafka Sender sends messages to a Kafka Cluster.   Attribute|Identifying|Write-Only|Opaque :---|:---:|:---:|:---: authenticationAwsMskIamSecretAccessKey||x| authenticationBasicPassword||x|x authenticationClientCertContent||x|x authenticationClientCertPassword||x| authenticationKerberosKeytabContent||x| authenticationOauthClientSecret||x|x authenticationScramPassword||x|x kafkaSenderName|x|| msgVpnName|x||    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 100.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***KafkaSenderApiGetMsgVpnKafkaSendersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiGetMsgVpnKafkaSendersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSendersResponse**](MsgVpnKafkaSendersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnKafkaSender**
> MsgVpnKafkaSenderResponse ReplaceMsgVpnKafkaSender(ctx, body, msgVpnName, kafkaSenderName, optional)
Replace a Kafka Sender object.

Replace a Kafka Sender object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  A Kafka Sender sends messages to a Kafka Cluster.   Attribute|Identifying|Const|Read-Only|Write-Only|Auto-Disable|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationAwsMskIamAccessKeyId|||||x| authenticationAwsMskIamRegion|||||x| authenticationAwsMskIamSecretAccessKey||||x|x| authenticationAwsMskIamStsExternalId|||||x| authenticationAwsMskIamStsRoleArn|||||x| authenticationAwsMskIamStsRoleSessionName|||||x| authenticationBasicPassword||||x|x|x authenticationBasicUsername|||||x| authenticationClientCertContent||||x|x|x authenticationClientCertPassword||||x|x| authenticationKerberosKeytabContent||||x|x| authenticationKerberosKeytabFileName|||||x| authenticationKerberosServiceName|||||x| authenticationKerberosUserPrincipalName|||||x| authenticationOauthClientId|||||x| authenticationOauthClientScope|||||x| authenticationOauthClientSecret||||x|x|x authenticationOauthClientTokenEndpoint|||||x| authenticationScheme|||||x| authenticationScramHash|||||x| authenticationScramPassword||||x|x|x authenticationScramUsername|||||x| batchDelay|||||x| batchMaxMsgCount|||||x| batchMaxSize|||||x| bootstrapAddressList|||||x| idempotenceEnabled|||||x| kafkaSenderName|x|x|||| msgVpnName|x||x||| transportCompressionEnabled|||||x| transportCompressionLevel|||||x| transportCompressionType|||||x| transportTlsEnabled|||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires :---|:---|:--- MsgVpnKafkaSender|authenticationBasicPassword|authenticationBasicUsername MsgVpnKafkaSender|authenticationClientCertPassword|authenticationClientCertContent MsgVpnKafkaSender|authenticationKerberosKeytabContent|authenticationKerberosKeytabFileName, authenticationKerberosUserPrincipalName MsgVpnKafkaSender|authenticationKerberosKeytabFileName|authenticationKerberosKeytabContent, authenticationKerberosUserPrincipalName MsgVpnKafkaSender|authenticationKerberosUserPrincipalName|authenticationKerberosKeytabContent, authenticationKerberosKeytabFileName MsgVpnKafkaSender|authenticationScramPassword|authenticationScramUsername    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"vpn/read-write\" is required if this operation creates an object.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaSender**](MsgVpnKafkaSender.md)| The Kafka Sender object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
 **optional** | ***KafkaSenderApiReplaceMsgVpnKafkaSenderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiReplaceMsgVpnKafkaSenderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderResponse**](MsgVpnKafkaSenderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnKafkaSenderQueueBinding**
> MsgVpnKafkaSenderQueueBindingResponse ReplaceMsgVpnKafkaSenderQueueBinding(ctx, body, msgVpnName, kafkaSenderName, queueName, optional)
Replace a Queue Binding object.

Replace a Queue Binding object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  A Queue Binding sends messages from a local Solace Queue to a remote Kafka topic.   Attribute|Identifying|Const|Read-Only|Auto-Disable :---|:---:|:---:|:---:|:---: ackMode||||x kafkaSenderName|x||x| msgVpnName|x||x| partitionConsistentHash||||x partitionExplicitNumber||||x partitionRandomFallbackEnabled||||x partitionScheme||||x queueName|x|x|| remoteKey||||x remoteTopic||||x    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"vpn/read-write\" is required if this operation creates an object.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaSenderQueueBinding**](MsgVpnKafkaSenderQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***KafkaSenderApiReplaceMsgVpnKafkaSenderQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiReplaceMsgVpnKafkaSenderQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderQueueBindingResponse**](MsgVpnKafkaSenderQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnKafkaSender**
> MsgVpnKafkaSenderResponse UpdateMsgVpnKafkaSender(ctx, body, msgVpnName, kafkaSenderName, optional)
Update a Kafka Sender object.

Update a Kafka Sender object. Any attribute missing from the request will be left unchanged.  A Kafka Sender sends messages to a Kafka Cluster.   Attribute|Identifying|Const|Read-Only|Write-Only|Auto-Disable|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationAwsMskIamAccessKeyId|||||x| authenticationAwsMskIamRegion|||||x| authenticationAwsMskIamSecretAccessKey||||x|x| authenticationAwsMskIamStsExternalId|||||x| authenticationAwsMskIamStsRoleArn|||||x| authenticationAwsMskIamStsRoleSessionName|||||x| authenticationBasicPassword||||x|x|x authenticationBasicUsername|||||x| authenticationClientCertContent||||x|x|x authenticationClientCertPassword||||x|x| authenticationKerberosKeytabContent||||x|x| authenticationKerberosKeytabFileName|||||x| authenticationKerberosServiceName|||||x| authenticationKerberosUserPrincipalName|||||x| authenticationOauthClientId|||||x| authenticationOauthClientScope|||||x| authenticationOauthClientSecret||||x|x|x authenticationOauthClientTokenEndpoint|||||x| authenticationScheme|||||x| authenticationScramHash|||||x| authenticationScramPassword||||x|x|x authenticationScramUsername|||||x| batchDelay|||||x| batchMaxMsgCount|||||x| batchMaxSize|||||x| bootstrapAddressList|||||x| idempotenceEnabled|||||x| kafkaSenderName|x|x|||| msgVpnName|x||x||| transportCompressionEnabled|||||x| transportCompressionLevel|||||x| transportCompressionType|||||x| transportTlsEnabled|||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires :---|:---|:--- MsgVpnKafkaSender|authenticationBasicPassword|authenticationBasicUsername MsgVpnKafkaSender|authenticationClientCertPassword|authenticationClientCertContent MsgVpnKafkaSender|authenticationKerberosKeytabContent|authenticationKerberosKeytabFileName, authenticationKerberosUserPrincipalName MsgVpnKafkaSender|authenticationKerberosKeytabFileName|authenticationKerberosKeytabContent, authenticationKerberosUserPrincipalName MsgVpnKafkaSender|authenticationKerberosUserPrincipalName|authenticationKerberosKeytabContent, authenticationKerberosKeytabFileName MsgVpnKafkaSender|authenticationScramPassword|authenticationScramUsername    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaSender**](MsgVpnKafkaSender.md)| The Kafka Sender object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
 **optional** | ***KafkaSenderApiUpdateMsgVpnKafkaSenderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiUpdateMsgVpnKafkaSenderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderResponse**](MsgVpnKafkaSenderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnKafkaSenderQueueBinding**
> MsgVpnKafkaSenderQueueBindingResponse UpdateMsgVpnKafkaSenderQueueBinding(ctx, body, msgVpnName, kafkaSenderName, queueName, optional)
Update a Queue Binding object.

Update a Queue Binding object. Any attribute missing from the request will be left unchanged.  A Queue Binding sends messages from a local Solace Queue to a remote Kafka topic.   Attribute|Identifying|Const|Read-Only|Auto-Disable :---|:---:|:---:|:---:|:---: ackMode||||x kafkaSenderName|x||x| msgVpnName|x||x| partitionConsistentHash||||x partitionExplicitNumber||||x partitionRandomFallbackEnabled||||x partitionScheme||||x queueName|x|x|| remoteKey||||x remoteTopic||||x    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaSenderQueueBinding**](MsgVpnKafkaSenderQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaSenderName** | **string**| The name of the Kafka Sender. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***KafkaSenderApiUpdateMsgVpnKafkaSenderQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaSenderApiUpdateMsgVpnKafkaSenderQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaSenderQueueBindingResponse**](MsgVpnKafkaSenderQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

