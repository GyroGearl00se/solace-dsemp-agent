# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnAuthenticationKerberosRealm**](AuthenticationKerberosRealmApi.md#CreateMsgVpnAuthenticationKerberosRealm) | **Post** /msgVpns/{msgVpnName}/authenticationKerberosRealms | Create a Realm object.
[**DeleteMsgVpnAuthenticationKerberosRealm**](AuthenticationKerberosRealmApi.md#DeleteMsgVpnAuthenticationKerberosRealm) | **Delete** /msgVpns/{msgVpnName}/authenticationKerberosRealms/{kerberosRealmName} | Delete a Realm object.
[**GetMsgVpnAuthenticationKerberosRealm**](AuthenticationKerberosRealmApi.md#GetMsgVpnAuthenticationKerberosRealm) | **Get** /msgVpns/{msgVpnName}/authenticationKerberosRealms/{kerberosRealmName} | Get a Realm object.
[**GetMsgVpnAuthenticationKerberosRealms**](AuthenticationKerberosRealmApi.md#GetMsgVpnAuthenticationKerberosRealms) | **Get** /msgVpns/{msgVpnName}/authenticationKerberosRealms | Get a list of Realm objects.
[**ReplaceMsgVpnAuthenticationKerberosRealm**](AuthenticationKerberosRealmApi.md#ReplaceMsgVpnAuthenticationKerberosRealm) | **Put** /msgVpns/{msgVpnName}/authenticationKerberosRealms/{kerberosRealmName} | Replace a Realm object.
[**UpdateMsgVpnAuthenticationKerberosRealm**](AuthenticationKerberosRealmApi.md#UpdateMsgVpnAuthenticationKerberosRealm) | **Patch** /msgVpns/{msgVpnName}/authenticationKerberosRealms/{kerberosRealmName} | Update a Realm object.

# **CreateMsgVpnAuthenticationKerberosRealm**
> MsgVpnAuthenticationKerberosRealmResponse CreateMsgVpnAuthenticationKerberosRealm(ctx, body, msgVpnName, optional)
Create a Realm object.

Create a Realm object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Kerberos Realm.   Attribute|Identifying|Const|Required|Read-Only :---|:---:|:---:|:---:|:---: kerberosRealmName|x|x|x| msgVpnName|x|||x    The minimum access scope/level required to perform this operation is \"vpn/read-write\". A different access scope/level may also be required when providing values for specific attributes.  This has been available since 2.40.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationKerberosRealm**](MsgVpnAuthenticationKerberosRealm.md)| The Realm object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AuthenticationKerberosRealmApiCreateMsgVpnAuthenticationKerberosRealmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationKerberosRealmApiCreateMsgVpnAuthenticationKerberosRealmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationKerberosRealmResponse**](MsgVpnAuthenticationKerberosRealmResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthenticationKerberosRealm**
> SempMetaOnlyResponse DeleteMsgVpnAuthenticationKerberosRealm(ctx, msgVpnName, kerberosRealmName)
Delete a Realm object.

Delete a Realm object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Kerberos Realm.  The minimum access scope/level required to perform this operation is \"vpn/read-write\".  This has been available since 2.40.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kerberosRealmName** | **string**| The Realm Name. Must start with \&quot;@\&quot;, typically all uppercase. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationKerberosRealm**
> MsgVpnAuthenticationKerberosRealmResponse GetMsgVpnAuthenticationKerberosRealm(ctx, msgVpnName, kerberosRealmName, optional)
Get a Realm object.

Get a Realm object.  Kerberos Realm.   Attribute|Identifying :---|:---: kerberosRealmName|x msgVpnName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  This has been available since 2.40.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kerberosRealmName** | **string**| The Realm Name. Must start with \&quot;@\&quot;, typically all uppercase. | 
 **optional** | ***AuthenticationKerberosRealmApiGetMsgVpnAuthenticationKerberosRealmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationKerberosRealmApiGetMsgVpnAuthenticationKerberosRealmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationKerberosRealmResponse**](MsgVpnAuthenticationKerberosRealmResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationKerberosRealms**
> MsgVpnAuthenticationKerberosRealmsResponse GetMsgVpnAuthenticationKerberosRealms(ctx, msgVpnName, optional)
Get a list of Realm objects.

Get a list of Realm objects.  Kerberos Realm.   Attribute|Identifying :---|:---: kerberosRealmName|x msgVpnName|x    The minimum access scope/level required to perform this operation is \"vpn/read-only\".  The maximum number of objects that can be returned in a single page is 100.  This has been available since 2.40.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AuthenticationKerberosRealmApiGetMsgVpnAuthenticationKerberosRealmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationKerberosRealmApiGetMsgVpnAuthenticationKerberosRealmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationKerberosRealmsResponse**](MsgVpnAuthenticationKerberosRealmsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnAuthenticationKerberosRealm**
> MsgVpnAuthenticationKerberosRealmResponse ReplaceMsgVpnAuthenticationKerberosRealm(ctx, body, msgVpnName, kerberosRealmName, optional)
Replace a Realm object.

Replace a Realm object. Any attribute missing from the request will be set to its default value, subject to the exceptions [here](https://docs.solace.com/Admin/SEMP/SEMP-API-Archit.htm#HTTP_Methods).  Kerberos Realm.   Attribute|Identifying|Const|Read-Only|Auto-Disable :---|:---:|:---:|:---:|:---: kdcAddress||||x kerberosRealmName|x|x|| msgVpnName|x||x|    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request. In addition, a minimum access scope/level of \"vpn/read-write\" is required if this operation creates an object.  This has been available since 2.40.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationKerberosRealm**](MsgVpnAuthenticationKerberosRealm.md)| The Realm object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kerberosRealmName** | **string**| The Realm Name. Must start with \&quot;@\&quot;, typically all uppercase. | 
 **optional** | ***AuthenticationKerberosRealmApiReplaceMsgVpnAuthenticationKerberosRealmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationKerberosRealmApiReplaceMsgVpnAuthenticationKerberosRealmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationKerberosRealmResponse**](MsgVpnAuthenticationKerberosRealmResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnAuthenticationKerberosRealm**
> MsgVpnAuthenticationKerberosRealmResponse UpdateMsgVpnAuthenticationKerberosRealm(ctx, body, msgVpnName, kerberosRealmName, optional)
Update a Realm object.

Update a Realm object. Any attribute missing from the request will be left unchanged.  Kerberos Realm.   Attribute|Identifying|Const|Read-Only|Auto-Disable :---|:---:|:---:|:---:|:---: kdcAddress||||x kerberosRealmName|x|x|| msgVpnName|x||x|    The minimum access scope/level required to perform this operation is determined by the attributes provided in the request.  This has been available since 2.40.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationKerberosRealm**](MsgVpnAuthenticationKerberosRealm.md)| The Realm object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **kerberosRealmName** | **string**| The Realm Name. Must start with \&quot;@\&quot;, typically all uppercase. | 
 **optional** | ***AuthenticationKerberosRealmApiUpdateMsgVpnAuthenticationKerberosRealmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationKerberosRealmApiUpdateMsgVpnAuthenticationKerberosRealmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationKerberosRealmResponse**](MsgVpnAuthenticationKerberosRealmResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

