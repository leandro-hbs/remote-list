package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":5000")
	if err != nil {
		fmt.Print("dialing:", err)
	}

	// Synchronous call
	var reply_i int

	args := []interface{}{10, 0}
	err = client.Call("RemoteList.Append", args, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento adicionado:", reply_i)
	}

	args = []interface{}{20, 0}
	err = client.Call("RemoteList.Append", args, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento adicionado:", reply_i)
	}

	args = []interface{}{10, 1}
	err = client.Call("RemoteList.Append", args, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento adicionado:", reply_i)
	}

	args = []interface{}{1, 0}
	err = client.Call("RemoteList.Get", args, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento encontrado:", reply_i)
	}

	err = client.Call("RemoteList.Remove", 0, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", reply_i)
	}

	err = client.Call("RemoteList.Size", 0, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("O tamanho é: ", reply_i)
	}

}
