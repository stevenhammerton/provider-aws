/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type Domain string

const (
	Domain_S3  Domain = "S3"
	Domain_EFS Domain = "EFS"
)

type EndpointType string

const (
	EndpointType_PUBLIC       EndpointType = "PUBLIC"
	EndpointType_VPC          EndpointType = "VPC"
	EndpointType_VPC_ENDPOINT EndpointType = "VPC_ENDPOINT"
)

type HomeDirectoryType string

const (
	HomeDirectoryType_PATH    HomeDirectoryType = "PATH"
	HomeDirectoryType_LOGICAL HomeDirectoryType = "LOGICAL"
)

type IDentityProviderType string

const (
	IDentityProviderType_SERVICE_MANAGED IDentityProviderType = "SERVICE_MANAGED"
	IDentityProviderType_API_GATEWAY     IDentityProviderType = "API_GATEWAY"
)

type Protocol string

const (
	Protocol_SFTP Protocol = "SFTP"
	Protocol_FTP  Protocol = "FTP"
	Protocol_FTPS Protocol = "FTPS"
)

type State string

const (
	State_OFFLINE      State = "OFFLINE"
	State_ONLINE       State = "ONLINE"
	State_STARTING     State = "STARTING"
	State_STOPPING     State = "STOPPING"
	State_START_FAILED State = "START_FAILED"
	State_STOP_FAILED  State = "STOP_FAILED"
)