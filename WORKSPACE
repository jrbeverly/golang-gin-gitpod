workspace(name = "golang-gin-consignment")

load("//:bazel/rules/deps.bzl", "bazel_dependencies")

bazel_dependencies()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:bazel/go/deps.bzl", "BAZEL_GOLANG_VERSION", "go_dependencies")

# gazelle:repository_macro bazel/go/deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = BAZEL_GOLANG_VERSION)

gazelle_dependencies()

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")

_go_image_repos()
