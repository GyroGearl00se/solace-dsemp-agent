# OauthProfileAccessLevelGroupMsgVpnAccessLevelException

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | **string** | The message VPN access level.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;global/read-write\&quot;. Changes to this attribute are synchronized to HA mates via config-sync. The default value is &#x60;\&quot;none\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;none\&quot; - User has no access to a Message VPN. \&quot;read-only\&quot; - User has read-only access to a Message VPN. \&quot;read-write\&quot; - User has read-write access to most Message VPN settings. &lt;/pre&gt;  | [optional] [default to null]
**GroupName** | **string** | The name of the group.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the message VPN.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. | [optional] [default to null]
**OauthProfileName** | **string** | The name of the OAuth profile.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

