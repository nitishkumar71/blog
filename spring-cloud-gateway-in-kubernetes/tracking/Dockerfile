FROM openjdk:8-jdk-alpine as build
COPY . /app/tracking
WORKDIR /app/tracking
RUN ./gradlew clean build


FROM openjdk:8-jdk-alpine
COPY --from=build /app/tracking/build/libs/tracking-0.0.1-SNAPSHOT.jar /app/tracking.jar
EXPOSE 8080
ENTRYPOINT [ "java", "-jar", "/app/tracking.jar" ]