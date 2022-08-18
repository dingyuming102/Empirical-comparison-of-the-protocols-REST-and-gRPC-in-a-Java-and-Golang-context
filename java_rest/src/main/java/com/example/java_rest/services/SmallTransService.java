package com.example.java_rest.services;

import org.springframework.stereotype.Service;

@Service
public class SmallTransService {

    public int getUnarySquare(int number) {
        return number * number;
    }

}
