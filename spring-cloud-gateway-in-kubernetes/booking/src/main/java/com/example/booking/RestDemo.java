package com.example.booking;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class RestDemo {

    @GetMapping("/hello")
    public String getBookingMessage(){
        return "Welcome from Booking Application!";
    }
}