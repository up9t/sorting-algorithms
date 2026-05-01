FROM golang:1-trixie

ARG USERNAME=dev
ARG UID=1000
ARG GID=1000

RUN groupadd --gid $GID $USERNAME \
  && useradd --gid $GID --uid $UID --create-home --shell /usr/bin/bash $USERNAME

RUN apt-get update \
  && apt-get install -y sudo \
  && echo "$USERNAME ALL=(root) NOPASSWD: ALL" > /etc/sudoers.d/$USERNAME

USER dev