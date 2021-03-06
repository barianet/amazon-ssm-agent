// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package appconfig manages the configuration of the agent.
package appconfig

// CredentialProfile represents configurations for aws credential profile
type CredentialProfile struct {
	Path         string
	Name         string
	ShareCreds   bool
	ShareProfile string
}

// MdsCfg represents configuration for Message delivery service (MDS)
type MdsCfg struct {
	Endpoint            string
	CommandWorkersLimit int
	StopTimeoutMillis   int64
	CommandRetryLimit   int
}

// SsmCfg represents configuration for Simple system manager (SSM)
type SsmCfg struct {
	Endpoint                    string
	HealthFrequencyMinutes      int
	AssociationFrequencyMinutes int
	AssociationRetryLimit       int
	// TODO: test hook, can be removed before release
	// this is to skip ssl verification for the beta self signed certs
	InsecureSkipVerify             bool
	CustomInventoryDefaultLocation string
}

// AgentInfo represents metadata for amazon-ssm-agent
type AgentInfo struct {
	Name                 string
	Version              string
	Region               string
	OrchestrationRootDir string
	DownloadRootDir      string
}

// OsInfo represents os related information
type OsInfo struct {
	Lang    string
	Name    string
	Version string
}

// S3Cfg represents configurations related to S3 bucket and key for SSM
type S3Cfg struct {
	Region    string
	LogBucket string
	LogKey    string
}

// SsmagentConfig stores agent configuration values.
type SsmagentConfig struct {
	Profile CredentialProfile
	Mds     MdsCfg
	Ssm     SsmCfg
	Agent   AgentInfo
	Os      OsInfo
	S3      S3Cfg
}
