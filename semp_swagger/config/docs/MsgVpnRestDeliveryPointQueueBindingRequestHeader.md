# MsgVpnRestDeliveryPointQueueBindingRequestHeader

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderName** | **string** | The name of the HTTP request header.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. | [optional] [default to null]
**HeaderValue** | **string** | A substitution expression for the value of the HTTP request header.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;vpn/read-write\&quot;. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. | [optional] [default to null]
**QueueBindingName** | **string** | The name of a queue in the Message VPN.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. | [optional] [default to null]
**RestDeliveryPointName** | **string** | The name of the REST Delivery Point.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

