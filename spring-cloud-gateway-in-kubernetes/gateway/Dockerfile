FROM openjdk:8-jdk-alpine as build
COPY . /app/gateway
WORKDIR /app/gateway
RUN ./gradlew clean build


FROM openjdk:8-jdk-alpine
COPY --from=build /app/gateway/build/libs/gateway-0.0.1-SNAPSHOT.jar /app/gateway.jar
EXPOSE 8080
ENTRYPOINT [ "java", "-jar", "/app/gateway.jar" ]