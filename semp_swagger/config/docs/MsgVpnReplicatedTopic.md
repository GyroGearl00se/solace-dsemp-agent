# MsgVpnReplicatedTopic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgVpnName** | **string** | The name of the Message VPN.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. | [optional] [default to null]
**ReplicatedTopic** | **string** | The topic for applying replication. Published messages matching this topic will be replicated to the standby site.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. | [optional] [default to null]
**ReplicationMode** | **string** | The replication mode for the Replicated Topic.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;vpn/read-write\&quot;. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;\&quot;async\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;sync\&quot; - Messages are acknowledged when replicated (spooled remotely). \&quot;async\&quot; - Messages are acknowledged when pending replication (spooled locally). &lt;/pre&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

