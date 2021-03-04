// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	klog "k8s.io/klog"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucketControllerConfiguration) DeepCopyInto(out *BackupBucketControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucketControllerConfiguration.
func (in *BackupBucketControllerConfiguration) DeepCopy() *BackupBucketControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(BackupBucketControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEntryControllerConfiguration) DeepCopyInto(out *BackupEntryControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.DeletionGracePeriodHours != nil {
		in, out := &in.DeletionGracePeriodHours, &out.DeletionGracePeriodHours
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEntryControllerConfiguration.
func (in *BackupEntryControllerConfiguration) DeepCopy() *BackupEntryControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(BackupEntryControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionThreshold) DeepCopyInto(out *ConditionThreshold) {
	*out = *in
	out.Duration = in.Duration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionThreshold.
func (in *ConditionThreshold) DeepCopy() *ConditionThreshold {
	if in == nil {
		return nil
	}
	out := new(ConditionThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationCareControllerConfiguration) DeepCopyInto(out *ControllerInstallationCareControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationCareControllerConfiguration.
func (in *ControllerInstallationCareControllerConfiguration) DeepCopy() *ControllerInstallationCareControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationCareControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationControllerConfiguration) DeepCopyInto(out *ControllerInstallationControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationControllerConfiguration.
func (in *ControllerInstallationControllerConfiguration) DeepCopy() *ControllerInstallationControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationRequiredControllerConfiguration) DeepCopyInto(out *ControllerInstallationRequiredControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationRequiredControllerConfiguration.
func (in *ControllerInstallationRequiredControllerConfiguration) DeepCopy() *ControllerInstallationRequiredControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationRequiredControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentBit) DeepCopyInto(out *FluentBit) {
	*out = *in
	if in.ServiceSection != nil {
		in, out := &in.ServiceSection, &out.ServiceSection
		*out = new(string)
		**out = **in
	}
	if in.InputSection != nil {
		in, out := &in.InputSection, &out.InputSection
		*out = new(string)
		**out = **in
	}
	if in.OutputSection != nil {
		in, out := &in.OutputSection, &out.OutputSection
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentBit.
func (in *FluentBit) DeepCopy() *FluentBit {
	if in == nil {
		return nil
	}
	out := new(FluentBit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenClientConnection) DeepCopyInto(out *GardenClientConnection) {
	*out = *in
	out.ClientConnectionConfiguration = in.ClientConnectionConfiguration
	if in.GardenClusterAddress != nil {
		in, out := &in.GardenClusterAddress, &out.GardenClusterAddress
		*out = new(string)
		**out = **in
	}
	if in.GardenClusterCACert != nil {
		in, out := &in.GardenClusterCACert, &out.GardenClusterCACert
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.BootstrapKubeconfig != nil {
		in, out := &in.BootstrapKubeconfig, &out.BootstrapKubeconfig
		*out = new(corev1.SecretReference)
		**out = **in
	}
	if in.KubeconfigSecret != nil {
		in, out := &in.KubeconfigSecret, &out.KubeconfigSecret
		*out = new(corev1.SecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenClientConnection.
func (in *GardenClientConnection) DeepCopy() *GardenClientConnection {
	if in == nil {
		return nil
	}
	out := new(GardenClientConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenletConfiguration) DeepCopyInto(out *GardenletConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.GardenClientConnection != nil {
		in, out := &in.GardenClientConnection, &out.GardenClientConnection
		*out = new(GardenClientConnection)
		(*in).DeepCopyInto(*out)
	}
	if in.SeedClientConnection != nil {
		in, out := &in.SeedClientConnection, &out.SeedClientConnection
		*out = new(SeedClientConnection)
		**out = **in
	}
	if in.ShootClientConnection != nil {
		in, out := &in.ShootClientConnection, &out.ShootClientConnection
		*out = new(ShootClientConnection)
		**out = **in
	}
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = new(GardenletControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourcesConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(LeaderElectionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.KubernetesLogLevel != nil {
		in, out := &in.KubernetesLogLevel, &out.KubernetesLogLevel
		*out = new(klog.Level)
		**out = **in
	}
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(ServerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SeedConfig != nil {
		in, out := &in.SeedConfig, &out.SeedConfig
		*out = new(SeedConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SeedSelector != nil {
		in, out := &in.SeedSelector, &out.SeedSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(Logging)
		(*in).DeepCopyInto(*out)
	}
	if in.SNI != nil {
		in, out := &in.SNI, &out.SNI
		*out = new(SNI)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenletConfiguration.
func (in *GardenletConfiguration) DeepCopy() *GardenletConfiguration {
	if in == nil {
		return nil
	}
	out := new(GardenletConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GardenletConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenletControllerConfiguration) DeepCopyInto(out *GardenletControllerConfiguration) {
	*out = *in
	if in.BackupBucket != nil {
		in, out := &in.BackupBucket, &out.BackupBucket
		*out = new(BackupBucketControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupEntry != nil {
		in, out := &in.BackupEntry, &out.BackupEntry
		*out = new(BackupEntryControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerInstallation != nil {
		in, out := &in.ControllerInstallation, &out.ControllerInstallation
		*out = new(ControllerInstallationControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerInstallationCare != nil {
		in, out := &in.ControllerInstallationCare, &out.ControllerInstallationCare
		*out = new(ControllerInstallationCareControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerInstallationRequired != nil {
		in, out := &in.ControllerInstallationRequired, &out.ControllerInstallationRequired
		*out = new(ControllerInstallationRequiredControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Seed != nil {
		in, out := &in.Seed, &out.Seed
		*out = new(SeedControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Shoot != nil {
		in, out := &in.Shoot, &out.Shoot
		*out = new(ShootControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ShootCare != nil {
		in, out := &in.ShootCare, &out.ShootCare
		*out = new(ShootCareControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ShootStateSync != nil {
		in, out := &in.ShootStateSync, &out.ShootStateSync
		*out = new(ShootStateSyncControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ShootedSeedRegistration != nil {
		in, out := &in.ShootedSeedRegistration, &out.ShootedSeedRegistration
		*out = new(ShootedSeedRegistrationControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.SeedAPIServerNetworkPolicy != nil {
		in, out := &in.SeedAPIServerNetworkPolicy, &out.SeedAPIServerNetworkPolicy
		*out = new(SeedAPIServerNetworkPolicyControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenletControllerConfiguration.
func (in *GardenletControllerConfiguration) DeepCopy() *GardenletControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(GardenletControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSServer) DeepCopyInto(out *HTTPSServer) {
	*out = *in
	out.Server = in.Server
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSServer)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSServer.
func (in *HTTPSServer) DeepCopy() *HTTPSServer {
	if in == nil {
		return nil
	}
	out := new(HTTPSServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	out.LeaderElectionConfiguration = in.LeaderElectionConfiguration
	if in.LockObjectNamespace != nil {
		in, out := &in.LockObjectNamespace, &out.LockObjectNamespace
		*out = new(string)
		**out = **in
	}
	if in.LockObjectName != nil {
		in, out := &in.LockObjectName, &out.LockObjectName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logging) DeepCopyInto(out *Logging) {
	*out = *in
	if in.FluentBit != nil {
		in, out := &in.FluentBit, &out.FluentBit
		*out = new(FluentBit)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logging.
func (in *Logging) DeepCopy() *Logging {
	if in == nil {
		return nil
	}
	out := new(Logging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesConfiguration) DeepCopyInto(out *ResourcesConfiguration) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Reserved != nil {
		in, out := &in.Reserved, &out.Reserved
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesConfiguration.
func (in *ResourcesConfiguration) DeepCopy() *ResourcesConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourcesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNI) DeepCopyInto(out *SNI) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(SNIIngress)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNI.
func (in *SNI) DeepCopy() *SNI {
	if in == nil {
		return nil
	}
	out := new(SNI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNIIngress) DeepCopyInto(out *SNIIngress) {
	*out = *in
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNIIngress.
func (in *SNIIngress) DeepCopy() *SNIIngress {
	if in == nil {
		return nil
	}
	out := new(SNIIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedAPIServerNetworkPolicyControllerConfiguration) DeepCopyInto(out *SeedAPIServerNetworkPolicyControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedAPIServerNetworkPolicyControllerConfiguration.
func (in *SeedAPIServerNetworkPolicyControllerConfiguration) DeepCopy() *SeedAPIServerNetworkPolicyControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SeedAPIServerNetworkPolicyControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedClientConnection) DeepCopyInto(out *SeedClientConnection) {
	*out = *in
	out.ClientConnectionConfiguration = in.ClientConnectionConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedClientConnection.
func (in *SeedClientConnection) DeepCopy() *SeedClientConnection {
	if in == nil {
		return nil
	}
	out := new(SeedClientConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedConfig) DeepCopyInto(out *SeedConfig) {
	*out = *in
	in.SeedTemplate.DeepCopyInto(&out.SeedTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedConfig.
func (in *SeedConfig) DeepCopy() *SeedConfig {
	if in == nil {
		return nil
	}
	out := new(SeedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedControllerConfiguration) DeepCopyInto(out *SeedControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedControllerConfiguration.
func (in *SeedControllerConfiguration) DeepCopy() *SeedControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SeedControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	in.HTTPS.DeepCopyInto(&out.HTTPS)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootCareControllerConfiguration) DeepCopyInto(out *ShootCareControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.StaleExtensionHealthChecks != nil {
		in, out := &in.StaleExtensionHealthChecks, &out.StaleExtensionHealthChecks
		*out = new(StaleExtensionHealthChecks)
		(*in).DeepCopyInto(*out)
	}
	if in.ConditionThresholds != nil {
		in, out := &in.ConditionThresholds, &out.ConditionThresholds
		*out = make([]ConditionThreshold, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootCareControllerConfiguration.
func (in *ShootCareControllerConfiguration) DeepCopy() *ShootCareControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootCareControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootClientConnection) DeepCopyInto(out *ShootClientConnection) {
	*out = *in
	out.ClientConnectionConfiguration = in.ClientConnectionConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootClientConnection.
func (in *ShootClientConnection) DeepCopy() *ShootClientConnection {
	if in == nil {
		return nil
	}
	out := new(ShootClientConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootControllerConfiguration) DeepCopyInto(out *ShootControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.ProgressReportPeriod != nil {
		in, out := &in.ProgressReportPeriod, &out.ProgressReportPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ReconcileInMaintenanceOnly != nil {
		in, out := &in.ReconcileInMaintenanceOnly, &out.ReconcileInMaintenanceOnly
		*out = new(bool)
		**out = **in
	}
	if in.RespectSyncPeriodOverwrite != nil {
		in, out := &in.RespectSyncPeriodOverwrite, &out.RespectSyncPeriodOverwrite
		*out = new(bool)
		**out = **in
	}
	if in.RetryDuration != nil {
		in, out := &in.RetryDuration, &out.RetryDuration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.DNSEntryTTLSeconds != nil {
		in, out := &in.DNSEntryTTLSeconds, &out.DNSEntryTTLSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootControllerConfiguration.
func (in *ShootControllerConfiguration) DeepCopy() *ShootControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootStateSyncControllerConfiguration) DeepCopyInto(out *ShootStateSyncControllerConfiguration) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootStateSyncControllerConfiguration.
func (in *ShootStateSyncControllerConfiguration) DeepCopy() *ShootStateSyncControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootStateSyncControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootedSeedRegistrationControllerConfiguration) DeepCopyInto(out *ShootedSeedRegistrationControllerConfiguration) {
	*out = *in
	if in.SyncJitterPeriod != nil {
		in, out := &in.SyncJitterPeriod, &out.SyncJitterPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootedSeedRegistrationControllerConfiguration.
func (in *ShootedSeedRegistrationControllerConfiguration) DeepCopy() *ShootedSeedRegistrationControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootedSeedRegistrationControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaleExtensionHealthChecks) DeepCopyInto(out *StaleExtensionHealthChecks) {
	*out = *in
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaleExtensionHealthChecks.
func (in *StaleExtensionHealthChecks) DeepCopy() *StaleExtensionHealthChecks {
	if in == nil {
		return nil
	}
	out := new(StaleExtensionHealthChecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSServer) DeepCopyInto(out *TLSServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSServer.
func (in *TLSServer) DeepCopy() *TLSServer {
	if in == nil {
		return nil
	}
	out := new(TLSServer)
	in.DeepCopyInto(out)
	return out
}
