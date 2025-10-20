# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnProxy**](ProxyApi.md#CreateMsgVpnProxy) | **Post** /msgVpns/{msgVpnName}/proxies | Create a Proxy object.
[**CreateProxy**](ProxyApi.md#CreateProxy) | **Post** /proxies | Create a Proxy object.
[**DeleteMsgVpnProxy**](ProxyApi.md#DeleteMsgVpnProxy) | **Delete** /msgVpns/{msgVpnName}/proxies/{proxyName} | Delete a Proxy object.
[**DeleteProxy**](ProxyApi.md#DeleteProxy) | **Delete** /proxies/{proxyName} | Delete a Proxy object.
[**GetMsgVpnProxies**](ProxyApi.md#GetMsgVpnProxies) | **Get** /msgVpns/{msgVpnName}/proxies | Get a list of Proxy objects.
[**GetMsgVpnProxy**](ProxyApi.md#GetMsgVpnProxy) | **Get** /msgVpns/{msgVpnName}/proxies/{proxyName} | Get a Proxy object.
[**GetProxies**](ProxyApi.md#GetProxies) | **Get** /proxies | Get a list of Proxy objects.
[**GetProxy**](ProxyApi.md#GetProxy) | **Get** /proxies/{proxyName} | Get a Proxy object.
[**ReplaceMsgVpnProxy**](ProxyApi.md#ReplaceMsgVpnProxy) | **Put** /msgVpns/{msgVpnName}/proxies/{proxyName} | Replace a Proxy object.
[**ReplaceProxy**](ProxyApi.md#ReplaceProxy) | **Put** /proxies/{proxyName} | Replace a Proxy object.
[**UpdateMsgVpnProxy**](ProxyApi.md#UpdateMsgVpnProxy) | **Patch** /msgVpns/{msgVpnName}/proxies/{proxyName} | Update a Proxy object.
[**UpdateProxy**](ProxyApi.md#UpdateProxy) | **Patch** /proxies/{proxyName} | Update a Proxy object.

# **CreateMsgVpnProxy**
> MsgVpnProxyResponse CreateMsgVpnProxy(ctx, body, msgVpnName, optional)
Create a Proxy object.

Create a Proxy object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a REST Consumer, select the proxy by name in the configuration for that object.   Attribute|Identifying|Const|Required|Read-Only|Write-Only|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationBasicPassword|||||x|x msgVpnName|x|||x|| proxyName|x|x|x|||    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnProxy**](MsgVpnProxy.md)| The Proxy object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***ProxyApiCreateMsgVpnProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiCreateMsgVpnProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnProxyResponse**](MsgVpnProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProxy**
> ProxyResponse CreateProxy(ctx, body, optional)
Create a Proxy object.

Create a Proxy object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a OAuth Provider, select the proxy by name in the configuration for that object.   Attribute|Identifying|Const|Required|Write-Only|Opaque :---|:---:|:---:|:---:|:---:|:---: authenticationBasicPassword||||x|x proxyName|x|x|x||    The minimum access scope/level required to perform this operation is \"global/admin\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.41.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Proxy**](Proxy.md)| The Proxy object&#x27;s attributes. | 
 **optional** | ***ProxyApiCreateProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiCreateProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ProxyResponse**](ProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnProxy**
> SempMetaOnlyResponse DeleteMsgVpnProxy(ctx, msgVpnName, proxyName)
Delete a Proxy object.

Delete a Proxy object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a REST Consumer, select the proxy by name in the configuration for that object.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **proxyName** | **string**| The name of the proxy. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProxy**
> SempMetaOnlyResponse DeleteProxy(ctx, proxyName)
Delete a Proxy object.

Delete a Proxy object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a OAuth Provider, select the proxy by name in the configuration for that object.  The minimum access scope/level required to perform this operation is \"global/admin\".  This has been available since 2.41.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyName** | **string**| The name of the proxy. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnProxies**
> MsgVpnProxiesResponse GetMsgVpnProxies(ctx, msgVpnName, optional)
Get a list of Proxy objects.

Get a list of Proxy objects.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a REST Consumer, select the proxy by name in the configuration for that object.   Attribute|Identifying|Write-Only|Opaque :---|:---:|:---:|:---: authenticationBasicPassword||x|x msgVpnName|x|| proxyName|x||    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 500.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***ProxyApiGetMsgVpnProxiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiGetMsgVpnProxiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnProxiesResponse**](MsgVpnProxiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnProxy**
> MsgVpnProxyResponse GetMsgVpnProxy(ctx, msgVpnName, proxyName, optional)
Get a Proxy object.

Get a Proxy object.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a REST Consumer, select the proxy by name in the configuration for that object.   Attribute|Identifying|Write-Only|Opaque :---|:---:|:---:|:---: authenticationBasicPassword||x|x msgVpnName|x|| proxyName|x||    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **proxyName** | **string**| The name of the proxy. | 
 **optional** | ***ProxyApiGetMsgVpnProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiGetMsgVpnProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnProxyResponse**](MsgVpnProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxies**
> ProxiesResponse GetProxies(ctx, optional)
Get a list of Proxy objects.

Get a list of Proxy objects.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a OAuth Provider, select the proxy by name in the configuration for that object.   Attribute|Identifying|Write-Only|Opaque :---|:---:|:---:|:---: authenticationBasicPassword||x|x proxyName|x||    The minimum access scope/level required to perform this operation is \"global/read-only\".  The maximum number of objects that can be returned in a single page is 500.  This has been available since 2.41.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProxyApiGetProxiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiGetProxiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ProxiesResponse**](ProxiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxy**
> ProxyResponse GetProxy(ctx, proxyName, optional)
Get a Proxy object.

Get a Proxy object.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a OAuth Provider, select the proxy by name in the configuration for that object.   Attribute|Identifying|Write-Only|Opaque :---|:---:|:---:|:---: authenticationBasicPassword||x|x proxyName|x||    The minimum access scope/level required to perform this operation is \"global/read-only\".  This has been available since 2.41.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyName** | **string**| The name of the proxy. | 
 **optional** | ***ProxyApiGetProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiGetProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ProxyResponse**](ProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnProxy**
> MsgVpnProxyResponse ReplaceMsgVpnProxy(ctx, body, msgVpnName, proxyName, optional)
Replace a Proxy object.

Replace a Proxy object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a REST Consumer, select the proxy by name in the configuration for that object.   Attribute|Identifying|Const|Read-Only|Write-Only|Opaque :---|:---:|:---:|:---:|:---:|:---: authenticationBasicPassword||||x|x msgVpnName|x||x|| proxyName|x|x|||    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"vpn/read-write\" is required if this operation creates an object.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnProxy**](MsgVpnProxy.md)| The Proxy object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **proxyName** | **string**| The name of the proxy. | 
 **optional** | ***ProxyApiReplaceMsgVpnProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiReplaceMsgVpnProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnProxyResponse**](MsgVpnProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceProxy**
> ProxyResponse ReplaceProxy(ctx, body, proxyName, optional)
Replace a Proxy object.

Replace a Proxy object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a OAuth Provider, select the proxy by name in the configuration for that object.   Attribute|Identifying|Const|Write-Only|Opaque :---|:---:|:---:|:---:|:---: authenticationBasicPassword|||x|x proxyName|x|x||    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"global/admin\" is required if this operation creates an object.  This has been available since 2.41.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Proxy**](Proxy.md)| The Proxy object&#x27;s attributes. | 
  **proxyName** | **string**| The name of the proxy. | 
 **optional** | ***ProxyApiReplaceProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiReplaceProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ProxyResponse**](ProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnProxy**
> MsgVpnProxyResponse UpdateMsgVpnProxy(ctx, body, msgVpnName, proxyName, optional)
Update a Proxy object.

Update a Proxy object. Any attribute missing from the request will be left unchanged.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a REST Consumer, select the proxy by name in the configuration for that object.   Attribute|Identifying|Const|Read-Only|Write-Only|Opaque :---|:---:|:---:|:---:|:---:|:---: authenticationBasicPassword||||x|x msgVpnName|x||x|| proxyName|x|x|||    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.36.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnProxy**](MsgVpnProxy.md)| The Proxy object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **proxyName** | **string**| The name of the proxy. | 
 **optional** | ***ProxyApiUpdateMsgVpnProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiUpdateMsgVpnProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnProxyResponse**](MsgVpnProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProxy**
> ProxyResponse UpdateProxy(ctx, body, proxyName, optional)
Update a Proxy object.

Update a Proxy object. Any attribute missing from the request will be left unchanged.  Proxy objects define the connection parameters for a proxy server. To use a proxy for a particular connection such as a OAuth Provider, select the proxy by name in the configuration for that object.   Attribute|Identifying|Const|Write-Only|Opaque :---|:---:|:---:|:---:|:---: authenticationBasicPassword|||x|x proxyName|x|x||    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.41.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Proxy**](Proxy.md)| The Proxy object&#x27;s attributes. | 
  **proxyName** | **string**| The name of the proxy. | 
 **optional** | ***ProxyApiUpdateProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiUpdateProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ProxyResponse**](ProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

