package main 

import (
	"sync"
	//"fmt"
)

type cacheShard struct {
	items map[uint64]interface{}
	sync.RWMutex
	maxSize int
}

/*
func (s *cacheShard) Lock(){
	 fmt.Println("Lock of cacheShard");
}

func (s *cacheShard) Unlock(){
     fmt.Println("UnLock of cacheShard"); 
}

func (s *cacheShard) Lock() bool{
     fmt.Println("Lock1 of cacheShard");
	 return true;
}

func (s *cacheShard) Unlock() bool{
     fmt.Println("UnLock1 of cacheShard"); 
	 return true;
}
*/

func (s *cacheShard) add() bool {
	s.Lock() //call the "sync.RWMutex" method dirctly"" 
	defer s.Unlock() //call the "sync.RWMutex" method dirctly
	//s.mtx.Lock();
	//defer s.mtx.Unlock();
	return true;
}

func main() {
	var s cacheShard;
	s.add();
}
