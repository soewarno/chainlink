package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/asdine/storm"
	"github.com/smartcontractkit/chainlink/logger"
	"github.com/smartcontractkit/chainlink/store"
	"github.com/smartcontractkit/chainlink/store/models"
)

// Reaper interface is a gateway to an instance that can reap stale objects such as sessions.
type Reaper interface {
	Start() error
	Stop() error
	ReapSessions()
}

type storeReaper struct {
	store    *store.Store
	config   store.Config
	cond     *sync.Cond
	mutex    *sync.Mutex
	sessions int
}

// NewStoreReaper creates a reaper that cleans stale objects from the store.
func NewStoreReaper(store *store.Store) Reaper {
	mutex := &sync.Mutex{}
	sr := &storeReaper{
		store:  store,
		config: store.Config,
		cond:   sync.NewCond(mutex),
		mutex:  mutex,
	}
	go sr.reaperLoop()
	return sr
}

// Start starts the reaper instance so that it can listen for cleanup asynchronously.
func (sr *storeReaper) Start() error {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()

	sr.sessions++
	sr.cond.Signal()
	return nil
}

// Stop stops the reaper from listening to clean up messages asynchronously.
func (sr *storeReaper) Stop() error {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()

	sr.sessions = -1
	sr.cond.Signal()
	return nil
}

// ReapSessions signals the reaper to clean up sessions asynchronously.
func (sr *storeReaper) ReapSessions() {
	//go sr.reaperLoop()
}

func (sr *storeReaper) reaperLoop() {
	for {
		sr.mutex.Lock()
		for sr.sessions == 0 {
			sr.cond.Wait()
		}
		if sr.sessions < 0 {
			return
		}
		sr.deleteStaleSessions()
		sr.sessions--
		fmt.Println("reaperLoop")
		sr.mutex.Unlock()
	}
}

func (sr *storeReaper) deleteStaleSessions() {
	var sessions []models.Session
	offset := time.Now().Add(-sr.config.ReaperExpiration.Duration).Add(-sr.config.SessionTimeout.Duration)
	stale := models.Time{offset}
	err := sr.store.Range("LastUsed", models.Time{}, stale, &sessions)
	if err == storm.ErrNotFound {
		return
	} else if err != nil {
		logger.Error("unable to reap stale sessions: ", err)
		return
	}

	for _, s := range sessions {
		err := sr.store.DeleteStruct(&s)
		if err != nil {
			logger.Error("unable to delete stale session: ", err)
		}
	}
}
