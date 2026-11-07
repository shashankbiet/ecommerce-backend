FROM node:alpine

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies from the lockfile.
COPY user-service/package.json user-service/package-lock.json ./
RUN --mount=type=cache,target=/root/.npm \
    npm ci --omit=dev \
    && rm -rf "$(npm root -g)/npm" \
    && rm -f "$(command -v npm)" "$(command -v npx)"

# Bundle app source
COPY user-service/. .

USER node

ENV NODE_ENV=prod

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 CMD ["node", "-e", "fetch('http://127.0.0.1:3002/api/user/health').then((response) => process.exit(response.ok ? 0 : 1)).catch(() => process.exit(1))"]

CMD ["node", "server.js"]