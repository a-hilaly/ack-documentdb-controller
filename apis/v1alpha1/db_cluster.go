// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DBClusterSpec defines the desired state of DBCluster.
//
// Detailed information about a cluster.
type DBClusterSpec struct {

	// A list of Amazon EC2 Availability Zones that instances in the cluster can
	// be created in.
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`
	// The number of days for which automated backups are retained. You must specify
	// a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	//   - Must be a value from 1 to 35.
	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty"`
	// The cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster
	// +kubebuilder:validation:Required
	DBClusterIdentifier *string `json:"dbClusterIdentifier"`
	// The name of the cluster parameter group to associate with this cluster.
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty"`
	// A subnet group to associate with this cluster.
	//
	// Constraints: Must match the name of an existing DBSubnetGroup. Must not be
	// default.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty"`
	// Specifies whether this cluster can be deleted. If DeletionProtection is enabled,
	// the cluster cannot be deleted unless it is modified and DeletionProtection
	// is disabled. DeletionProtection protects clusters from being accidentally
	// deleted.
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
	// DestinationRegion is used for presigning the request to a given region.
	DestinationRegion *string `json:"destinationRegion,omitempty"`
	// A list of log types that need to be enabled for exporting to Amazon CloudWatch
	// Logs. You can enable audit logs or profiler logs. For more information, see
	// Auditing Amazon DocumentDB Events (https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html)
	// and Profiling Amazon DocumentDB Operations (https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html).
	EnableCloudwatchLogsExports []*string `json:"enableCloudwatchLogsExports,omitempty"`
	// The name of the database engine to be used for this cluster.
	//
	// Valid values: docdb
	// +kubebuilder:validation:Required
	Engine *string `json:"engine"`
	// The version number of the database engine to use. The --engine-version will
	// default to the latest major engine version. For production workloads, we
	// recommend explicitly declaring this parameter with the intended major engine
	// version.
	EngineVersion *string `json:"engineVersion,omitempty"`
	// The cluster identifier of the new global cluster.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty"`
	// The KMS key identifier for an encrypted cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a cluster using the same Amazon Web Services account
	// that owns the KMS encryption key that is used to encrypt the new cluster,
	// you can use the KMS key alias instead of the ARN for the KMS encryption key.
	//
	// If an encryption key is not specified in KmsKeyId:
	//
	//   - If the StorageEncrypted parameter is true, Amazon DocumentDB uses your
	//     default encryption key.
	//
	// KMS creates the default encryption key for your Amazon Web Services account.
	// Your Amazon Web Services account has a different default encryption key for
	// each Amazon Web Services Regions.
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// The password for the master database user. This password can contain any
	// printable ASCII character except forward slash (/), double quote ("), or
	// the "at" symbol (@).
	//
	// Constraints: Must contain from 8 to 100 characters.
	MasterUserPassword *ackv1alpha1.SecretKeyReference `json:"masterUserPassword,omitempty"`
	// The name of the master user for the cluster.
	//
	// Constraints:
	//
	//   - Must be from 1 to 63 letters or numbers.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot be a reserved word for the chosen database engine.
	MasterUsername *string `json:"masterUsername,omitempty"`
	// The port number on which the instances in the cluster accept connections.
	Port *int64 `json:"port,omitempty"`
	// Not currently supported.
	PreSignedURL *string `json:"preSignedURL,omitempty"`
	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each Amazon Web Services Region.
	//
	// Constraints:
	//
	//   - Must be in the format hh24:mi-hh24:mi.
	//
	//   - Must be in Universal Coordinated Time (UTC).
	//
	//   - Must not conflict with the preferred maintenance window.
	//
	//   - Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty"`
	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each Amazon Web Services Region, occurring on a random day of
	// the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty"`
	// SourceRegion is the source region where the resource exists. This is not
	// sent over the wire and is only used for presigning. This value should always
	// have the same region as the source ARN.
	SourceRegion *string `json:"sourceRegion,omitempty"`
	// Specifies whether the cluster is encrypted.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty"`
	// The storage type to associate with the DB cluster.
	//
	// For information on storage types for Amazon DocumentDB clusters, see Cluster
	// storage configurations in the Amazon DocumentDB Developer Guide.
	//
	// # Valid values for storage type - standard | iopt1
	//
	// # Default value is standard
	//
	// When you create a DocumentDB DB cluster with the storage type set to iopt1,
	// the storage type is returned in the response. The storage type isn't returned
	// when you set it to standard.
	StorageType *string `json:"storageType,omitempty"`
	// The tags to be assigned to the cluster.
	Tags []*Tag `json:"tags,omitempty"`
	// A list of EC2 VPC security groups to associate with this cluster.
	VPCSecurityGroupIDs []*string `json:"vpcSecurityGroupIDs,omitempty"`
}

