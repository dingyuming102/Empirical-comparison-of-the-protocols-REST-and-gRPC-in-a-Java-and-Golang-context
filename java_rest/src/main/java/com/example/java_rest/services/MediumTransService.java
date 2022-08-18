package com.example.java_rest.services;

import org.opencv.core.Core;
import org.opencv.core.Mat;
import org.opencv.core.MatOfByte;
import org.opencv.imgcodecs.Imgcodecs;
import org.opencv.imgproc.Imgproc;

public class MediumTransService {

    byte[] buf;

    public MediumTransService(byte[] buf)  {
        this.buf = buf;
    }

    public byte[] grayscaleFilter() {

        //nu.pattern.OpenCV.loadShared();
        nu.pattern.OpenCV.loadLocally();

        /* Receive and decode input */
        Mat src = Imgcodecs.imdecode(new MatOfByte(this.buf), Imgcodecs.IMREAD_UNCHANGED);
        Mat dest = new Mat();

        /* Apply filter */
        //long init_op = System.nanoTime();
        Imgproc.cvtColor(src, dest, Imgproc.COLOR_RGB2GRAY, src.channels());
        //long end_op = System.nanoTime();

        /* Decode and create output */
        MatOfByte dest_img = new MatOfByte();
        Imgcodecs.imencode(".jpg", dest, dest_img);

        return dest_img.toArray();
    }
}
