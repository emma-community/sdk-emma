/*
Public EMMA API

### About Infrastructure API  **Base URL:** **<u>https://api.emma.ms/external</u>**   This **Infrastructure API** is for managing the emma cloud infrastructure within a project. The API enables you to view, create, edit, and delete _Virtual machines, Spot instances, Applications, Kubernetes clusters, SSH keys, Security groups, and Backup policies_. For creating the resources you can use the endpoints with the dictionaries: _Data centers, Locations, Providers, Operating systems, Virtual machines configurations, Spot instances configurations, Kubernetes clusters configurations._   ### Authentication   #### 1. Create service application   To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.   - **Read** - only GET methods are allowed in the API.   - **Operate** - some operations are allowed with the resources (e.g. _Start, Reboot,_ and _Shutdown_ of the Virtual machines).   - **Manage** - full creating, updating, and deleting of the resources is allowed.    #### 2. Get access token   - Copy the **Client ID** and **Client Secret** in the service application.  - Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.   _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token \\  -H \"Content-Type: application/json\" \\  -d '{\"clientId\": \"YOUR-CLIENT-ID\", \"clientSecret\": \"YOUR-CLIENT-SECRET\"}'  ```  _For Windows:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token ^  -H \"Content-Type: application/json\" ^  -d \"{\\\"clientId\\\": \\\"YOUR-CLIENT-ID\\\", \\\"clientSecret\\\": \\\"YOUR-CLIENT-SECRET\\\"}\"  ```   #### 3. Use access token in requests  The Bearer access token is a text string, included in the request header, for example:   _For Linux / Mac:_  ```  curl -X GET https://api.emma.ms/external/v1/locations -H \"Authorization: Bearer YOUR-ACCESS-TOKEN-HERE\"  ```   Use this token for the API requests.    #### 4. Refresh token  The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).   To get a new access token send a request to the **_/refresh-token** endpoint:    _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/refresh-token \\  -H \"Content-Type: application/json\" \\  -d '{\"refreshToken\": \"YOUR-REFRESH-TOKEN\"}'  ```       ### Possible response status codes   We use standard HTTP response codes to show the success or failure of requests.   `2xx` - successful responses.   `4xx` - client error responses (the response contains an explanation of the error).   `5xx` - server error responses.   The API uses the following status codes:   | Status Code | Description                  | Notes                                                                  |  |-------------|------------------------------|------------------------------------------------------------------------|  | 200         | OK                           | The request was successful.                                             |  | 201         | Created                      | The object was successfully created. This code is only used with objects that are created immediately.  | 204         | No content                   | A successful request, but there is no additional information to send back in the response body (in a case when the object was deleted).    | 400         | Bad Request                  | The request could not be understood by the server. Incoming parameters might not be valid. |  | 401         | Unauthorized            | The client is unauthenticated. The client must authenticate itself to get the requested response. |  | 403         | Forbidden                   | The client does not have access rights to the content.  | 404         | Not Found                    | The requested resource is not found.                                    |  | 409         | Conflict | This response is sent when a request conflicts with the current state of the object (e.g. deleting the security group with the compute instances in it).|  | 422         | Unprocessable Content   | The request was well-formed but was unable to be followed due to incorrect field values (e.g. creation of a virtual machine in the non-existent data center).  |  | 500         | Internal server Error                 | The server could not return the representation due to an internal server error. | 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VmResizeCompute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmResizeCompute{}

// VmResizeCompute struct for VmResizeCompute
type VmResizeCompute struct {
	// Action with a virtual machine.  **Note:** To retrieve available configurations for DigitalOcean system volumes, use the `edithardware` action 
	Action string `json:"action"`
	// Number of virtual Central Processing Units (vCPUs)
	VCpu int32 `json:"vCpu"`
	// Type of virtual Central Processing Units (vCPUs)
	VCpuType string `json:"vCpuType"`
	// Capacity of the RAM in gigabytes
	RamGb int32 `json:"ramGb"`
}

type _VmResizeCompute VmResizeCompute

// NewVmResizeCompute instantiates a new VmResizeCompute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmResizeCompute(action string, vCpu int32, vCpuType string, ramGb int32) *VmResizeCompute {
	this := VmResizeCompute{}
	this.Action = action
	this.VCpu = vCpu
	this.VCpuType = vCpuType
	this.RamGb = ramGb
	return &this
}

// NewVmResizeComputeWithDefaults instantiates a new VmResizeCompute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmResizeComputeWithDefaults() *VmResizeCompute {
	this := VmResizeCompute{}
	return &this
}

// GetAction returns the Action field value
func (o *VmResizeCompute) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *VmResizeCompute) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *VmResizeCompute) SetAction(v string) {
	o.Action = v
}

// GetVCpu returns the VCpu field value
func (o *VmResizeCompute) GetVCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VCpu
}

// GetVCpuOk returns a tuple with the VCpu field value
// and a boolean to check if the value has been set.
func (o *VmResizeCompute) GetVCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VCpu, true
}

// SetVCpu sets field value
func (o *VmResizeCompute) SetVCpu(v int32) {
	o.VCpu = v
}

// GetVCpuType returns the VCpuType field value
func (o *VmResizeCompute) GetVCpuType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VCpuType
}

// GetVCpuTypeOk returns a tuple with the VCpuType field value
// and a boolean to check if the value has been set.
func (o *VmResizeCompute) GetVCpuTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VCpuType, true
}

// SetVCpuType sets field value
func (o *VmResizeCompute) SetVCpuType(v string) {
	o.VCpuType = v
}

// GetRamGb returns the RamGb field value
func (o *VmResizeCompute) GetRamGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RamGb
}

// GetRamGbOk returns a tuple with the RamGb field value
// and a boolean to check if the value has been set.
func (o *VmResizeCompute) GetRamGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamGb, true
}

// SetRamGb sets field value
func (o *VmResizeCompute) SetRamGb(v int32) {
	o.RamGb = v
}

func (o VmResizeCompute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmResizeCompute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["vCpu"] = o.VCpu
	toSerialize["vCpuType"] = o.VCpuType
	toSerialize["ramGb"] = o.RamGb
	return toSerialize, nil
}

func (o *VmResizeCompute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"vCpu",
		"vCpuType",
		"ramGb",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVmResizeCompute := _VmResizeCompute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVmResizeCompute)

	if err != nil {
		return err
	}

	*o = VmResizeCompute(varVmResizeCompute)

	return err
}

type NullableVmResizeCompute struct {
	value *VmResizeCompute
	isSet bool
}

func (v NullableVmResizeCompute) Get() *VmResizeCompute {
	return v.value
}

func (v *NullableVmResizeCompute) Set(val *VmResizeCompute) {
	v.value = val
	v.isSet = true
}

func (v NullableVmResizeCompute) IsSet() bool {
	return v.isSet
}

func (v *NullableVmResizeCompute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmResizeCompute(val *VmResizeCompute) *NullableVmResizeCompute {
	return &NullableVmResizeCompute{value: val, isSet: true}
}

func (v NullableVmResizeCompute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmResizeCompute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


