# The buildstage only compiles the node modules and copies the fnal artifacts into the deployment stage.
FROM node:20.11-alpine as buildstage

COPY ./workspace /build
WORKDIR /build

RUN npm install && \
    npm run build

# The final deployment stage which is used for hosting the compiled build.
FROM nginx
COPY --from=buildstage /build/dist /usr/share/nginx/html
