package com.example.java_rest.controller;

import com.example.java_rest.services.MediumTransService;
import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import org.springframework.boot.json.JacksonJsonParser;
import org.springframework.http.codec.json.Jackson2JsonDecoder;
import org.springframework.http.codec.json.Jackson2JsonEncoder;
import org.springframework.web.bind.annotation.*;


@RestController
public class MediumTransController {

//    @PostMapping(path = "/rest/unary/GrayscaleFilter")
//    public @ResponseBody
//    byte[] grayscaleFilter(@RequestBody byte[] barr) throws Exception {
//
////        long init_proc = System.nanoTime() / 1000000;
//
//        byte[] dst_en = new MediumTransService(barr).grayscaleFilter();
//
////        long end_proc = System.nanoTime() / 1000000;
////
////        System.out.printf("Java, REST, Medium, %d, %d\n", dst_en.length, end_proc-init_proc);
//
//        return dst_en;
//
//    }

    @PostMapping(path = "/rest/unary/GrayscaleFilter")
    public @ResponseBody
    byte[] grayscaleFilter(@RequestBody byte[] barr) throws Exception {

//        long init_proc = System.nanoTime() / 1000000;

        ObjectMapper mapper = new ObjectMapper();
        barr = mapper.readValue(barr, barr.getClass());

        byte[] dst_en = new MediumTransService(barr).grayscaleFilter();
        dst_en = mapper.writeValueAsBytes(dst_en);
//        byte[] dst_en = mapper.writeValueAsBytes(barr);

//        long end_proc = System.nanoTime() / 1000000;
//
//        System.out.printf("Java, REST, Medium, %d, %d\n", dst_en.length, end_proc-init_proc);

        return dst_en;

    }

}
