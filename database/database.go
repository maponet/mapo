package database

import (
    "mapo/log"
    
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

var database *mgo.Database

// TODO: definire una funzione che si occupa con la creazione e gestione della
// connessione verso un database.
func NewConnection(databaseName string) {
    log.Debug("executing NewConnection function")
    
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    database = session.DB(databaseName)
    
    // connessione alla data base avvenne usando diversi livelli di autenticazione
    // come admin, user, ... e probabile altri?
}

// Store salva nella database un singolo oggetto
func Store(data interface{}, table string) error {
    
    c := database.C(table)
    
    err := c.Insert(data)
    
    return err
}

// RestoreOne riprende dalla database un singolo oggetto identificato da un id
func RestoreOne(data interface{}, id string, table string) error {
    
    c := database.C(table)
    
    err := c.Find(bson.M{"_id" : id}).One(data)
    
    return err
}

// RestoreList riprende dalla database una lista (tutti) di oggetti, senza alcun filtro
func RestoreList(data interface{}, table string) error {
    
    c := database.C(table)
    
    err := c.Find(bson.M{}).All(data)
    
    return err
}

// Update aggiorna i valori di un oggetto nella database, identificato da un id
func Update(data interface{}, id string, table string) error {
    
    c := database.C(table)
    
    err := c.Update(bson.M{"_id": id}, data)
    
    return err
}