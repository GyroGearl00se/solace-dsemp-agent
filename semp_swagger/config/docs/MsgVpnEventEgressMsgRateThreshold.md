# MsgVpnEventEgressMsgRateThreshold

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearValue** | **int64** | The clear threshold for the absolute value of this counter or rate. Falling below this value will trigger a corresponding event.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;vpn/read-write\&quot;. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;3000000&#x60;. | [optional] [default to null]
**SetValue** | **int64** | The set threshold for the absolute value of this counter or rate. Exceeding this value will trigger a corresponding event.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;vpn/read-write\&quot;. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;4000000&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

