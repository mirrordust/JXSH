FROM node:14.18.1-alpine
LABEL maintainer="shhliao754@outlook.com"

# used in building like --build-arg NPM_REG="some value"
ARG NPM_REG=https://registry.npm.taobao.org
ARG API_PREFIX=http://localhost:4000/api

ENV NODE_ENV=production \
    APP_PORT=3000 \
    REACT_APP_PREFIX=${API_PREFIX}

WORKDIR /app

COPY package*.json /app/
RUN npm install -g serve --registry=${NPM_REG} \
    && npm install --registry=${NPM_REG}

# The `COPY` instruction copies *new* files or directories from <src>
# and adds them to the filesystem of the container at the path <dest>.
COPY . /app/
# envs used in JavaScript code must be set before building
RUN npm run-script build

EXPOSE ${APP_PORT}

CMD ["sh", "-c", "serve -s build"]
