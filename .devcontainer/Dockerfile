FROM mcr.microsoft.com/devcontainers/go:0-1.20

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends postgresql-client httpie

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install github.com/githubnemo/CompileDaemon@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest
USER root 

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
