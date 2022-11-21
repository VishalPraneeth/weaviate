//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InvertedIndexConfig Configure the inverted index built into Weaviate
//
// swagger:model InvertedIndexConfig
type InvertedIndexConfig struct {

	// bm25
	Bm25 *BM25Config `json:"bm25,omitempty"`

	// Asynchronous index clean up happens every n seconds
	CleanupIntervalSeconds int64 `json:"cleanupIntervalSeconds,omitempty"`

	// hybrid search
	HybridSearch *HybridConfig `json:"hybridSearch,omitempty"`

	// Index each object with the null state
	IndexNullState bool `json:"indexNullState,omitempty"`

	// Index length of properties
	IndexPropertyLength bool `json:"indexPropertyLength,omitempty"`

	// Index each object by its internal timestamps
	IndexTimestamps bool `json:"indexTimestamps,omitempty"`

	// stopwords
	Stopwords *StopwordConfig `json:"stopwords,omitempty"`
}

// Validate validates this inverted index config
func (m *InvertedIndexConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBm25(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHybridSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopwords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvertedIndexConfig) validateBm25(formats strfmt.Registry) error {

	if swag.IsZero(m.Bm25) { // not required
		return nil
	}

	if m.Bm25 != nil {
		if err := m.Bm25.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bm25")
			}
			return err
		}
	}

	return nil
}

func (m *InvertedIndexConfig) validateHybridSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.HybridSearch) { // not required
		return nil
	}

	if m.HybridSearch != nil {
		if err := m.HybridSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hybridSearch")
			}
			return err
		}
	}

	return nil
}

func (m *InvertedIndexConfig) validateStopwords(formats strfmt.Registry) error {

	if swag.IsZero(m.Stopwords) { // not required
		return nil
	}

	if m.Stopwords != nil {
		if err := m.Stopwords.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stopwords")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvertedIndexConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvertedIndexConfig) UnmarshalBinary(b []byte) error {
	var res InvertedIndexConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
