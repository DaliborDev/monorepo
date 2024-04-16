load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "lib-main",
    srcs = glob(["**/*.go"]),
    visibility = ["//visibility:private"],
    importpath = "github.com/DaliborDev/monorepo/pulse"
)

go_binary(
    name = "run",
    embed = [
        ":lib-main"
    ],
    data = glob([
        "files/*"
    ])
)