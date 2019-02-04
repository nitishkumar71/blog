package com.example.sampleservice.demo;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/hello")
public class HelloWorldRestController{

    @GetMapping("/message")
    public String get_hello_message(){
        return "Hello World!!!";
    }
}