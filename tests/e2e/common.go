package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func setUpClient() (*kubernetes.Clientset, error) {
	kubeConf := getKubeConfig()
	// create client for pod
	clientset, err := kubernetes.NewForConfig(kubeConf)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func getJsonData(path string) ([]byte, error) {
	// Return buffer data for json
	jsonData, err := ioutil.ReadFile(path)
	return jsonData, err
}

func decodeJson(data []byte) (map[string]interface{}, error) {
	// Return JSON from buffer data
	var jsonData map[string]interface{}

	err := json.Unmarshal(data, &jsonData)
	return jsonData, err
}

// Keeping it for now.
func createOADPTestNamespace(namespace string) error {
	// default OADP Namespace
	kubeConf := getKubeConfig()
	clientset, err := kubernetes.NewForConfig(kubeConf)
	if err != nil {
		return err
	}
	ns := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &ns, metav1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return nil
	}

	return err
}

// Keeping it for now.
func deleteOADPTestNamespace(namespace string) error {
	// default OADP Namespace
	kubeConf := getKubeConfig()
	clientset, err := kubernetes.NewForConfig(kubeConf)

	if err != nil {
		return err
	}
	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	return err
}

func getKubeConfig() *rest.Config {
	return config.GetConfigOrDie()
}

// Keeping it for now
func doesNamespaceExists(namespace string) (bool, error) {
	clientset, err := setUpClient()
	if err != nil {
		return false, err
	}
	_, err = clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	fmt.Println("Test namespace already exists")
	return true, nil
}

// Keeping it for now.
func isNamespaceDeleted(namespace string) wait.ConditionFunc {
	fmt.Println("Checking test namespace has been deleted...")
	return func() (bool, error) {
		clientset, err := setUpClient()
		if err != nil {
			return false, err
		}
		_, err = clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
		if err != nil {
			fmt.Println("Test namespace has been deleted")
			return true, nil
		}
		return false, err
	}
}

func getCredsData(cloud string) ([]byte, error) {
	// pass in aws credentials by cli flag
	// from cli:  -cloud=<"filepath">
	// go run main.go -cloud="/Users/emilymcmullan/.aws/credentials"
	// cloud := flag.String("cloud", "", "file path for aws credentials")
	// flag.Parse()
	// save passed in cred file as []byte
	credsFile, err := ioutil.ReadFile(cloud)
	return credsFile, err
}

func createCredentialsSecret(data []byte, namespace string, credSecretRef string) error {
	clientset, err := setUpClient()
	if err != nil {
		return err
	}
	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      credSecretRef,
			Namespace: namespace,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: metav1.SchemeGroupVersion.String(),
		},
		Data: map[string][]byte{
			"cloud": data,
		},
		Type: corev1.SecretTypeOpaque,
	}
	_, err = clientset.CoreV1().Secrets(namespace).Create(context.TODO(), &secret, metav1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {

		return nil
	}
	return err
}

func deleteSecret(namespace string, credSecretRef string) error {
	clientset, err := setUpClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Secrets(namespace).Delete(context.Background(), credSecretRef, metav1.DeleteOptions{})
	return err
}

func isCredentialsSecretDeleted(namespace string, credSecretRef string) wait.ConditionFunc {
	fmt.Println("Checking secret has been deleted...")
	return func() (bool, error) {
		clientset, err := setUpClient()
		if err != nil {
			return false, err
		}
		_, err = clientset.CoreV1().Secrets(namespace).Get(context.Background(), credSecretRef, metav1.GetOptions{})
		if err != nil {
			fmt.Println("Secret in test namespace has been deleted")
			return true, nil
		}
		fmt.Println("Secret still exists in namespace")
		return false, err
	}
}
