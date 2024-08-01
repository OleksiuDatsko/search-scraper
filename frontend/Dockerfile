FROM node:20 AS build

WORKDIR /app

COPY package*.json ./

# Install dependencies
RUN --mount=type=cache,target=/root/.yarn \
    --mount=type=cache,target=/app/node_modules \
    yarn

COPY . .
# Build
RUN --mount=type=cache,target=/root/.yarn \
    --mount=type=cache,target=/app/.svelte-kit \
    yarn run build

FROM node:20-alpine

RUN addgroup -S server_group && adduser -S server -G server_group

COPY --from=build /app/build /app/build
COPY --from=build /app/package*.json /app/

USER server

EXPOSE 3000

CMD [ "node", "/app/build/index.js" ]
