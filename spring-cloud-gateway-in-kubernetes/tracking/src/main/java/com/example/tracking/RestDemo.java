package com.example.tracking;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class RestDemo {

    @GetMapping("/hello")
    public String getBookingMessage(){
        return "Welcome from Tracking Application!";
    }
}