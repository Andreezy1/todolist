package todolist

import (
	"time"

	"github.com/k0kubun/pp"
)

type ToDoList struct {
	Opisanie    string
	Done        bool
	TimeNachalo time.Time
	TimeKonec   time.Time
}

var zadachi = make(map[string]ToDoList)

func Todolist(text []string) string {
	_, ok := zadachi[text[1]]
	if ok {
		pp.Println(HaveZadacha)
		return HaveZadacha
	}
	opis := ""
	for i := 2; i < len(text); i++ {
		opis += text[i]
		if i != len(text)-1 {
			opis += " "
		}
	}
	zadachi[text[1]] = ToDoList{
		Opisanie:    opis,
		Done:        false,
		TimeNachalo: time.Now(),
	}
	_, ok = zadachi[text[1]]
	if ok {
		pp.Println("Задача добавлена")
	}
	return ""
}

func Alltodolist() string {
	if len(zadachi) == 0 {
		pp.Println(donthaveZadacha)
		return donthaveZadacha
	}
	for i, v := range zadachi {
		pp.Println("___________________________")
		pp.Println("Заголовок:", i)
		pp.Println("Описание:", v.Opisanie)
		pp.Println("Выполнено:", v.Done)
		pp.Println("Начало задачи:", v.TimeNachalo)
		if v.Done {
			pp.Println("Конец задачи:", v.TimeKonec)
		}
		pp.Println("___________________________")
	}
	return ""
}

func TodolistDone(s string) string {
	v, ok := zadachi[s]
	if ok {
		v.Done = true
		v.TimeKonec = time.Now()
		zadachi[s] = v
		pp.Println("Статус задачи изменен")
		return ""
	}
	pp.Println(Notfound)
	return Notfound
}

func Deletetodolist(s string) string {
	_, ok := zadachi[s]
	if ok {
		delete(zadachi, s)
		pp.Println("Задача удалена")
		return ""
	}
	pp.Println(Notfound)
	return Notfound
}
