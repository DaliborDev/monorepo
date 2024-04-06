load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "lib-daemonize",
    srcs = glob(["*.go"]),
    visibility = ["//visibility:private"],
)

go_binary(
    name = "daemonize",
    embed = [
        ":lib-daemonize"
    ],
    data = glob([
        "files/*"
    ])
)