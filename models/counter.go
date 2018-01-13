package models

import "log"

type Counter struct {
    Id int64
    Name string
    Count int16
}

func (counter *Counter) Create() {

    statement, err := database.Prepare("INSERT INTO counter (name, count) VALUES (?, ?)")

    if err == nil {
        response, err := statement.Exec(counter.Name, counter.Count)
        if err == nil {
            counter.Id, err = response.LastInsertId()
        }
    }

    if err != nil {
        log.Panic(err)
    }

}

func (counter Counter) Delete() {

}

func (counter *Counter) Increment() {
    counter.Count++
}

func (counter *Counter) Decrement() {
    counter.Count--
}

func (counter *Counter) Update() {

    statement, err := database.Prepare("UPDATE counter SET name = ?, count = ? WHERE id = ?")

    if err == nil {
        statement.Exec(counter.Name, counter.Count, counter.Id)
    } else {
        log.Panic(err)
    }

}
