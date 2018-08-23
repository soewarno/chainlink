package models

import (
	"fmt"
	"log"
)

func (orm ORM) migrate() {
	orm.initializeModel(&JobSpec{})
	orm.initializeModel(&JobRun{})
	orm.initializeModel(&Initiator{})
	orm.initializeModel(&Tx{})
	orm.initializeModel(&TxAttempt{})
	orm.initializeModel(&BridgeType{})
	orm.initializeModel(&IndexableBlockNumber{})
	orm.initializeModel(&User{})
	orm.initializeModel(&Session{})
	orm.initializeModel(&ServiceAgreement{})
}

func (orm ORM) initializeModel(klass interface{}) {
	err := orm.InitBucket(klass)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to migrate %T: %+v", klass, err))
	}
}

//func performMigrations(db *bolt.DB) error {
//	// 1. If there doesn't exist a schema table / record, create it (do we need to do this?)
//	// 2. Lookup our set of migrations, each migration needs some kind of unique identifier
//	// 3. If a migrations identifier doesn't exist in this schema table, execute it, write the id if it succeeds
//}
//
//func migrateServiceAgreementAmounts(db *bolt.DB) error {
//	// Get all SAs
//
//	type CurrentEncumbrance struct {
//		Payment    *assets.Link `json:"payment"`
//		Expiration uint64       `json:"expiration"`
//	}
//	type CurrentSA struct {
//		CreatedAt   Time               `json:"createdAt" storm:"index"`
//		Encumbrance CurrentEncumbrance `json:"encumbrance" storm:"inline"`
//		ID          string             `json:"id" storm:"id,unique"`
//		JobSpecID   string             `json:"jobSpecID"`
//		RequestBody string             `json:"requestBody"`
//		Signature   string             `json:"signature"`
//		jobSpec     JobSpec            // jobSpec is used during the initial SA creation.
//	}
//
//	type PreviousEncumbrance struct {
//		Payment    *big.Int `json:"payment"`
//		Expiration uint64   `json:"expiration"`
//	}
//	type PreviousSA struct {
//		CreatedAt   Time                `json:"createdAt" storm:"index"`
//		Encumbrance PreviousEncumbrance `json:"encumbrance" storm:"inline"`
//		ID          string              `json:"id" storm:"id,unique"`
//		JobSpecID   string              `json:"jobSpecID"`
//		RequestBody string              `json:"requestBody"`
//		Signature   string              `json:"signature"`
//		jobSpec     JobSpec             // jobSpec is used during the initial SA creation.
//	}
//
//	// Iterate over all SAs
//	// Convert amount from float to string
//	// Save to DB
//}
//
//// GetBolt returns BoltDB from the ORM
//func (orm *ORM) GetBolt() *bolt.DB {
//	return orm.DB.Bolt
//}
//
