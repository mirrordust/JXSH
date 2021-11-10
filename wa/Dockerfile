FROM node:14.18.1-alpine
LABEL maintainer="shhliao754@outlook.com"

WORKDIR /app

COPY package*.json /app/
RUN npm install -g server --registry=https://registry.npm.taobao.org \
    && npm install --registry=https://registry.npm.taobao.org

# The `COPY` instruction copies *new* files or directories from <src>
# and adds them to the filesystem of the container at the path <dest>.
COPY . /app/
RUN npm run-script build

EXPOSE 3000

CMD ["serve", "-s", "build"]
