package gc_io

import (
	"encoding/binary"
	"fmt"
	"os"
)

type BitmapHeader struct {
	HeadA     byte   // B
	HeadB     byte   // M
	Size      uint32 // 文件大小
	ReservedA uint16 // 0
	ReservedB uint16 // 0
	OffBits   uint32 // 数据偏移
}

type BitmapInfoHeader struct {
	Size           uint32 // 结构体大小
	Width          int32  // 宽度
	Height         int32  // 高度
	Planes         uint16 // 面， 恒定为1
	BitCount       uint16 // 每个像素占用的字节数
	Compression    uint32 // 压缩类型
	SizeImage      uint32 // 图形大小
	XPerlsPerMeter int32  // 水平分辨率 每米的像素数
	YPerlsPerMeter int32  // 每米的像素数
	ClrUsed        uint32 // 颜色数
	ClrImportant   uint32 // 调色版
}

func BinaryUsage() {

	file, err := os.Open("image.bmp")

	if err != nil {
		fmt.Println(err)
		return
	}

	var headA, headB byte
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	var reserveA, reserveB uint16
	binary.Read(file, binary.LittleEndian, &reserveA)
	binary.Read(file, binary.LittleEndian, &reserveB)

	var offbits uint32
	binary.Read(file, binary.LittleEndian, &offbits)

	fmt.Println(headA, headB, size, reserveA, reserveB, offbits)

	infoHeader := new(BitmapInfoHeader)
	if err := binary.Read(file, binary.LittleEndian, infoHeader); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(infoHeader)

	//	/fmt.Println("size", binary.Size(header), binary.Size(infoHeader))
}
