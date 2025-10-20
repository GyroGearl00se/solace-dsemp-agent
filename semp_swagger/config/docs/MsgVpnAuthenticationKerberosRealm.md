# MsgVpnAuthenticationKerberosRealm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | [****bool**](*bool.md) | Enable or disable the Realm.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;vpn/read-write\&quot;. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]
**KdcAddress** | **string** | Address (FQDN or IP) and optional port of the Key Distribution Center for principals in this Realm.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;vpn/read-write\&quot;. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**KerberosRealmName** | **string** | The Realm Name. Must start with \&quot;@\&quot;, typically all uppercase.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN.  The minimum access scope/level required to retrieve this attribute is \&quot;vpn/read-only\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

