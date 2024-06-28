/*
Public EMMA API

### About Infrastructure API   **Base URL:** **<u>https://api.emma.ms/external</u>**   This **Infrastructure API** is for managing the emma cloud infrastructure within a project. The API enables you to view, create, edit, and delete _Virtual machines, Spot instances, Applications, Kubernetes clusters, SSH keys, Security groups, and Backup policies_. For creating the resources you can use the endpoints with the dictionaries: _Data centers, Locations, Providers, Operating systems, Virtual machines configurations, Spot instances configurations, Kubernetes clusters configurations._   ### Authentication   #### 1. Create service application   To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.   - **Read** - only GET methods are allowed in the API.   - **Operate** - some operations are allowed with the resources (e.g. _Start, Reboot,_ and _Shutdown_ of the Virtual machines).   - **Manage** - full creating, updating, and deleting of the resources is allowed.     #### 2. Get access token   - Copy the **Client ID** and **Client Secret** in the service application.  - Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.   _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token \\  -H \"Content-Type: application/json\" \\  -d '{\"clientId\": \"YOUR-CLIENT-ID\", \"clientSecret\": \"YOUR-CLIENT-SECRET\"}'  ```  _For Windows:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token ^  -H \"Content-Type: application/json\" ^  -d \"{\\\"clientId\\\": \\\"YOUR-CLIENT-ID\\\", \\\"clientSecret\\\": \\\"YOUR-CLIENT-SECRET\\\"}\"  ```      #### 3. Use access token in requests  The Bearer access token is a text string, included in the request header, for example:   _For Linux / Mac:_  ```  curl -X GET https://api.emma.ms/external/v1/locations -H \"Authorization: Bearer YOUR-ACCESS-TOKEN-HERE\"  ```   Use this token for the API requests.     #### 4. Refresh token  The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).   To get a new access token send a request to the **_/refresh-token** endpoint:    _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/refresh-token \\  -H \"Content-Type: application/json\" \\  -H \"Authorization: Bearer YOUR-ACCESS-TOKEN\" \\  -d '{\"refreshToken\": \"YOUR-REFRESH-TOKEN\"}'  ```       ### Possible response status codes   We use standard HTTP response codes to show the success or failure of requests.   `2xx` - successful responses.   `4xx` - client error responses (the response contains an explanation of the error).   `5xx` - server error responses.   The API uses the following status codes:   | Status Code | Description                  | Notes                                                                  |  |-------------|------------------------------|------------------------------------------------------------------------|  | 200         | OK                           | The request was successful.                                             |  | 201         | Created                      | The object was successfully created. This code is only used with objects that are created immediately.  | 400         | Bad Request                  | The request could not be understood by the server. Incoming parameters might not be valid. |  | 401         | Unauthorized            | The client is unauthenticated. The client must authenticate itself to get the requested response. |  | 403         | Forbidden                   | The client does not have access rights to the content.  | 404         | Not Found                    | The requested resource is not found.                                    |  | 409         | Conflict | This response is sent when a request conflicts with the current state of the object (e.g. deleting the security group with the compute instances in it).|  | 422         | Unprocessable Content   | The request was well-formed but was unable to be followed due to incorrect field values (e.g. creation of a virtual machine in the non-existent data center).  |  | 500         | Internal server Error                 | The server could not return the representation due to an internal server error. |

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
)

// checks if the SshKeyGenerated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshKeyGenerated{}

// SshKeyGenerated struct for SshKeyGenerated
type SshKeyGenerated struct {
	// SSH key ID
	Id *int32 `json:"id,omitempty"`
	// SSH key name
	Name *string `json:"name,omitempty"`
	// SSH public key
	Key *string `json:"key,omitempty"`
	// SSH key fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`
	// SSH key type (RSA or ED25519)
	KeyType *string `json:"keyType,omitempty"`
	// SSH private key
	PrivateKey *string `json:"privateKey,omitempty"`
	// Date and time when the SSH key was created
	CreatedAt *string `json:"createdAt,omitempty"`
	// Name of the user who created the SSH key
	CreatedByName *string `json:"createdByName,omitempty"`
	// ID of the user who created the SSH key
	CreatedById *int32 `json:"createdById,omitempty"`
}

// NewSshKeyGenerated instantiates a new SshKeyGenerated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyGenerated() *SshKeyGenerated {
	this := SshKeyGenerated{}
	return &this
}

// NewSshKeyGeneratedWithDefaults instantiates a new SshKeyGenerated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyGeneratedWithDefaults() *SshKeyGenerated {
	this := SshKeyGenerated{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SshKeyGenerated) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SshKeyGenerated) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SshKeyGenerated) SetKey(v string) {
	o.Key = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *SshKeyGenerated) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetKeyType() string {
	if o == nil || IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasKeyType() bool {
	if o != nil && !IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *SshKeyGenerated) SetKeyType(v string) {
	o.KeyType = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *SshKeyGenerated) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SshKeyGenerated) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetCreatedByName() string {
	if o == nil || IsNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetCreatedByNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByName) {
		return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasCreatedByName() bool {
	if o != nil && !IsNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *SshKeyGenerated) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *SshKeyGenerated) GetCreatedById() int32 {
	if o == nil || IsNil(o.CreatedById) {
		var ret int32
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyGenerated) GetCreatedByIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *SshKeyGenerated) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given int32 and assigns it to the CreatedById field.
func (o *SshKeyGenerated) SetCreatedById(v int32) {
	o.CreatedById = &v
}

func (o SshKeyGenerated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshKeyGenerated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !IsNil(o.KeyType) {
		toSerialize["keyType"] = o.KeyType
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatedByName) {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	return toSerialize, nil
}

type NullableSshKeyGenerated struct {
	value *SshKeyGenerated
	isSet bool
}

func (v NullableSshKeyGenerated) Get() *SshKeyGenerated {
	return v.value
}

func (v *NullableSshKeyGenerated) Set(val *SshKeyGenerated) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyGenerated) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyGenerated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyGenerated(val *SshKeyGenerated) *NullableSshKeyGenerated {
	return &NullableSshKeyGenerated{value: val, isSet: true}
}

func (v NullableSshKeyGenerated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyGenerated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
