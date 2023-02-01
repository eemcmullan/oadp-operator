//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationConfig) DeepCopyInto(out *ApplicationConfig) {
	*out = *in
	if in.Velero != nil {
		in, out := &in.Velero, &out.Velero
		*out = new(VeleroConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Restic != nil {
		in, out := &in.Restic, &out.Restic
		*out = new(ResticConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationConfig.
func (in *ApplicationConfig) DeepCopy() *ApplicationConfig {
	if in == nil {
		return nil
	}
	out := new(ApplicationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocation) DeepCopyInto(out *BackupLocation) {
	*out = *in
	if in.Velero != nil {
		in, out := &in.Velero, &out.Velero
		*out = new(velerov1.BackupStorageLocationSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudStorage != nil {
		in, out := &in.CloudStorage, &out.CloudStorage
		*out = new(CloudStorageLocation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocation.
func (in *BackupLocation) DeepCopy() *BackupLocation {
	if in == nil {
		return nil
	}
	out := new(BackupLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorage) DeepCopyInto(out *CloudStorage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorage.
func (in *CloudStorage) DeepCopy() *CloudStorage {
	if in == nil {
		return nil
	}
	out := new(CloudStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStorage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageList) DeepCopyInto(out *CloudStorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageList.
func (in *CloudStorageList) DeepCopy() *CloudStorageList {
	if in == nil {
		return nil
	}
	out := new(CloudStorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageLocation) DeepCopyInto(out *CloudStorageLocation) {
	*out = *in
	out.CloudStorageRef = in.CloudStorageRef
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupSyncPeriod != nil {
		in, out := &in.BackupSyncPeriod, &out.BackupSyncPeriod
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageLocation.
func (in *CloudStorageLocation) DeepCopy() *CloudStorageLocation {
	if in == nil {
		return nil
	}
	out := new(CloudStorageLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageSpec) DeepCopyInto(out *CloudStorageSpec) {
	*out = *in
	in.CreationSecret.DeepCopyInto(&out.CreationSecret)
	if in.EnableSharedConfig != nil {
		in, out := &in.EnableSharedConfig, &out.EnableSharedConfig
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageSpec.
func (in *CloudStorageSpec) DeepCopy() *CloudStorageSpec {
	if in == nil {
		return nil
	}
	out := new(CloudStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageStatus) DeepCopyInto(out *CloudStorageStatus) {
	*out = *in
	if in.LastSynced != nil {
		in, out := &in.LastSynced, &out.LastSynced
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageStatus.
func (in *CloudStorageStatus) DeepCopy() *CloudStorageStatus {
	if in == nil {
		return nil
	}
	out := new(CloudStorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPlugin) DeepCopyInto(out *CustomPlugin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPlugin.
func (in *CustomPlugin) DeepCopy() *CustomPlugin {
	if in == nil {
		return nil
	}
	out := new(CustomPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataMover) DeepCopyInto(out *DataMover) {
	*out = *in
	if in.MaxConcurrentBackupVolumes != nil {
		in, out := &in.MaxConcurrentBackupVolumes, &out.MaxConcurrentBackupVolumes
		*out = new(int64)
		**out = **in
	}
	if in.MaxConcurrentRestoreVolumes != nil {
		in, out := &in.MaxConcurrentRestoreVolumes, &out.MaxConcurrentRestoreVolumes
		*out = new(int64)
		**out = **in
	}
	if in.MoverConfig != nil {
		in, out := &in.MoverConfig, &out.MoverConfig
		*out = new(MoverConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataMover.
func (in *DataMover) DeepCopy() *DataMover {
	if in == nil {
		return nil
	}
	out := new(DataMover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProtectionApplication) DeepCopyInto(out *DataProtectionApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProtectionApplication.
func (in *DataProtectionApplication) DeepCopy() *DataProtectionApplication {
	if in == nil {
		return nil
	}
	out := new(DataProtectionApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataProtectionApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProtectionApplicationList) DeepCopyInto(out *DataProtectionApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataProtectionApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProtectionApplicationList.
func (in *DataProtectionApplicationList) DeepCopy() *DataProtectionApplicationList {
	if in == nil {
		return nil
	}
	out := new(DataProtectionApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataProtectionApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProtectionApplicationSpec) DeepCopyInto(out *DataProtectionApplicationSpec) {
	*out = *in
	if in.BackupLocations != nil {
		in, out := &in.BackupLocations, &out.BackupLocations
		*out = make([]BackupLocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SnapshotLocations != nil {
		in, out := &in.SnapshotLocations, &out.SnapshotLocations
		*out = make([]SnapshotLocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UnsupportedOverrides != nil {
		in, out := &in.UnsupportedOverrides, &out.UnsupportedOverrides
		*out = make(map[UnsupportedImageKey]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PodDnsConfig.DeepCopyInto(&out.PodDnsConfig)
	if in.BackupImages != nil {
		in, out := &in.BackupImages, &out.BackupImages
		*out = new(bool)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(ApplicationConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(Features)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProtectionApplicationSpec.
func (in *DataProtectionApplicationSpec) DeepCopy() *DataProtectionApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(DataProtectionApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProtectionApplicationStatus) DeepCopyInto(out *DataProtectionApplicationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProtectionApplicationStatus.
func (in *DataProtectionApplicationStatus) DeepCopy() *DataProtectionApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(DataProtectionApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	if in.DataMover != nil {
		in, out := &in.DataMover, &out.DataMover
		*out = new(DataMover)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MoverConfig) DeepCopyInto(out *MoverConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MoverConfig.
func (in *MoverConfig) DeepCopy() *MoverConfig {
	if in == nil {
		return nil
	}
	out := new(MoverConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodConfig) DeepCopyInto(out *PodConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ResourceAllocations.DeepCopyInto(&out.ResourceAllocations)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodConfig.
func (in *PodConfig) DeepCopy() *PodConfig {
	if in == nil {
		return nil
	}
	out := new(PodConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResticConfig) DeepCopyInto(out *ResticConfig) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.SupplementalGroups != nil {
		in, out := &in.SupplementalGroups, &out.SupplementalGroups
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.PodConfig != nil {
		in, out := &in.PodConfig, &out.PodConfig
		*out = new(PodConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResticConfig.
func (in *ResticConfig) DeepCopy() *ResticConfig {
	if in == nil {
		return nil
	}
	out := new(ResticConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotLocation) DeepCopyInto(out *SnapshotLocation) {
	*out = *in
	if in.Velero != nil {
		in, out := &in.Velero, &out.Velero
		*out = new(velerov1.VolumeSnapshotLocationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotLocation.
func (in *SnapshotLocation) DeepCopy() *SnapshotLocation {
	if in == nil {
		return nil
	}
	out := new(SnapshotLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroConfig) DeepCopyInto(out *VeleroConfig) {
	*out = *in
	if in.FeatureFlags != nil {
		in, out := &in.FeatureFlags, &out.FeatureFlags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultPlugins != nil {
		in, out := &in.DefaultPlugins, &out.DefaultPlugins
		*out = make([]DefaultPlugin, len(*in))
		copy(*out, *in)
	}
	if in.CustomPlugins != nil {
		in, out := &in.CustomPlugins, &out.CustomPlugins
		*out = make([]CustomPlugin, len(*in))
		copy(*out, *in)
	}
	if in.PodConfig != nil {
		in, out := &in.PodConfig, &out.PodConfig
		*out = new(PodConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroConfig.
func (in *VeleroConfig) DeepCopy() *VeleroConfig {
	if in == nil {
		return nil
	}
	out := new(VeleroConfig)
	in.DeepCopyInto(out)
	return out
}
