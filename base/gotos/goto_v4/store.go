package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

const saveQueueLength = 1000

type URLStore struct {
	urls map[string]string

	mu sync.RWMutex

	// 存储到文件中
	save chan record
}

type record struct {
	Key, URL string
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),
	}

	if err := s.load(filename); err != nil {
		log.Println("Error loading URLStore:", err)
	}
	go s.saveLoop(filename)
	return s
}

func (s *URLStore) Get(key string) string  {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool{
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, present := s.urls[key]; present {
		return false
	}

	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if ok := s.Set(key, url); ok {

			// 存储到通道中
			s.save <- record{key, url}
			return key
		}
	}
	panic("shouldn't get here")
}

//https://blog.csdn.net/qq_40500045/article/details/106220410

func (s *URLStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}
	defer f.Close()

	d:= json.NewDecoder(f)

	for err == nil{
		var r record
		// err 局部变量 如果在解码中出错 退出当前循环
		if err = d.Decode(&r); err == nil {
			log.Printf("loading ... key : %s  values : %s ", r.Key, r.URL)
			s.Set(r.Key, r.URL)
		}
	}

	if err == io.EOF {
		return nil
	}

	// error occurred:
	log.Println("Error decoding URLStore:", err) // map hasn't been read correctly
	return err
}

func (s *URLStore) saveLoop(filename string) error {

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore: ", err)
	}
	defer f.Close()

	e := json.NewEncoder(f)
	for  {
		r := <- s.save
		if err := e.Encode(r); err != nil{
			log.Println("Error saving to URLStore: ", err)
		}
	}
}