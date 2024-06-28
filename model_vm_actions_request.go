/*
Public EMMA API

### About Infrastructure API   **Base URL:** **<u>https://api.emma.ms/external</u>**   This **Infrastructure API** is for managing the emma cloud infrastructure within a project. The API enables you to view, create, edit, and delete _Virtual machines, Spot instances, Applications, Kubernetes clusters, SSH keys, Security groups, and Backup policies_. For creating the resources you can use the endpoints with the dictionaries: _Data centers, Locations, Providers, Operating systems, Virtual machines configurations, Spot instances configurations, Kubernetes clusters configurations._   ### Authentication   #### 1. Create service application   To access the API, enter your project, navigate to **Settings** > **Service Apps**, and create a service application. Select the access level **Read**, **Operate**, or **Manage**.   - **Read** - only GET methods are allowed in the API.   - **Operate** - some operations are allowed with the resources (e.g. _Start, Reboot,_ and _Shutdown_ of the Virtual machines).   - **Manage** - full creating, updating, and deleting of the resources is allowed.     #### 2. Get access token   - Copy the **Client ID** and **Client Secret** in the service application.  - Send an API request to the endpoint **_/issue-token** as specified in the **Authentication** section of the API documentation. You will receive access and refresh tokens in the response.   _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token \\  -H \"Content-Type: application/json\" \\  -d '{\"clientId\": \"YOUR-CLIENT-ID\", \"clientSecret\": \"YOUR-CLIENT-SECRET\"}'  ```  _For Windows:_  ```  curl -X POST https://api.emma.ms/external/v1/issue-token ^  -H \"Content-Type: application/json\" ^  -d \"{\\\"clientId\\\": \\\"YOUR-CLIENT-ID\\\", \\\"clientSecret\\\": \\\"YOUR-CLIENT-SECRET\\\"}\"  ```      #### 3. Use access token in requests  The Bearer access token is a text string, included in the request header, for example:   _For Linux / Mac:_  ```  curl -X GET https://api.emma.ms/external/v1/locations -H \"Authorization: Bearer YOUR-ACCESS-TOKEN-HERE\"  ```   Use this token for the API requests.     #### 4. Refresh token  The access token will expire in 10 minutes. A new access token may be created using the refresh token (without Client ID and Client Secret).   To get a new access token send a request to the **_/refresh-token** endpoint:    _For Linux / Mac:_  ```  curl -X POST https://api.emma.ms/external/v1/refresh-token \\  -H \"Content-Type: application/json\" \\  -H \"Authorization: Bearer YOUR-ACCESS-TOKEN\" \\  -d '{\"refreshToken\": \"YOUR-REFRESH-TOKEN\"}'  ```       ### Possible response status codes   We use standard HTTP response codes to show the success or failure of requests.   `2xx` - successful responses.   `4xx` - client error responses (the response contains an explanation of the error).   `5xx` - server error responses.   The API uses the following status codes:   | Status Code | Description                  | Notes                                                                  |  |-------------|------------------------------|------------------------------------------------------------------------|  | 200         | OK                           | The request was successful.                                             |  | 201         | Created                      | The object was successfully created. This code is only used with objects that are created immediately.  | 400         | Bad Request                  | The request could not be understood by the server. Incoming parameters might not be valid. |  | 401         | Unauthorized            | The client is unauthenticated. The client must authenticate itself to get the requested response. |  | 403         | Forbidden                   | The client does not have access rights to the content.  | 404         | Not Found                    | The requested resource is not found.                                    |  | 409         | Conflict | This response is sent when a request conflicts with the current state of the object (e.g. deleting the security group with the compute instances in it).|  | 422         | Unprocessable Content   | The request was well-formed but was unable to be followed due to incorrect field values (e.g. creation of a virtual machine in the non-existent data center).  |  | 500         | Internal server Error                 | The server could not return the representation due to an internal server error. |

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emma

import (
	"encoding/json"
	"fmt"
)

// VmActionsRequest - struct for VmActionsRequest
type VmActionsRequest struct {
	VmClone        *VmClone
	VmEditHardware *VmEditHardware
	VmReboot       *VmReboot
	VmShutdown     *VmShutdown
	VmStart        *VmStart
	VmTransfer     *VmTransfer
}

// VmCloneAsVmActionsRequest is a convenience function that returns VmClone wrapped in VmActionsRequest
func VmCloneAsVmActionsRequest(v *VmClone) VmActionsRequest {
	return VmActionsRequest{
		VmClone: v,
	}
}

// VmEditHardwareAsVmActionsRequest is a convenience function that returns VmEditHardware wrapped in VmActionsRequest
func VmEditHardwareAsVmActionsRequest(v *VmEditHardware) VmActionsRequest {
	return VmActionsRequest{
		VmEditHardware: v,
	}
}

// VmRebootAsVmActionsRequest is a convenience function that returns VmReboot wrapped in VmActionsRequest
func VmRebootAsVmActionsRequest(v *VmReboot) VmActionsRequest {
	return VmActionsRequest{
		VmReboot: v,
	}
}

// VmShutdownAsVmActionsRequest is a convenience function that returns VmShutdown wrapped in VmActionsRequest
func VmShutdownAsVmActionsRequest(v *VmShutdown) VmActionsRequest {
	return VmActionsRequest{
		VmShutdown: v,
	}
}

// VmStartAsVmActionsRequest is a convenience function that returns VmStart wrapped in VmActionsRequest
func VmStartAsVmActionsRequest(v *VmStart) VmActionsRequest {
	return VmActionsRequest{
		VmStart: v,
	}
}

// VmTransferAsVmActionsRequest is a convenience function that returns VmTransfer wrapped in VmActionsRequest
func VmTransferAsVmActionsRequest(v *VmTransfer) VmActionsRequest {
	return VmActionsRequest{
		VmTransfer: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VmActionsRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into VmClone
	err = newStrictDecoder(data).Decode(&dst.VmClone)
	if err == nil {
		jsonVmClone, _ := json.Marshal(dst.VmClone)
		if string(jsonVmClone) == "{}" { // empty struct
			dst.VmClone = nil
		} else {
			match++
		}
	} else {
		dst.VmClone = nil
	}

	// try to unmarshal data into VmEditHardware
	err = newStrictDecoder(data).Decode(&dst.VmEditHardware)
	if err == nil {
		jsonVmEditHardware, _ := json.Marshal(dst.VmEditHardware)
		if string(jsonVmEditHardware) == "{}" { // empty struct
			dst.VmEditHardware = nil
		} else {
			match++
		}
	} else {
		dst.VmEditHardware = nil
	}

	// try to unmarshal data into VmReboot
	err = newStrictDecoder(data).Decode(&dst.VmReboot)
	if err == nil {
		jsonVmReboot, _ := json.Marshal(dst.VmReboot)
		if string(jsonVmReboot) == "{}" { // empty struct
			dst.VmReboot = nil
		} else {
			match++
		}
	} else {
		dst.VmReboot = nil
	}

	// try to unmarshal data into VmShutdown
	err = newStrictDecoder(data).Decode(&dst.VmShutdown)
	if err == nil {
		jsonVmShutdown, _ := json.Marshal(dst.VmShutdown)
		if string(jsonVmShutdown) == "{}" { // empty struct
			dst.VmShutdown = nil
		} else {
			match++
		}
	} else {
		dst.VmShutdown = nil
	}

	// try to unmarshal data into VmStart
	err = newStrictDecoder(data).Decode(&dst.VmStart)
	if err == nil {
		jsonVmStart, _ := json.Marshal(dst.VmStart)
		if string(jsonVmStart) == "{}" { // empty struct
			dst.VmStart = nil
		} else {
			match++
		}
	} else {
		dst.VmStart = nil
	}

	// try to unmarshal data into VmTransfer
	err = newStrictDecoder(data).Decode(&dst.VmTransfer)
	if err == nil {
		jsonVmTransfer, _ := json.Marshal(dst.VmTransfer)
		if string(jsonVmTransfer) == "{}" { // empty struct
			dst.VmTransfer = nil
		} else {
			match++
		}
	} else {
		dst.VmTransfer = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.VmClone = nil
		dst.VmEditHardware = nil
		dst.VmReboot = nil
		dst.VmShutdown = nil
		dst.VmStart = nil
		dst.VmTransfer = nil

		return fmt.Errorf("data matches more than one schema in oneOf(VmActionsRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(VmActionsRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VmActionsRequest) MarshalJSON() ([]byte, error) {
	if src.VmClone != nil {
		return json.Marshal(&src.VmClone)
	}

	if src.VmEditHardware != nil {
		return json.Marshal(&src.VmEditHardware)
	}

	if src.VmReboot != nil {
		return json.Marshal(&src.VmReboot)
	}

	if src.VmShutdown != nil {
		return json.Marshal(&src.VmShutdown)
	}

	if src.VmStart != nil {
		return json.Marshal(&src.VmStart)
	}

	if src.VmTransfer != nil {
		return json.Marshal(&src.VmTransfer)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VmActionsRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.VmClone != nil {
		return obj.VmClone
	}

	if obj.VmEditHardware != nil {
		return obj.VmEditHardware
	}

	if obj.VmReboot != nil {
		return obj.VmReboot
	}

	if obj.VmShutdown != nil {
		return obj.VmShutdown
	}

	if obj.VmStart != nil {
		return obj.VmStart
	}

	if obj.VmTransfer != nil {
		return obj.VmTransfer
	}

	// all schemas are nil
	return nil
}

type NullableVmActionsRequest struct {
	value *VmActionsRequest
	isSet bool
}

func (v NullableVmActionsRequest) Get() *VmActionsRequest {
	return v.value
}

func (v *NullableVmActionsRequest) Set(val *VmActionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVmActionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVmActionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmActionsRequest(val *VmActionsRequest) *NullableVmActionsRequest {
	return &NullableVmActionsRequest{value: val, isSet: true}
}

func (v NullableVmActionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmActionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
