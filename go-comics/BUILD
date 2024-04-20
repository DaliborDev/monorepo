load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# cmd -> load
go_library(
    name = "load_lib",
    srcs = glob(["cmd/load/*.go"]),
    visibility = ["//visibility:private"],
)

go_binary(
    name = "load",
    embed = [
        ":load_lib",
    ],
)

# cmd -> search
go_library(
    name = "search_lib",
    srcs = glob(["cmd/search/*.go"]),
    visibility = ["//visibility:private"],
)

go_binary(
    name = "search",
    data = [
        "files/xkcd.json",
    ],
    embed = [
        ":search_lib",
    ],
)
