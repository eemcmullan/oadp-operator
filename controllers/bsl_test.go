package controllers

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-logr/logr"
	oadpv1alpha1 "github.com/openshift/oadp-operator/api/v1alpha1"
	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func getSchemeForFakeClient() (*runtime.Scheme, error) {
	err := oadpv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	err = velerov1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	return scheme.Scheme, nil
}

func getFakeClientFromObjects(objs ...client.Object) (client.WithWatch, error) {
	schemeForFakeClient, err := getSchemeForFakeClient()
	if err != nil {
		return nil, err
	}

	return fake.NewClientBuilder().WithScheme(schemeForFakeClient).WithObjects(objs...).Build(), nil
}

func TestVeleroReconciler_ValidateBackupStorageLocations(t *testing.T) {
	tests := []struct {
		name     string
		VeleroCR *oadpv1alpha1.Velero
		want     bool
		wantErr  bool
	}{
		{
			name: "test no BSLs, no noobaa",
			VeleroCR: &oadpv1alpha1.Velero{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "bar",
				},
				Spec: oadpv1alpha1.VeleroSpec{},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "test BSLs specified, no noobaa",
			VeleroCR: &oadpv1alpha1.Velero{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "bar",
				},
				Spec: oadpv1alpha1.VeleroSpec{
					BackupStorageLocations: []velerov1.BackupStorageLocationSpec{
						{
							// TODO: foo is invalid provider, add test cases for it
							Provider: "foo",
						},
					},
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test no BSL, noobaa configured",
			VeleroCR: &oadpv1alpha1.Velero{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "bar",
				},
				Spec: oadpv1alpha1.VeleroSpec{
					Noobaa: true,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name:     "test get error",
			VeleroCR: &oadpv1alpha1.Velero{},
			want:     false,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient, err := getFakeClientFromObjects(tt.VeleroCR)
			if err != nil {
				t.Errorf("error in creating fake client, likely programmer error")
			}
			r := &VeleroReconciler{
				Client:  fakeClient,
				Scheme:  fakeClient.Scheme(),
				Log:     logr.Discard(),
				Context: newContextForTest(tt.name),
				NamespacedName: types.NamespacedName{
					Namespace: tt.VeleroCR.Namespace,
					Name:      tt.VeleroCR.Name,
				},
				EventRecorder: record.NewFakeRecorder(10),
			}
			got, err := r.ValidateBackupStorageLocations(r.Log)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBackupStorageLocations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateBackupStorageLocations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func newContextForTest(name string) context.Context {
	return context.TODO()
}

func TestVeleroReconciler_updateBSLFromSpec(t *testing.T) {
	tests := []struct {
		name    string
		bsl     *velerov1.BackupStorageLocation
		velero  *oadpv1alpha1.Velero
		wantErr bool
	}{
		{
			name: "BSL without owner reference and labels",
			bsl: &velerov1.BackupStorageLocation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo-1",
					Namespace: "bar",
				},
			},
			velero: &oadpv1alpha1.Velero{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheme, err := getSchemeForFakeClient()
			if err != nil {
				t.Errorf("error getting scheme for the test: %#v", err)
			}
			r := &VeleroReconciler{
				Scheme: scheme,
			}

			wantBSl := &velerov1.BackupStorageLocation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo-1",
					Namespace: "bar",
					Labels: map[string]string{
						"app.kubernetes.io/name":     "oadp-operator-velero",
						"app.kubernetes.io/instance": tt.velero.Name + "-1",
						//"app.kubernetes.io/version":    "x.y.z",
						"app.kubernetes.io/managed-by": "oadp-operator",
						"app.kubernetes.io/component":  "bsl",
					},
					OwnerReferences: []metav1.OwnerReference{{
						APIVersion:         oadpv1alpha1.SchemeBuilder.GroupVersion.String(),
						Kind:               "Velero",
						Name:               tt.velero.Name,
						UID:                tt.velero.UID,
						Controller:         pointer.BoolPtr(true),
						BlockOwnerDeletion: pointer.BoolPtr(true),
					}},
				},
			}

			err = r.updateBSLFromSpec(tt.bsl, tt.velero)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateBSLFromSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.bsl.Labels, wantBSl.Labels) {
				t.Errorf("expected bsl labels to be %#v, got %#v", wantBSl.Labels, tt.bsl.Labels)
			}
			if !reflect.DeepEqual(tt.bsl.OwnerReferences, wantBSl.OwnerReferences) {
				t.Errorf("expected bsl owner references to be %#v, got %#v", wantBSl.OwnerReferences, tt.bsl.OwnerReferences)
			}
		})
	}
}
