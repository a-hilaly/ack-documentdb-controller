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

package db_instance

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customPreCompare(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade) {
		delta.Add("Spec.AutoMinorVersionUpgrade", a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade)
	} else if a.ko.Spec.AutoMinorVersionUpgrade != nil && b.ko.Spec.AutoMinorVersionUpgrade != nil {
		if *a.ko.Spec.AutoMinorVersionUpgrade != *b.ko.Spec.AutoMinorVersionUpgrade {
			delta.Add("Spec.AutoMinorVersionUpgrade", a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone) {
		delta.Add("Spec.AvailabilityZone", a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone)
	} else if a.ko.Spec.AvailabilityZone != nil && b.ko.Spec.AvailabilityZone != nil {
		if *a.ko.Spec.AvailabilityZone != *b.ko.Spec.AvailabilityZone {
			delta.Add("Spec.AvailabilityZone", a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CACertificateIdentifier, b.ko.Spec.CACertificateIdentifier) {
		delta.Add("Spec.CACertificateIdentifier", a.ko.Spec.CACertificateIdentifier, b.ko.Spec.CACertificateIdentifier)
	} else if a.ko.Spec.CACertificateIdentifier != nil && b.ko.Spec.CACertificateIdentifier != nil {
		if *a.ko.Spec.CACertificateIdentifier != *b.ko.Spec.CACertificateIdentifier {
			delta.Add("Spec.CACertificateIdentifier", a.ko.Spec.CACertificateIdentifier, b.ko.Spec.CACertificateIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot) {
		delta.Add("Spec.CopyTagsToSnapshot", a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot)
	} else if a.ko.Spec.CopyTagsToSnapshot != nil && b.ko.Spec.CopyTagsToSnapshot != nil {
		if *a.ko.Spec.CopyTagsToSnapshot != *b.ko.Spec.CopyTagsToSnapshot {
			delta.Add("Spec.CopyTagsToSnapshot", a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier) {
		delta.Add("Spec.DBClusterIdentifier", a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier)
	} else if a.ko.Spec.DBClusterIdentifier != nil && b.ko.Spec.DBClusterIdentifier != nil {
		if *a.ko.Spec.DBClusterIdentifier != *b.ko.Spec.DBClusterIdentifier {
			delta.Add("Spec.DBClusterIdentifier", a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBInstanceClass, b.ko.Spec.DBInstanceClass) {
		delta.Add("Spec.DBInstanceClass", a.ko.Spec.DBInstanceClass, b.ko.Spec.DBInstanceClass)
	} else if a.ko.Spec.DBInstanceClass != nil && b.ko.Spec.DBInstanceClass != nil {
		if *a.ko.Spec.DBInstanceClass != *b.ko.Spec.DBInstanceClass {
			delta.Add("Spec.DBInstanceClass", a.ko.Spec.DBInstanceClass, b.ko.Spec.DBInstanceClass)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBInstanceIdentifier, b.ko.Spec.DBInstanceIdentifier) {
		delta.Add("Spec.DBInstanceIdentifier", a.ko.Spec.DBInstanceIdentifier, b.ko.Spec.DBInstanceIdentifier)
	} else if a.ko.Spec.DBInstanceIdentifier != nil && b.ko.Spec.DBInstanceIdentifier != nil {
		if *a.ko.Spec.DBInstanceIdentifier != *b.ko.Spec.DBInstanceIdentifier {
			delta.Add("Spec.DBInstanceIdentifier", a.ko.Spec.DBInstanceIdentifier, b.ko.Spec.DBInstanceIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Engine, b.ko.Spec.Engine) {
		delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
	} else if a.ko.Spec.Engine != nil && b.ko.Spec.Engine != nil {
		if *a.ko.Spec.Engine != *b.ko.Spec.Engine {
			delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PerformanceInsightsEnabled, b.ko.Spec.PerformanceInsightsEnabled) {
		delta.Add("Spec.PerformanceInsightsEnabled", a.ko.Spec.PerformanceInsightsEnabled, b.ko.Spec.PerformanceInsightsEnabled)
	} else if a.ko.Spec.PerformanceInsightsEnabled != nil && b.ko.Spec.PerformanceInsightsEnabled != nil {
		if *a.ko.Spec.PerformanceInsightsEnabled != *b.ko.Spec.PerformanceInsightsEnabled {
			delta.Add("Spec.PerformanceInsightsEnabled", a.ko.Spec.PerformanceInsightsEnabled, b.ko.Spec.PerformanceInsightsEnabled)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PerformanceInsightsKMSKeyID, b.ko.Spec.PerformanceInsightsKMSKeyID) {
		delta.Add("Spec.PerformanceInsightsKMSKeyID", a.ko.Spec.PerformanceInsightsKMSKeyID, b.ko.Spec.PerformanceInsightsKMSKeyID)
	} else if a.ko.Spec.PerformanceInsightsKMSKeyID != nil && b.ko.Spec.PerformanceInsightsKMSKeyID != nil {
		if *a.ko.Spec.PerformanceInsightsKMSKeyID != *b.ko.Spec.PerformanceInsightsKMSKeyID {
			delta.Add("Spec.PerformanceInsightsKMSKeyID", a.ko.Spec.PerformanceInsightsKMSKeyID, b.ko.Spec.PerformanceInsightsKMSKeyID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.PerformanceInsightsKMSKeyRef, b.ko.Spec.PerformanceInsightsKMSKeyRef) {
		delta.Add("Spec.PerformanceInsightsKMSKeyRef", a.ko.Spec.PerformanceInsightsKMSKeyRef, b.ko.Spec.PerformanceInsightsKMSKeyRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow) {
		delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
	} else if a.ko.Spec.PreferredMaintenanceWindow != nil && b.ko.Spec.PreferredMaintenanceWindow != nil {
		if *a.ko.Spec.PreferredMaintenanceWindow != *b.ko.Spec.PreferredMaintenanceWindow {
			delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PromotionTier, b.ko.Spec.PromotionTier) {
		delta.Add("Spec.PromotionTier", a.ko.Spec.PromotionTier, b.ko.Spec.PromotionTier)
	} else if a.ko.Spec.PromotionTier != nil && b.ko.Spec.PromotionTier != nil {
		if *a.ko.Spec.PromotionTier != *b.ko.Spec.PromotionTier {
			delta.Add("Spec.PromotionTier", a.ko.Spec.PromotionTier, b.ko.Spec.PromotionTier)
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
