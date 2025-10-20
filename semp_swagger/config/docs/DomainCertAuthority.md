# DomainCertAuthority

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertAuthorityName** | **string** | The name of the Certificate Authority.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. | [optional] [default to null]
**CertContent** | **string** | The PEM formatted content for the trusted root certificate of a domain Certificate Authority.  The minimum access scope/level required to retrieve this attribute is \&quot;global/read-only\&quot;. The minimum access scope/level required to change this attribute is \&quot;global/admin\&quot;. Changes to this attribute are synchronized to HA mates via config-sync. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

