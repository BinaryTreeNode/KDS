package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["proxy_test.go"],
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/proxy",
    library = ":go_default_library",
    deps = [
        "//cmd/kubeadm/app/apis/kubeadm/v1alpha1:go_default_library",
        "//cmd/kubeadm/app/util:go_default_library",
        "//cmd/kubeadm/app/util/config:go_default_library",
        "//pkg/apis/core:go_default_library",
        "//pkg/proxy/apis/kubeproxyconfig/v1alpha1:go_default_library",
        "//pkg/util/pointer:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/fake:go_default_library",
        "//vendor/k8s.io/client-go/testing:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "manifests.go",
        "proxy.go",
    ],
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/proxy",
    deps = [
        "//cmd/kubeadm/app/apis/kubeadm:go_default_library",
        "//cmd/kubeadm/app/constants:go_default_library",
        "//cmd/kubeadm/app/util:go_default_library",
        "//cmd/kubeadm/app/util/apiclient:go_default_library",
        "//pkg/api/legacyscheme:go_default_library",
        "//pkg/proxy/apis/kubeproxyconfig/scheme:go_default_library",
        "//pkg/proxy/apis/kubeproxyconfig/v1alpha1:go_default_library",
        "//pkg/util/version:go_default_library",
        "//plugin/pkg/scheduler/algorithm:go_default_library",
        "//vendor/k8s.io/api/apps/v1beta2:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/api/rbac/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
