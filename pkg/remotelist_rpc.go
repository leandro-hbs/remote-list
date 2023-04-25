package remotelist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
)

type RemoteList struct {
	mu []sync.Mutex
	matriz [][]int
}

func (l *RemoteList) Append(args []interface{}, reply *int) error {
	value := args[0].(int)
    list_id := args[1].(int)
	if list_id >= len(l.matriz) {
        l.matriz = append(l.matriz, make([]int, 0))
		l.mu = append(l.mu, sync.Mutex{})
    }
	l.mu[list_id].Lock()
    defer l.mu[list_id].Unlock()
	l.matriz[list_id] = append(l.matriz[list_id], value)
	fmt.Println("Cliente: ", list_id)
	fmt.Println("Lista: ", l.matriz[list_id])
	fmt.Println("")
	l.UpdateJson()
	*reply = value
	return nil
}

func (l *RemoteList) Get(args []interface{}, reply *int) error {
	i := args[0].(int)
    list_id := args[1].(int)
	l.mu[list_id].Lock()
    defer l.mu[list_id].Unlock()
    if list_id >= len(l.matriz) {
        return errors.New("invalid list id")
    }
    targetList := l.matriz[list_id]
    if i >= len(targetList) {
        return errors.New("invalid value ID")
    }
    *reply = targetList[i]
    return nil
}

func (l *RemoteList) Remove(list_id int, reply *int) error {
	l.mu[list_id].Lock()
    defer l.mu[list_id].Unlock()
	if len(l.matriz[list_id]) > 0 {
		l.matriz[list_id] = l.matriz[list_id][:len(l.matriz[list_id])-1]
		fmt.Println("Cliente: ", list_id)
		fmt.Println("Lista: ", l.matriz[list_id])
		fmt.Println("")
		*reply = l.matriz[list_id][len(l.matriz[list_id])-1]
		l.UpdateJson()
	} else {
		return errors.New("empty list")
	}
	return nil
}

func (l *RemoteList) Size(list_id int, reply *int) error {
	l.mu[list_id].Lock()
    defer l.mu[list_id].Unlock()
    if list_id >= len(l.matriz) {
        return errors.New("invalid list id")
    }
    targetList := l.matriz[list_id]
    *reply = len(targetList)
    return nil
}

func (l *RemoteList) LoadMatriz() error {
	data, err := ioutil.ReadFile("matriz.json")
	if err != nil {
		return err
	}
	var matriz [][]int
	err = json.Unmarshal(data, &matriz)
	if err != nil {
		return err
	}
	l.matriz = matriz
	l.mu = make([]sync.Mutex, len(l.matriz))
	return nil
}

func (l *RemoteList) UpdateJson() error {
	jsonData, err := json.Marshal(l.matriz)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("matriz.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func NewRemoteList() *RemoteList {
	return new(RemoteList)
}
