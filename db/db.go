package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	localHostDB = "mongodb://localhost:27017"
	mlabHost    = "mongodb://archiver:!2017Dlab@ds155737.mlab.com:55737/draglabsdev"
	dbName      = "dsound"
	userC       = "users"
	jamC        = "jams"
	recordings  = "recordings"
	//mongodb://marlon:4803marlon@ds035856.mlab.com:35856/draglabs
)

// DataStore struct holds our DB interaction
type DataStore struct {
	session *mgo.Session
}

//Collection func, connect to DB
func (ds *DataStore) Collection(name string) *mgo.Collection {

	return ds.session.DB(dbName).C(name)
}

// UserCollection return our user Collection on DB
func (ds *DataStore) UserCollection() *mgo.Collection {
	return ds.session.DB(dbName).C(userC)
}

// JamCollection func, gives us a new jam collection
// is a connvenience func  for `session.DB(dbName).C(cName)`
func (ds *DataStore) JamCollection() *mgo.Collection {

	return ds.session.DB(dbName).C(jamC)
}
func (ds *DataStore) RecordingsCollection() *mgo.Collection {
	return ds.session.DB(dbName).C(recordings)
}

// Close func closes our session on DB
func (ds *DataStore) Close() {
	ds.session.Close()
}

// NewDataStore func, returns our new store
func NewDataStore() *DataStore {
	info := mgo.DialInfo{
		Addrs:    []string{"54.183.100.139:27017"},
		Database: dbName,
		Username: "soundBoy",
		Password: "soundBoy",
	}
	session, err := mgo.DialWithInfo(&info)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	session.SetMode(mgo.Monotonic, true)

	return &DataStore{session: session.Copy()}
}
