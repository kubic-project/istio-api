// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rbac/v1alpha1/rbac.proto

// Istio RBAC (Role Based Access Control) defines ServiceRole and ServiceRoleBinding
// objects.
//
// A ServiceRole specification includes a list of rules (permissions). Each rule has
// the following standard fields:
//
//   * services: a list of services.
//   * methods: A list of HTTP methods. You can set the value to `\*` to include all HTTP methods.
//              This field should not be set for TCP services. The policy will be ignored.
//              For gRPC services, only `POST` is allowed; other methods will result in denying services.
//   * paths: HTTP paths or gRPC methods. Note that gRPC methods should be
//     presented in the form of "/packageName.serviceName/methodName" and are case sensitive.
//
// In addition to the standard fields, operators can also use custom keys in the `constraints` field,
// the supported keys are listed in the "constraints and properties" page.
//
// Below is an example of ServiceRole object "product-viewer", which has "read" ("GET" and "HEAD")
// access to "products.svc.cluster.local" service at versions "v1" and "v2". "path" is not specified,
// so it applies to any path in the service.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: ServiceRole
// metadata:
//   name: products-viewer
//   namespace: default
// spec:
//   rules:
//   - services: ["products.svc.cluster.local"]
//     methods: ["GET", "HEAD"]
//     constraints:
//     - key: "destination.labels[version]"
//       values: ["v1", "v2"]
// ```
//
// A ServiceRoleBinding specification includes two parts:
//
//  * The `roleRef` field that refers to a ServiceRole object in the same namespace.
//  * A list of `subjects` that are assigned the roles.
//
// In addition to a simple `user` field, operators can also use custom keys in the `properties` field,
// the supported keys are listed in the "constraints and properties" page.
//
// Below is an example of ServiceRoleBinding object "test-binding-products", which binds two subjects
// to ServiceRole "product-viewer":
//
//   * User "alice@yahoo.com"
//   * Services in "abc" namespace.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: ServiceRoleBinding
// metadata:
//   name: test-binding-products
//   namespace: default
// spec:
//   subjects:
//   - user: alice@yahoo.com
//   - properties:
//       source.namespace: "abc"
//   roleRef:
//     kind: ServiceRole
//     name: "products-viewer"
// ```

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using ServiceRole within kubernetes types, where deepcopy-gen is used.
func (in *ServiceRole) DeepCopyInto(out *ServiceRole) {
	proto.Merge(out, in)
}

// DeepCopyInto supports using ServiceRoleBinding within kubernetes types, where deepcopy-gen is used.
func (in *ServiceRoleBinding) DeepCopyInto(out *ServiceRoleBinding) {
	proto.Merge(out, in)
}

// DeepCopyInto supports using RbacConfig within kubernetes types, where deepcopy-gen is used.
func (in *RbacConfig) DeepCopyInto(out *RbacConfig) {
	proto.Merge(out, in)
}
