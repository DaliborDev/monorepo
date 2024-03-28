load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "lib-roman2integer",
    srcs = glob(["roman2integer/*.go"]),
    visibility = ["//visibility:private"],
)

go_binary(
    name = "roman2integer",
    embed = [
        ":lib-roman2integer"
    ],
)