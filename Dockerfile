FROM busybox AS bin
COPY ./dist /binaries
RUN if [[ "$(arch)" == "x86_64" ]]; then \
        architecture="amd64"; \
    else \
        architecture="arm64"; \
    fi; \
    cp /binaries/imap-to-chat-bridge_linux-${architecture} /bin/imap-to-chat-bridge && \
    chmod +x /bin/imap-to-chat-bridge && \
    chown 65532:65532 /bin/imap-to-chat-bridge

FROM scratch
LABEL org.opencontainers.image.title="imap-to-chat-bridge"
LABEL org.opencontainers.image.description="Bridge to transfer mails from IMAP account to chat apps - timo-reymann/imap-to-chat-bridge"
LABEL org.opencontainers.image.ref.name="main"
LABEL org.opencontainers.image.licenses='GNU GPL v3'
LABEL org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.url="https://github.com/timo-reymann/imap-to-chat-bridge"
LABEL org.opencontainers.image.documentation="https://github.com/timo-reymann/imap-to-chat-bridge"
LABEL org.opencontainers.image.source="https://github.com/timo-reymann/imap-to-chat-bridge.git"
COPY --from=gcr.io/distroless/static-debian12:nonroot / /
USER nonroot
COPY --from=bin /bin/imap-to-chat-bridge /bin/imap-to-chat-bridge
ENTRYPOINT ["/bin/imap-to-chat-bridge"]
