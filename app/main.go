package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/danieldevpy/bot-golang/app/controller"
	"github.com/danieldevpy/bot-golang/app/database"
	"gorm.io/gorm"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("Aguardando conexões...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Conexão estabelecida com:", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var db *gorm.DB
	var botid int

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Erro ao receber mensagem:", err)
			return
		}
		msg := string(buf[:n])

		if botid != 0 {
			split := strings.Split(msg, "|")
			profile, answer := controller.GetResponse(db, botid, split[0], split[1])
			conn.Write([]byte("" + answer))
			controller.SaveProfile(db, profile)
		} else {
			u64, err := strconv.ParseUint(msg, 10, 32)
			if err != nil {
				fmt.Println(err)
			}
			botid = int(u64)

			db, err = database.ConnectDB()
			if err != nil {
				fmt.Println("error: ", err)
			}

		}

	}
}
