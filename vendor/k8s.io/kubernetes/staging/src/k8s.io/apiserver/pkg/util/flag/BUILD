package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = [
        "colon_separated_multimap_string_string_test.go",
        "langle_separated_map_string_string_test.go",
        "map_string_bool_test.go",
        "map_string_string_test.go",
        "namedcertkey_flag_test.go",
    ],
    importpath = "k8s.io/apiserver/pkg/util/flag",
    library = ":go_default_library",
    deps = ["//vendor/github.com/spf13/pflag:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "colon_separated_multimap_string_string.go",
        "configuration_map.go",
        "flags.go",
        "langle_separated_map_string_string.go",
        "map_string_bool.go",
        "map_string_string.go",
        "namedcertkey_flag.go",
        "omitempty.go",
        "string_flag.go",
        "tristate.go",
    ],
    importpath = "k8s.io/apiserver/pkg/util/flag",
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
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
