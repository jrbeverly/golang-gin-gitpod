FROM gitpod/workspace-full
USER root

RUN mkdir $HOME/bin
RUN npm install -g @bazel/bazelisk
RUN npm install -g @bazel/ibazel

ARG BAZEL_BUILDTOOLS_URL="https://github.com/bazelbuild/buildtools/releases/download"
RUN curl -sSL "${BAZEL_BUILDTOOLS_URL}/5.1.0/buildozer-linux-amd64" > /usr/local/bin/buildozer && \
    curl -sSL "${BAZEL_BUILDTOOLS_URL}/5.1.0/buildifier-linux-amd64" >  /usr/local/bin/buildifier && \
    chmod +x /usr/local/bin/buildozer && \
    chmod +x /usr/local/bin/buildifier 

USER gitpod
