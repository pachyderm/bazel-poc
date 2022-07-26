load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

BAZEL_TOOLCHAIN_TAG = "0.7.2"

BAZEL_TOOLCHAIN_SHA = "f7aa8e59c9d3cafde6edb372d9bd25fb4ee7293ab20b916d867cd0baaa642529"

http_archive(
    name = "com_grail_bazel_toolchain",
    canonical_id = BAZEL_TOOLCHAIN_TAG,
    sha256 = BAZEL_TOOLCHAIN_SHA,
    strip_prefix = "bazel-toolchain-{tag}".format(tag = BAZEL_TOOLCHAIN_TAG),
    url = "https://github.com/grailbio/bazel-toolchain/archive/{tag}.tar.gz".format(tag = BAZEL_TOOLCHAIN_TAG),
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "685052b498b6ddfe562ca7a97736741d87916fe536623afb7da2824c0211c369",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.33.0/rules_go-v0.33.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.33.0/rules_go-v0.33.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "501deb3d5695ab658e82f6f6f549ba681ea3ca2a5fb7911154b5aa45596183fa",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.26.0/bazel-gazelle-v0.26.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.26.0/bazel-gazelle-v0.26.0.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "c29d8b4b79389463c546f98b15aa4391d4ed7ec459340c47bffe15db63eb9126",
    strip_prefix = "protobuf-3.21.3",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.21.3.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.21.3.tar.gz",
    ],
)

load("@com_grail_bazel_toolchain//toolchain:deps.bzl", "bazel_toolchain_dependencies")

bazel_toolchain_dependencies()

load("@com_grail_bazel_toolchain//toolchain:rules.bzl", "llvm_toolchain")

LLVM_VERSION = "14.0.0"

llvm_toolchain(
    name = "llvm_toolchain",
    llvm_version = LLVM_VERSION,
)

# For new sysroots, read this Python script: https://chromium.googlesource.com/chromium/src/build/+/refs/heads/main/linux/sysroot_scripts/install-sysroot.py
# which reads this data file: https://chromium.googlesource.com/chromium/src/build/+/refs/heads/main/linux/sysroot_scripts/sysroots.json
http_archive(
    name = "org_chromium_sysroot_linux_x64",
    build_file_content = """
filegroup(
  name = "sysroot",
  srcs = glob(["*/**"]),
  visibility = ["//visibility:public"],
)
""",
    urls = ["https://commondatastorage.googleapis.com/chrome-linux-sysroot/toolchain/cb4fa34f1faddafb72cace35faf62a611f2ca7c9/debian_bullseye_amd64_sysroot.tar.xz"],
)

http_archive(
    name = "org_chromium_sysroot_linux_arm64",
    build_file_content = """
filegroup(
  name = "sysroot",
  srcs = glob(["*/**"]),
  visibility = ["//visibility:public"],
)
""",
    urls = ["https://commondatastorage.googleapis.com/chrome-linux-sysroot/toolchain/f00ece500aef0ff2a431a07de48bd3e1aa6d1caf/debian_bullseye_arm64_sysroot.tar.xz"],
)

llvm_toolchain(
    name = "llvm_toolchain_with_sysroot",
    llvm_version = LLVM_VERSION,
    sysroot = {
        "linux-x86_64": "@org_chromium_sysroot_linux_x64//:sysroot",
        "linux-aarch64": "@org_chromium_sysroot_linux_arm64//:sysroot",
    },
    # We can share the downloaded LLVM distribution with the first configuration.
    toolchain_roots = {
        "": "@llvm_toolchain_llvm//",
    },
)

load("@llvm_toolchain//:toolchains.bzl", "llvm_register_toolchains")

llvm_register_toolchains()

load("@llvm_toolchain_with_sysroot//:toolchains.bzl", "llvm_register_toolchains")

llvm_register_toolchains()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("//:deps.bzl", "go_dependencies")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:go_lint",
    version = "1.18.4",
)

gazelle_dependencies()

protobuf_deps()

git_repository(
    name = "com_github_sluongng_nogo_analyzer",
    commit = "d6a3958307f2b485cda307f453e06db025415177",
    remote = "https://github.com/sluongng/nogo-analyzer",
    shallow_since = "1656919128 +0200",
)
