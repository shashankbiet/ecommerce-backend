FROM node:alpine

WORKDIR /usr/src/app

COPY api-gateway/package.json api-gateway/package-lock.json ./
RUN --mount=type=cache,target=/root/.npm \
    npm ci --omit=dev \
    && rm -rf "$(npm root -g)/npm" \
    && rm -f "$(command -v npm)" "$(command -v npx)"

COPY api-gateway/. .

USER node

ENV NODE_ENV=prod

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 CMD ["node", "-e", "fetch('http://127.0.0.1:3001/health').then((response) => process.exit(response.ok ? 0 : 1)).catch(() => process.exit(1))"]

CMD ["node", "server.js"]