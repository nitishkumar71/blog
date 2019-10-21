FROM openjdk:8-jdk-alpine as build
COPY . /app/booking
WORKDIR /app/booking
RUN ./gradlew clean build


FROM openjdk:8-jdk-alpine
COPY --from=build /app/booking/build/libs/booking-0.0.1-SNAPSHOT.jar /app/booking.jar
EXPOSE 8080
ENTRYPOINT [ "java", "-jar", "/app/booking.jar" ]