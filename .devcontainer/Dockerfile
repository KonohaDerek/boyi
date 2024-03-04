# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.191.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.16, 1.15
FROM mcr.microsoft.com/devcontainers/go:1-1.21

RUN apt-get update
RUN go env GOENV
# RUN su vscode -c "go env -w GOOS=darwin"
# RUN su vscode -c "go env -w GOARCH=arm64"
RUN su vscode -c "go install github.com/swaggo/swag/cmd/swag@latest"
RUN su vscode -c "go install github.com/99designs/gqlgen@latest"
RUN su vscode -c "go install github.com/golang/mock/mockgen@latest" 
RUN su vscode -c "go install github.com/pressly/goose/v3/cmd/goose@latest"
RUN su vscode -c "go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest"