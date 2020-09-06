package main

import (
	"fmt"
	"time"
	"errors"
	
	"github.com/keonjeo/taskflow"
)

func SerialAndParallelTaskDemo() {
	revertMethod := func(param interface{}) error {
		fmt.Printf("revertMethod => [%v] Oh My God! Something goes wrong here!\n", param)
		return nil
	}

	sayHi := func(param interface{}) error {
		for i := 0; i < 10; i++ {
			fmt.Printf("Hi, I am %v\n", param)
		}

		return nil
	}
	task1 := taskflow.NewTask("sayHi Task", "Tony", sayHi, revertMethod)

	sayGoodBye := func(param interface{}) error {
		time.Sleep(1 * time.Second)
		for i := 0; i < 10; i++ {
			fmt.Printf("goodbye %v\n", param)
			time.Sleep(10 * time.Millisecond)
		}
		return errors.New("xx")
	}
	task2 := taskflow.NewTask("sayGoodBye Task", "Steven", sayGoodBye, revertMethod)

	serialTaskExecuteMethod := func(param interface{}) error {
		time.Sleep(3 * time.Second)
		for i := 0; i < 10; i++ {
			fmt.Printf("Hi %v\n", param)
			time.Sleep(100 * time.Millisecond)
		}
		return errors.New("xx")
	}

	serialTaskRevertMethod := func(param interface{}) error {
		time.Sleep(3 * time.Second)
		for i := 0; i < 10; i++ {
			fmt.Printf("Goodbye %v\n", param)
			time.Sleep(100 * time.Millisecond)
		}
		return nil
	}

	serialTask := taskflow.NewSerialTask("SerialTask Demo", "Pony", serialTaskExecuteMethod, serialTaskRevertMethod, []taskflow.ITask{task1, task2})
	serialTask.Execute()
}
