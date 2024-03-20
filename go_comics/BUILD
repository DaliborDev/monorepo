load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "load_lib",
    srcs = glob(["cmd/**/*.go"]),
    visibility = ["//visibility:public"],
)

go_binary(
    name = "load",
    deps = [
        ":load_lib",
    ],
)
