### BASE
################################################
FROM node:19-alpine AS base

WORKDIR /usr/src/app

RUN npm install -g npm

COPY ./package.json ./
COPY ./package-lock.json ./

RUN npm install

COPY ./ ./

### DEVELOPMENT
################################################
FROM base AS development

ENTRYPOINT ["npm", "run", "dev"]

### PRODUCTION
################################################
FROM base AS production

RUN npm run build

# ENTRYPOINT [ "echo", "Building frontend done" ]

# RUN rm -rf ./dist
# RUN mkdir ./dist
# RUN chown -R node:node ./dist

# USER node

# ENTRYPOINT ["npm", "run", "build"]