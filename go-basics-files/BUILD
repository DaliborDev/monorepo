load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# cmd -> generate
go_library(
    name = "generate-file-library",
    srcs = glob(["cmd/generate/*.go"]),
    visibility = ["//visibility:private"],
)

go_binary(
    name = "generate-file",
    embed = [
        ":generate-file-library"
    ],
    data = glob([
        "files/*"
    ])
)

# cmd -> read
go_library(
    name = "read-file-library",
    srcs = glob(["cmd/read/*.go"]),
    visibility = ["//visibility:private"],
)

go_binary(
    name = "read-file",
    embed = [
        ":read-file-library",
    ],
    data = glob([
        "files/*"
    ]),
)