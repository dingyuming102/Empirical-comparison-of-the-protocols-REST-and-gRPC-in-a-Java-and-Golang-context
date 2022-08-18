package com.example.java_rest.controller;


import com.example.java_rest.services.SmallTransService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.Random;


@RestController
public class SmallTransController {

    @Autowired
    private SmallTransService squareService;

    @GetMapping("/rest/unary/square/{number}")
    public int getUnarySquare(@PathVariable int number){
//        long init_proc = System.nanoTime();

        int retData = this.squareService.getUnarySquare(number);

//        long end_proc = System.nanoTime();

//        System.out.printf("Java, REST, Small, %d, %d\n", retData, end_proc-init_proc);

        return retData;
    }

    @PostMapping(path ="/rest/unary/square/")
    public @ResponseBody
    int postUnarySquare(@RequestBody int number){
//        long init_proc = System.nanoTime();

//        int retData = this.squareService.getUnarySquare(number);

//        long end_proc = System.nanoTime();

//        System.out.printf("Java, REST, Small, %d, %d\n", retData, end_proc-init_proc);

        return number*number;
    }
}