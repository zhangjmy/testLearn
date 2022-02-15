/**
 * File: x
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/9/22 8:36 下午
 **/

package main

import (
"fmt"
"net"
)

type Reader struct {
	Conn net.Conn
	Buff []byte
	Start int    //数据读取开始位置
	End int    //数据读取结束位置
	BuffLen int    //数据接收缓冲区大小
	HeaderLen int    //包头长度
}

//读取tcp数据流
func (reader *Reader)Read(acceptData chan string) error {
	for {
		reader.Move()
		if reader.End == reader.BuffLen {
			//缓冲区的宽度容纳不了一条消息的长度
			return fmt.Errorf("one message is too large:%v", reader)
		}
		len, err := reader.Conn.Read(reader.Buff[reader.End:])
		if err != nil {
			return err
		}
		reader.End += len
		reader.GetData(acceptData)
	}
}

//前移上一次未处理完的数据
func (reader *Reader)Move() {
	if reader.Start == 0 {
		return
	}
	copy(reader.Buff, reader.Buff[reader.Start:reader.End])
	reader.End -= reader.Start
	reader.Start = 0
}

//读取buff中的单条数据
func (reader *Reader)GetData(acceptData chan string) error {
	if reader.End - reader.Start < reader.HeaderLen {
		//包头的长度不够，继续接收
		return nil
	}
	//读取包头数据
	headerData := reader.Buff[reader.Start:(reader.Start + reader.HeaderLen)]
	if reader.End - reader.Start - reader.HeaderLen < int(headerData[reader.HeaderLen - 1]) {
		//包体的长度不够，继续接收
		return nil
	}
	//读取包体数据
	bodyData := reader.Buff[(reader.Start + reader.HeaderLen):reader.Start + reader.HeaderLen + int(headerData[reader.HeaderLen - 1])]
	//把完整的包用通道传递出去
	acceptData <- string(append(headerData, bodyData...))
	//每读完一次数据 start 后移
	reader.Start += reader.HeaderLen + int(headerData[reader.HeaderLen - 1])
	return reader.GetData(acceptData)
}
