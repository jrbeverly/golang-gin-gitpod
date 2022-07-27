#!/usr/bin/env bash
#===============================================================================
# Emit variables external to the bazel build process that are used in the process
# of stamping compiled artifacts.
#
# These variables are grouped by their function
#====================================

# Version properties for tool
echo STABLE_GIT_COMMIT $(git rev-parse HEAD)
echo TOOL_VERSION $(git tag --points-at HEAD | head -n 1)