// DBClusterStatus defines the observed state of DBCluster
type DBClusterStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Provides a list of the Identity and Access Management (IAM) roles that are
	// associated with the cluster. (IAM) roles that are associated with a cluster
	// grant permission for the cluster to access other Amazon Web Services services
	// on your behalf.
	// +kubebuilder:validation:Optional
	AssociatedRoles []*DBClusterRole `json:"associatedRoles,omitempty"`
	// Identifies the clone group to which the DB cluster is associated.
	// +kubebuilder:validation:Optional
	CloneGroupID *string `json:"cloneGroupID,omitempty"`
	// Specifies the time when the cluster was created, in Universal Coordinated
	// Time (UTC).
	// +kubebuilder:validation:Optional
	ClusterCreateTime *metav1.Time `json:"clusterCreateTime,omitempty"`
	// Provides the list of instances that make up the cluster.
	// +kubebuilder:validation:Optional
	DBClusterMembers []*DBClusterMember `json:"dbClusterMembers,omitempty"`
	// Specifies the name of the cluster parameter group for the cluster.
	// +kubebuilder:validation:Optional
	DBClusterParameterGroup *string `json:"dbClusterParameterGroup,omitempty"`
	// Specifies information on the subnet group that is associated with the cluster,
	// including the name, description, and subnets in the subnet group.
	// +kubebuilder:validation:Optional
	DBSubnetGroup *string `json:"dbSubnetGroup,omitempty"`
	// The Amazon Web Services Region-unique, immutable identifier for the cluster.
	// This identifier is found in CloudTrail log entries whenever the KMS key for
	// the cluster is accessed.
	// +kubebuilder:validation:Optional
	DBClusterResourceID *string `json:"dbClusterResourceID,omitempty"`
	// The earliest time to which a database can be restored with point-in-time
	// restore.
	// +kubebuilder:validation:Optional
	EarliestRestorableTime *metav1.Time `json:"earliestRestorableTime,omitempty"`
	// A list of log types that this cluster is configured to export to Amazon CloudWatch
	// Logs.
	// +kubebuilder:validation:Optional
	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty"`
	// Specifies the connection endpoint for the primary instance of the cluster.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty"`
	// Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.
	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneID,omitempty"`
	// Specifies the latest time to which a database can be restored with point-in-time
	// restore.
	// +kubebuilder:validation:Optional
	LatestRestorableTime *metav1.Time `json:"latestRestorableTime,omitempty"`
	// Specifies whether the cluster has instances in multiple Availability Zones.
	// +kubebuilder:validation:Optional
	MultiAZ *bool `json:"multiAZ,omitempty"`
	// Specifies the progress of the operation as a percentage.
	// +kubebuilder:validation:Optional
	PercentProgress *string `json:"percentProgress,omitempty"`
	// Contains one or more identifiers of the secondary clusters that are associated
	// with this cluster.
	// +kubebuilder:validation:Optional
	ReadReplicaIdentifiers []*string `json:"readReplicaIdentifiers,omitempty"`
	// The reader endpoint for the cluster. The reader endpoint for a cluster load
	// balances connections across the Amazon DocumentDB replicas that are available
	// in a cluster. As clients request new connections to the reader endpoint,
	// Amazon DocumentDB distributes the connection requests among the Amazon DocumentDB
	// replicas in the cluster. This functionality can help balance your read workload
	// across multiple Amazon DocumentDB replicas in your cluster.
	//
	// If a failover occurs, and the Amazon DocumentDB replica that you are connected
	// to is promoted to be the primary instance, your connection is dropped. To
	// continue sending your read workload to other Amazon DocumentDB replicas in
	// the cluster, you can then reconnect to the reader endpoint.
	// +kubebuilder:validation:Optional
	ReaderEndpoint *string `json:"readerEndpoint,omitempty"`
	// Contains the identifier of the source cluster if this cluster is a secondary
	// cluster.
	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier,omitempty"`
	// Specifies the current state of this cluster.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
	// Provides a list of virtual private cloud (VPC) security groups that the cluster
	// belongs to.
	// +kubebuilder:validation:Optional
	VPCSecurityGroups []*VPCSecurityGroupMembership `json:"vpcSecurityGroups,omitempty"`
}

// DBCluster is the Schema for the DBClusters API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type DBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBClusterSpec   `json:"spec,omitempty"`
	Status            DBClusterStatus `json:"status,omitempty"`
}

// DBClusterList contains a list of DBCluster
// +kubebuilder:object:root=true
type DBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBCluster{}, &DBClusterList{})
}
