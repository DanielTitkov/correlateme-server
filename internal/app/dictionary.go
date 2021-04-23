package app

import (
	"encoding/json"
	"io/ioutil"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

func (a *App) initDictionaries() error {
	err := a.createDictionaries()
	if err != nil {
		return err
	}

	err = a.createDictionaryEntries()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) createDictionaries() error {
	data, err := ioutil.ReadFile(a.cfg.Data.Presets.DictionaryPresetsPath)
	if err != nil {
		return err
	}

	var dictionaries []domain.Dictionary
	err = json.Unmarshal(data, &dictionaries)
	if err != nil {
		return err
	}

	for _, dictionary := range dictionaries {
		d, err := a.repo.GetDictionaryByID(dictionary.ID)
		if err == nil {
			a.logger.Debug("dictionary already exists", d.JSONString())
			continue
		}

		err = a.repo.CreateDictionaty(&dictionary)
		if err != nil {
			return err
		}
		a.logger.Debug("created dictionary", dictionary.JSONString())
	}

	return nil
}

func (a *App) createDictionaryEntries() error {
	data, err := ioutil.ReadFile(a.cfg.Data.Presets.DictionaryEntryPresetsPath)
	if err != nil {
		return err
	}

	var entries []domain.DictionaryEntry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		e, err := a.repo.GetDictionaryEntryByID(entry.ID)
		if err == nil {
			a.logger.Debug("dictionary entry already exists", e.JSONString())
			continue
		}

		err = a.repo.CreateDictionatyEntry(&entry)
		if err != nil {
			return err
		}
		a.logger.Debug("created dictionary entry", entry.JSONString())
	}

	return nil
}
