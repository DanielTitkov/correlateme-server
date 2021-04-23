package entgo

import (
	"context"
	"errors"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dictionary"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dictionaryentry"
)

func (r *EntgoRepository) CreateDictionaty(d *domain.Dictionary) error {
	_, err := r.client.Dictionary.
		Create().
		SetID(d.ID).
		SetCode(d.Code).
		SetDescription(d.Description).
		Save(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (r *EntgoRepository) GetDictionaryByCode(code string) (*domain.Dictionary, error) {
	dict, err := r.client.Dictionary.Query().
		Where(dictionary.CodeEQ(code)).
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainDictionary(dict), nil
}

func (r *EntgoRepository) GetDictionaryByID(id int) (*domain.Dictionary, error) {
	dict, err := r.client.Dictionary.Query().
		Where(dictionary.IDEQ(id)).
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainDictionary(dict), nil
}

func (r *EntgoRepository) CreateDictionatyEntry(d *domain.DictionaryEntry) error {
	_, err := r.client.DictionaryEntry.
		Create().
		SetID(d.ID).
		SetDictionaryID(d.DictionaryID).
		SetCode(d.Code).
		SetDescription(d.Description).
		Save(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (r *EntgoRepository) GetDictionaryEntryByID(id int) (*domain.DictionaryEntry, error) {
	entry, err := r.client.DictionaryEntry.Query().
		Where(dictionaryentry.IDEQ(id)).
		WithDictionary().
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	if entry.Edges.Dictionary == nil {
		return nil, errors.New("orphan dictionary entry")
	}

	return entToDomainDictionaryEntry(entry), nil
}

func entToDomainDictionary(d *ent.Dictionary) *domain.Dictionary {
	return &domain.Dictionary{
		ID:          d.ID,
		Code:        d.Code,
		Description: d.Description,
	}
}

func entToDomainDictionaryEntry(e *ent.DictionaryEntry) *domain.DictionaryEntry {
	return &domain.DictionaryEntry{
		DictionaryID: e.Edges.Dictionary.ID,
		ID:           e.ID,
		Code:         e.Code,
		Description:  e.Description,
	}
}
