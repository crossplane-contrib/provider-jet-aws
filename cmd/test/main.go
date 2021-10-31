package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	CRDPath = "/Users/monus/go/src/github.com/crossplane/provider-tf-aws/package/crds"

	// OpenAPIEndpoint         = "http://127.0.0.1:8001/openapi/v2"
	OpenAPIEndpoint         = "https://kubernetes.default/openapi/v2"
	SchemaCheckInterval     = time.Second * 3
	SchemaCheckBatchTimeout = time.Minute * 10
	CooldownPeriodPerBatch  = time.Second * 0
	BatchQuantity           = 1
)

func main() {
	crds, err := ReadCRDs()
	kingpin.FatalIfError(err, "cannot read crds")
	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "cannot get config")
	mgr, err := ctrl.NewManager(cfg, ctrl.Options{})
	kingpin.FatalIfError(err, "cannot create new manager")
	kingpin.FatalIfError(v1.AddToScheme(mgr.GetScheme()), "cannot add crd package to scheme")
	kube := mgr.GetClient()

	ctx := context.TODO()
	start := time.Now()
	fmt.Printf("started at %s\n", start.String())
	var created []*v1.CustomResourceDefinition
	batchNo := 1
	for i, crd := range crds {
		if i != 0 && i%BatchQuantity == 0 {
			fmt.Printf("let me check if %d CRDs made it to OpenAPI schema\n", len(created))
			if err := WaitSchemaAggregation(created); err != nil {
				panic(errors.Wrap(err, "cannot wait until crd schemas are aggregated"))
			}
			wait := CooldownPeriodPerBatch * time.Duration(batchNo)
			fmt.Printf("\n\n %d out of %d CRDs have been installed. Waiting for %s to start the next batch...\n\n", i+1, len(crds), wait.String())
			time.Sleep(wait)
			created = []*v1.CustomResourceDefinition{}
			batchNo++
		}
		if err := kube.Create(ctx, crd); resource.Ignore(kerrors.IsAlreadyExists, err) != nil {
			panic(errors.Wrapf(err, "cannot create crd: %s", crd.Name))
		}
		created = append(created, crd)
	}
	fmt.Println("Checking schema for all CRDs...")
	if err := WaitSchemaAggregation(crds); err != nil {
		panic(errors.Wrap(err, "cannot wait until crd schemas are aggregated"))
	}
	end := time.Now()
	fmt.Printf("finished at %s\n", end.String())
	fmt.Printf("the whole installation took %s\n", end.Sub(start).String())
}

func WaitSchemaAggregation(crds []*v1.CustomResourceDefinition) error {
	start := time.Now()
	for {
		// curl -s https://kubernetes.default/openapi/v2  --header "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
		client := http.DefaultClient
		req, err := http.NewRequest(http.MethodGet, OpenAPIEndpoint, nil)
		if err != nil {
			return errors.Wrap(err, "cannot create a new request")
		}
		if os.Getenv("SERVICE_ACCOUNT_TOKEN_PATH") != "" {
			token, err := os.ReadFile(os.Getenv("SERVICE_ACCOUNT_TOKEN_PATH"))
			if err != nil {
				return errors.Wrap(err, "cannot read service account token")
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", string(token)))
		}
		if os.Getenv("KUBERNETES_CA_CERT_PATH") != "" {
			caCert, err := os.ReadFile(os.Getenv("KUBERNETES_CA_CERT_PATH"))
			if err != nil {
				return errors.Wrap(err, "cannot read cert")
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)

			client = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						RootCAs: caCertPool,
					},
				},
			}
		}
		resp, err := client.Do(req)
		if err != nil {
			return errors.Wrap(err, "request failed")
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "cannot read response data")
		}
		schemaJSON := string(data)
		notFound := -1
		for i, crd := range crds {
			// globalaccelerator.aws.tf.crossplane.io/v1alpha1/accelerators
			str := fmt.Sprintf("%s/%s/%s", crd.Spec.Group, crd.Spec.Versions[0].Name, crd.Spec.Names.Plural)
			if !strings.Contains(schemaJSON, str) {
				notFound = i
				break
			}
		}
		if notFound == -1 {
			break
		}
		if time.Now().Sub(start) > SchemaCheckBatchTimeout {
			return errors.Errorf("could not go further than %s whose index is %d in %s aborting...\n", crds[notFound].Name, notFound, SchemaCheckBatchTimeout.String())
		}
		fmt.Printf("could not find %d, waiting for %s before retrying...\n", notFound, SchemaCheckInterval.String())
		time.Sleep(SchemaCheckInterval)
	}
	fmt.Printf("the batch with %d CRDs took %s to install\n", len(crds), time.Now().Sub(start).String())
	return nil
}

func ReadCRDs() ([]*v1.CustomResourceDefinition, error) {
	files, err := os.ReadDir(CRDPath)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read crds directory")
	}
	crds := make([]*v1.CustomResourceDefinition, len(files))
	for i, f := range files {
		if f.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(CRDPath, f.Name()))
		if err != nil {
			return nil, errors.Wrapf(err, "cannot read file: %s", f.Name())
		}
		crd := &v1.CustomResourceDefinition{}
		if err := yaml.Unmarshal(data, crd); err != nil {
			return nil, errors.Wrap(err, "cannot unmarshal file into crd")
		}
		crds[i] = crd
	}
	sort.SliceStable(crds, func(i, j int) bool {
		return crds[i].Name < crds[j].Name
	})
	return crds, nil
}
