package main

import (
	"flag"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/rlcao/learn-golang/protobuf/todo"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
    flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr,"missing subcommand: list or add")
		os.Exit(1)
	}

    var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:]," "))
	default:
		err = fmt.Errorf("unknown subcommand: %v",cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr,"got error: %v",err)
	}
	fmt.Println("TODO")
}

const dbPath = "mydb.pb"

func add(text string) error  {
	task := &todo.Task{
		Text:text,
		Done: false,
	}
	fmt.Printf(proto.MarshalTextString(task))

	bf,err := proto.Marshal(task)
	if err != nil {
		return fmt.Errorf("got error: %v", err)
	}

	f, err := os.OpenFile(dbPath,os.O_WRONLY|os.O_CREATE | os.O_APPEND, 0666)
	if err != err {
		return fmt.Errorf("could not open db file %v: %v",dbPath,err)
	}
	defer f.Close()
	f.Write(bf)
	return nil
}

func list() error {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("failed to list: %v",err)
	}
	for {
		var task todo.Task
		err := proto.Unmarshal(b,&task);
		if err == io.EOF {
			return nil
		} else if err != nil {
			return fmt.Errorf("failed")
		}
		if task.Done {
			fmt.Printf("Great!")
		} else {
			fmt.Printf("Bad!")
		}
		fmt.Printf("\n%v",task.Text)
	}

	return nil
}