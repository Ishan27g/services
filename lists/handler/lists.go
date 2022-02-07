package handler

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
	"github.com/micro/services/pkg/tenant"
	lists "lists/proto"
)

type Lists struct{}

func (e *Lists) New(ctx context.Context, req *lists.List, rsp *lists.Response) error {
	log.Info("Received Lists.New request")
	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	// generate a key (uuid v4)
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%s:%s", tnt, id)

	req.Id = id.String()
	store.NewRecord(key, req)

	rsp.Msg = "Added ----" + req.Id
	return nil
}
func (e *Lists) Update(ctx context.Context, req *lists.List, rsp *lists.Response) error {
	log.Info("Received Lists.update request")

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}
	key := fmt.Sprintf("%s:%s", tnt, req.Id)

	// read the specific note
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("lists.update", "List not found")
	} else if err != nil {
		return errors.InternalServerError("lists.update", "Error reading from store: %v", err.Error())
	}

	// Decode the note
	var list *lists.List
	if err := recs[0].Decode(&list); err != nil {
		return errors.InternalServerError("lists.update", "Error unmarshaling JSON: %v", err.Error())
	}

	list.Content = req.Content

	rec := store.NewRecord(key, list)

	// Write the updated note to the store
	if err = store.Write(rec); err != nil {
		return errors.InternalServerError("lists.update", "Error writing to store: %v", err.Error())
	}

	rsp.Msg = "Updated " + req.String()
	return nil
}
func (e *Lists) Delete(ctx context.Context, opts *lists.GetOptions, rsp *lists.Response) error {
	log.Info("Received Lists.Delete request")

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}
	for _, id := range opts.Id {

		key := fmt.Sprintf("%s:%s", tnt, id)

		// read the specific note
		recs, err := store.Read(key)
		if err == store.ErrNotFound {
			return nil
		} else if err != nil {
			return errors.InternalServerError("lists.delete", "Error reading from store: %v", err.Error())
		}

		// Decode the note
		var list *lists.List
		if err := recs[0].Decode(&list); err != nil {
			return errors.InternalServerError("lists.delete", "Error unmarshaling JSON: %v", err.Error())
		}

		// now delete it
		if err := store.Delete(key); err != nil && err != store.ErrNotFound {
			return errors.InternalServerError("lists.delete", "Failed to delete note")
		}
	}

	rsp.Msg = "Deleted " + opts.String()
	return nil
}
func (e *Lists) Get(ctx context.Context, opts *lists.GetOptions, rsp *lists.AllLists) error {
	log.Info("Received Lists.Get request")
	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	// for _, id := range opts.Id {
	key := fmt.Sprintf("%s:%s", tnt, opts.Id)
	log.Info("retrieving ", key)

	// read the specific list
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("lists.read", "List not found")
	} else if err != nil {
		return errors.InternalServerError("lists.read", "Error reading from store: %v", err.Error())
	}

	// Decode the list
	var list *lists.List
	if err := recs[0].Decode(&list); err != nil {
		return errors.InternalServerError("notes.update", "Error unmarshaling JSON: %v", err.Error())
	}

	rsp.Lists = append(rsp.Lists, list)
	//	}

	return nil
}
