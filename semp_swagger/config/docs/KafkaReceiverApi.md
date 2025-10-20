# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnKafkaReceiver**](KafkaReceiverApi.md#CreateMsgVpnKafkaReceiver) | **Post** /msgVpns/{msgVpnName}/kafkaReceivers | Create a Kafka Receiver object.
[**CreateMsgVpnKafkaReceiverTopicBinding**](KafkaReceiverApi.md#CreateMsgVpnKafkaReceiverTopicBinding) | **Post** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName}/topicBindings | Create a Topic Binding object.
[**DeleteMsgVpnKafkaReceiver**](KafkaReceiverApi.md#DeleteMsgVpnKafkaReceiver) | **Delete** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName} | Delete a Kafka Receiver object.
[**DeleteMsgVpnKafkaReceiverTopicBinding**](KafkaReceiverApi.md#DeleteMsgVpnKafkaReceiverTopicBinding) | **Delete** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName}/topicBindings/{topicName} | Delete a Topic Binding object.
[**GetMsgVpnKafkaReceiver**](KafkaReceiverApi.md#GetMsgVpnKafkaReceiver) | **Get** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName} | Get a Kafka Receiver object.
[**GetMsgVpnKafkaReceiverTopicBinding**](KafkaReceiverApi.md#GetMsgVpnKafkaReceiverTopicBinding) | **Get** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName}/topicBindings/{topicName} | Get a Topic Binding object.
[**GetMsgVpnKafkaReceiverTopicBindings**](KafkaReceiverApi.md#GetMsgVpnKafkaReceiverTopicBindings) | **Get** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName}/topicBindings | Get a list of Topic Binding objects.
[**GetMsgVpnKafkaReceivers**](KafkaReceiverApi.md#GetMsgVpnKafkaReceivers) | **Get** /msgVpns/{msgVpnName}/kafkaReceivers | Get a list of Kafka Receiver objects.
[**ReplaceMsgVpnKafkaReceiver**](KafkaReceiverApi.md#ReplaceMsgVpnKafkaReceiver) | **Put** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName} | Replace a Kafka Receiver object.
[**ReplaceMsgVpnKafkaReceiverTopicBinding**](KafkaReceiverApi.md#ReplaceMsgVpnKafkaReceiverTopicBinding) | **Put** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName}/topicBindings/{topicName} | Replace a Topic Binding object.
[**UpdateMsgVpnKafkaReceiver**](KafkaReceiverApi.md#UpdateMsgVpnKafkaReceiver) | **Patch** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName} | Update a Kafka Receiver object.
[**UpdateMsgVpnKafkaReceiverTopicBinding**](KafkaReceiverApi.md#UpdateMsgVpnKafkaReceiverTopicBinding) | **Patch** /msgVpns/{msgVpnName}/kafkaReceivers/{kafkaReceiverName}/topicBindings/{topicName} | Update a Topic Binding object.

# **CreateMsgVpnKafkaReceiver**
> MsgVpnKafkaReceiverResponse CreateMsgVpnKafkaReceiver(ctx, body, msgVpnName, optional)
Create a Kafka Receiver object.

Create a Kafka Receiver object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Kafka Receiver receives messages from a Kafka Cluster.   Attribute|Identifying|Const|Required|Read-Only|Write-Only|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationAwsMskIamSecretAccessKey|||||x| authenticationBasicPassword|||||x|x authenticationClientCertContent|||||x|x authenticationClientCertPassword|||||x| authenticationKerberosKeytabContent|||||x| authenticationOauthClientSecret|||||x|x authenticationScramPassword|||||x|x kafkaReceiverName|x|x|x||| msgVpnName|x|||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires :---|:---|:--- MsgVpnKafkaReceiver|authenticationBasicPassword|authenticationBasicUsername MsgVpnKafkaReceiver|authenticationClientCertPassword|authenticationClientCertContent MsgVpnKafkaReceiver|authenticationKerberosKeytabContent|authenticationKerberosKeytabFileName, authenticationKerberosUserPrincipalName MsgVpnKafkaReceiver|authenticationKerberosKeytabFileName|authenticationKerberosKeytabContent, authenticationKerberosUserPrincipalName MsgVpnKafkaReceiver|authenticationKerberosUserPrincipalName|authenticationKerberosKeytabContent, authenticationKerberosKeytabFileName MsgVpnKafkaReceiver|authenticationScramPassword|authenticationScramUsername    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaReceiver**](MsgVpnKafkaReceiver.md)| The Kafka Receiver object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***KafkaReceiverApiCreateMsgVpnKafkaReceiverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiCreateMsgVpnKafkaReceiverOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverResponse**](MsgVpnKafkaReceiverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnKafkaReceiverTopicBinding**
> MsgVpnKafkaReceiverTopicBindingResponse CreateMsgVpnKafkaReceiverTopicBinding(ctx, body, msgVpnName, kafkaReceiverName, optional)
Create a Topic Binding object.

Create a Topic Binding object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Binding receives messages from a remote Kafka Topic.   Attribute|Identifying|Const|Required|Read-Only :---|:---:|:---:|:---:|:---: kafkaReceiverName|x|||x msgVpnName|x|||x topicName|x|x|x|    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaReceiverTopicBinding**](MsgVpnKafkaReceiverTopicBinding.md)| The Topic Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
 **optional** | ***KafkaReceiverApiCreateMsgVpnKafkaReceiverTopicBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiCreateMsgVpnKafkaReceiverTopicBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverTopicBindingResponse**](MsgVpnKafkaReceiverTopicBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnKafkaReceiver**
> SempMetaOnlyResponse DeleteMsgVpnKafkaReceiver(ctx, msgVpnName, kafkaReceiverName)
Delete a Kafka Receiver object.

Delete a Kafka Receiver object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Kafka Receiver receives messages from a Kafka Cluster.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnKafkaReceiverTopicBinding**
> SempMetaOnlyResponse DeleteMsgVpnKafkaReceiverTopicBinding(ctx, msgVpnName, kafkaReceiverName, topicName)
Delete a Topic Binding object.

Delete a Topic Binding object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Binding receives messages from a remote Kafka Topic.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
  **topicName** | **string**| The name of the Topic or a POSIX.2 regular expression starting with &#x27;^&#x27;. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnKafkaReceiver**
> MsgVpnKafkaReceiverResponse GetMsgVpnKafkaReceiver(ctx, msgVpnName, kafkaReceiverName, optional)
Get a Kafka Receiver object.

Get a Kafka Receiver object.  A Kafka Receiver receives messages from a Kafka Cluster.   Attribute|Identifying|Write-Only|Opaque :---|:---:|:---:|:---: authenticationAwsMskIamSecretAccessKey||x| authenticationBasicPassword||x|x authenticationClientCertContent||x|x authenticationClientCertPassword||x| authenticationKerberosKeytabContent||x| authenticationOauthClientSecret||x|x authenticationScramPassword||x|x kafkaReceiverName|x|| msgVpnName|x||    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
 **optional** | ***KafkaReceiverApiGetMsgVpnKafkaReceiverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiGetMsgVpnKafkaReceiverOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverResponse**](MsgVpnKafkaReceiverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnKafkaReceiverTopicBinding**
> MsgVpnKafkaReceiverTopicBindingResponse GetMsgVpnKafkaReceiverTopicBinding(ctx, msgVpnName, kafkaReceiverName, topicName, optional)
Get a Topic Binding object.

Get a Topic Binding object.  A Topic Binding receives messages from a remote Kafka Topic.   Attribute|Identifying :---|:---: kafkaReceiverName|x msgVpnName|x topicName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
  **topicName** | **string**| The name of the Topic or a POSIX.2 regular expression starting with &#x27;^&#x27;. | 
 **optional** | ***KafkaReceiverApiGetMsgVpnKafkaReceiverTopicBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiGetMsgVpnKafkaReceiverTopicBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverTopicBindingResponse**](MsgVpnKafkaReceiverTopicBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnKafkaReceiverTopicBindings**
> MsgVpnKafkaReceiverTopicBindingsResponse GetMsgVpnKafkaReceiverTopicBindings(ctx, msgVpnName, kafkaReceiverName, optional)
Get a list of Topic Binding objects.

Get a list of Topic Binding objects.  A Topic Binding receives messages from a remote Kafka Topic.   Attribute|Identifying :---|:---: kafkaReceiverName|x msgVpnName|x topicName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 100.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
 **optional** | ***KafkaReceiverApiGetMsgVpnKafkaReceiverTopicBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiGetMsgVpnKafkaReceiverTopicBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverTopicBindingsResponse**](MsgVpnKafkaReceiverTopicBindingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnKafkaReceivers**
> MsgVpnKafkaReceiversResponse GetMsgVpnKafkaReceivers(ctx, msgVpnName, optional)
Get a list of Kafka Receiver objects.

Get a list of Kafka Receiver objects.  A Kafka Receiver receives messages from a Kafka Cluster.   Attribute|Identifying|Write-Only|Opaque :---|:---:|:---:|:---: authenticationAwsMskIamSecretAccessKey||x| authenticationBasicPassword||x|x authenticationClientCertContent||x|x authenticationClientCertPassword||x| authenticationKerberosKeytabContent||x| authenticationOauthClientSecret||x|x authenticationScramPassword||x|x kafkaReceiverName|x|| msgVpnName|x||    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 100.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***KafkaReceiverApiGetMsgVpnKafkaReceiversOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiGetMsgVpnKafkaReceiversOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiversResponse**](MsgVpnKafkaReceiversResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnKafkaReceiver**
> MsgVpnKafkaReceiverResponse ReplaceMsgVpnKafkaReceiver(ctx, body, msgVpnName, kafkaReceiverName, optional)
Replace a Kafka Receiver object.

Replace a Kafka Receiver object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  A Kafka Receiver receives messages from a Kafka Cluster.   Attribute|Identifying|Const|Read-Only|Write-Only|Auto-Disable|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationAwsMskIamAccessKeyId|||||x| authenticationAwsMskIamRegion|||||x| authenticationAwsMskIamSecretAccessKey||||x|x| authenticationAwsMskIamStsExternalId|||||x| authenticationAwsMskIamStsRoleArn|||||x| authenticationAwsMskIamStsRoleSessionName|||||x| authenticationBasicPassword||||x|x|x authenticationBasicUsername|||||x| authenticationClientCertContent||||x|x|x authenticationClientCertPassword||||x|x| authenticationKerberosKeytabContent||||x|x| authenticationKerberosKeytabFileName|||||x| authenticationKerberosServiceName|||||x| authenticationKerberosUserPrincipalName|||||x| authenticationOauthClientId|||||x| authenticationOauthClientScope|||||x| authenticationOauthClientSecret||||x|x|x authenticationOauthClientTokenEndpoint|||||x| authenticationScheme|||||x| authenticationScramHash|||||x| authenticationScramPassword||||x|x|x authenticationScramUsername|||||x| batchDelay|||||x| batchMaxSize|||||x| bootstrapAddressList|||||x| groupId|||||x| groupKeepaliveInterval|||||x| groupKeepaliveTimeout|||||x| groupMembershipType|||||x| groupPartitionSchemeList|||||x| kafkaReceiverName|x|x|||| metadataTopicExcludeList|||||x| metadataTopicRefreshInterval|||||x| msgVpnName|x||x||| transportTlsEnabled|||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires :---|:---|:--- MsgVpnKafkaReceiver|authenticationBasicPassword|authenticationBasicUsername MsgVpnKafkaReceiver|authenticationClientCertPassword|authenticationClientCertContent MsgVpnKafkaReceiver|authenticationKerberosKeytabContent|authenticationKerberosKeytabFileName, authenticationKerberosUserPrincipalName MsgVpnKafkaReceiver|authenticationKerberosKeytabFileName|authenticationKerberosKeytabContent, authenticationKerberosUserPrincipalName MsgVpnKafkaReceiver|authenticationKerberosUserPrincipalName|authenticationKerberosKeytabContent, authenticationKerberosKeytabFileName MsgVpnKafkaReceiver|authenticationScramPassword|authenticationScramUsername    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"vpn/read-write\" is required if this operation creates an object.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaReceiver**](MsgVpnKafkaReceiver.md)| The Kafka Receiver object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
 **optional** | ***KafkaReceiverApiReplaceMsgVpnKafkaReceiverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiReplaceMsgVpnKafkaReceiverOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverResponse**](MsgVpnKafkaReceiverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnKafkaReceiverTopicBinding**
> MsgVpnKafkaReceiverTopicBindingResponse ReplaceMsgVpnKafkaReceiverTopicBinding(ctx, body, msgVpnName, kafkaReceiverName, topicName, optional)
Replace a Topic Binding object.

Replace a Topic Binding object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  A Topic Binding receives messages from a remote Kafka Topic.   Attribute|Identifying|Const|Read-Only|Auto-Disable :---|:---:|:---:|:---:|:---: initialOffset||||x kafkaReceiverName|x||x| localKey||||x localTopic||||x msgVpnName|x||x| topicName|x|x||    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"vpn/read-write\" is required if this operation creates an object.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaReceiverTopicBinding**](MsgVpnKafkaReceiverTopicBinding.md)| The Topic Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
  **topicName** | **string**| The name of the Topic or a POSIX.2 regular expression starting with &#x27;^&#x27;. | 
 **optional** | ***KafkaReceiverApiReplaceMsgVpnKafkaReceiverTopicBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiReplaceMsgVpnKafkaReceiverTopicBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverTopicBindingResponse**](MsgVpnKafkaReceiverTopicBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnKafkaReceiver**
> MsgVpnKafkaReceiverResponse UpdateMsgVpnKafkaReceiver(ctx, body, msgVpnName, kafkaReceiverName, optional)
Update a Kafka Receiver object.

Update a Kafka Receiver object. Any attribute missing from the request will be left unchanged.  A Kafka Receiver receives messages from a Kafka Cluster.   Attribute|Identifying|Const|Read-Only|Write-Only|Auto-Disable|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationAwsMskIamAccessKeyId|||||x| authenticationAwsMskIamRegion|||||x| authenticationAwsMskIamSecretAccessKey||||x|x| authenticationAwsMskIamStsExternalId|||||x| authenticationAwsMskIamStsRoleArn|||||x| authenticationAwsMskIamStsRoleSessionName|||||x| authenticationBasicPassword||||x|x|x authenticationBasicUsername|||||x| authenticationClientCertContent||||x|x|x authenticationClientCertPassword||||x|x| authenticationKerberosKeytabContent||||x|x| authenticationKerberosKeytabFileName|||||x| authenticationKerberosServiceName|||||x| authenticationKerberosUserPrincipalName|||||x| authenticationOauthClientId|||||x| authenticationOauthClientScope|||||x| authenticationOauthClientSecret||||x|x|x authenticationOauthClientTokenEndpoint|||||x| authenticationScheme|||||x| authenticationScramHash|||||x| authenticationScramPassword||||x|x|x authenticationScramUsername|||||x| batchDelay|||||x| batchMaxSize|||||x| bootstrapAddressList|||||x| groupId|||||x| groupKeepaliveInterval|||||x| groupKeepaliveTimeout|||||x| groupMembershipType|||||x| groupPartitionSchemeList|||||x| kafkaReceiverName|x|x|||| metadataTopicExcludeList|||||x| metadataTopicRefreshInterval|||||x| msgVpnName|x||x||| transportTlsEnabled|||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires :---|:---|:--- MsgVpnKafkaReceiver|authenticationBasicPassword|authenticationBasicUsername MsgVpnKafkaReceiver|authenticationClientCertPassword|authenticationClientCertContent MsgVpnKafkaReceiver|authenticationKerberosKeytabContent|authenticationKerberosKeytabFileName, authenticationKerberosUserPrincipalName MsgVpnKafkaReceiver|authenticationKerberosKeytabFileName|authenticationKerberosKeytabContent, authenticationKerberosUserPrincipalName MsgVpnKafkaReceiver|authenticationKerberosUserPrincipalName|authenticationKerberosKeytabContent, authenticationKerberosKeytabFileName MsgVpnKafkaReceiver|authenticationScramPassword|authenticationScramUsername    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaReceiver**](MsgVpnKafkaReceiver.md)| The Kafka Receiver object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
 **optional** | ***KafkaReceiverApiUpdateMsgVpnKafkaReceiverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiUpdateMsgVpnKafkaReceiverOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverResponse**](MsgVpnKafkaReceiverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnKafkaReceiverTopicBinding**
> MsgVpnKafkaReceiverTopicBindingResponse UpdateMsgVpnKafkaReceiverTopicBinding(ctx, body, msgVpnName, kafkaReceiverName, topicName, optional)
Update a Topic Binding object.

Update a Topic Binding object. Any attribute missing from the request will be left unchanged.  A Topic Binding receives messages from a remote Kafka Topic.   Attribute|Identifying|Const|Read-Only|Auto-Disable :---|:---:|:---:|:---:|:---: initialOffset||||x kafkaReceiverName|x||x| localKey||||x localTopic||||x msgVpnName|x||x| topicName|x|x||    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnKafkaReceiverTopicBinding**](MsgVpnKafkaReceiverTopicBinding.md)| The Topic Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kafkaReceiverName** | **string**| The name of the Kafka Receiver. | 
  **topicName** | **string**| The name of the Topic or a POSIX.2 regular expression starting with &#x27;^&#x27;. | 
 **optional** | ***KafkaReceiverApiUpdateMsgVpnKafkaReceiverTopicBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaReceiverApiUpdateMsgVpnKafkaReceiverTopicBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnKafkaReceiverTopicBindingResponse**](MsgVpnKafkaReceiverTopicBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

