package utils

import (
	"gocv.io/x/gocv"
	"os"
	"strconv"
)

func MediumTransInitData(fileName string, size int) [][]byte {

	color_img := gocv.IMRead(fileName, gocv.IMReadUnchanged)
	req_img, err := gocv.IMEncode(".jpg", color_img)
	if err != nil {
		panic(err)
	}

	//data := make([][]byte, size)
	//for i := range data {
	//	data[i] = req_img // 每一行4列
	//}
	data := make([][]byte, 1)
	data[0] = req_img

	return data
}

func GrayscaleFilter(buf []byte) []byte {
	// Decode bytes sequence received
	src, _ := gocv.IMDecode(buf, gocv.IMReadUnchanged)
	dst := gocv.NewMat()

	// Apply Filter
	//init_op := time.Now().UnixNano() / int64(time.Millisecond)
	gocv.CvtColor(src, &dst, gocv.ColorRGBToGray)
	//dst = src
	//end_op := time.Now().UnixNano() / int64(time.Millisecond)

	// Encode bytes sequence result
	dst_en, _ := gocv.IMEncode(".jpg", dst)
	return dst_en
}

func DecodeAndSave(id int, buf []byte) bool {
	_, err := gocv.IMDecode(buf, gocv.IMReadUnchanged)
	if err != nil {
		panic(err)
	}

	//return SaveImage(id, gray_img)
	return true
}

func SaveImage(id int, to_save gocv.Mat) bool {
	mob, err := gocv.IMEncode(".jpg", to_save)
	if err != nil {
		panic(err)
		return false
	}
	ba := mob

	if err := os.WriteFile("asserts/pic/gray"+strconv.Itoa(id)+".jpg", ba, 0644); err != nil {
		panic(err)
	}

	return true
}
