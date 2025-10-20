# VirtualHostname

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | [****bool**](*bool.md) | Enable or disable Virtual Hostname to Message VPN mapping.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;global/read-write\&quot;. Changes to this attribute are synchronized to HA mates via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The message VPN to which this virtual hostname is mapped.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;global/read-write\&quot;. Changes to this attribute are synchronized to HA mates via config-sync. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**VirtualHostname** | **string** | The virtual hostname.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

