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

package storobj

import (
	"encoding/binary"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func ParseAndExtractProperty(data []byte, propName string) ([]string, bool, error) {
	if propName == "id" || propName == "_id" {
		return extractID(data)
	}
	return ParseAndExtractTextProp(data, propName)
}

func ParseAndExtractTextProp(data []byte, propName string) ([]string, bool, error) {
	vals := []string{}
	err := parseAndExtractValueProp(data, propName, func(value []byte) {
		vals = append(vals, string(value))
	})
	if err != nil {
		return nil, false, err
	}
	return vals, true, nil
}

func ParseAndExtractNumberArrayProp(data []byte, propName string) ([]float64, bool, error) {
	vals := []float64{}
	err := parseAndExtractValueProp(data, propName, func(value []byte) {
		vals = append(vals, mustExtractNumber(value))
	})
	if err != nil {
		return nil, false, err
	}
	return vals, true, nil
}

func parseAndExtractValueProp(data []byte, propName string, valueFn func(value []byte)) error {
	propsBytes, err := extractPropsBytes(data)
	if err != nil {
		return err
	}

	val, t, _, err := jsonparser.Get(propsBytes, propName)
	if err != nil {
		return err
	}

	if t == jsonparser.Array {
		jsonparser.ArrayEach(val, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			valueFn(value)
		})
	} else {
		valueFn(val)
	}

	return nil
}

func mustExtractNumber(value []byte) float64 {
	number, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		panic("not a float64")
	}
	return number
}

func extractID(data []byte) ([]string, bool, error) {
	start := 1 + 8 + 1
	end := start + 16
	if len(data) > end {
		uuidParsed, err := uuid.FromBytes(data[start:end])
		if err != nil {
			return nil, false, errors.New("cannot parse id property")
		}
		return []string{uuidParsed.String()}, true, nil
	}
	return nil, false, errors.New("id property not found")
}

func extractPropsBytes(data []byte) ([]byte, error) {
	version := uint8(data[0])
	if version != 1 {
		return nil, errors.Errorf("unsupported binary marshaller version %d", version)
	}

	vecLen := binary.LittleEndian.Uint16(data[discardBytesPreVector : discardBytesPreVector+2])

	classNameStart := discardBytesPreVector + 2 + vecLen*4

	classNameLen := binary.LittleEndian.Uint16(data[classNameStart : classNameStart+2])

	propsLenStart := classNameStart + 2 + classNameLen
	propsLen := binary.LittleEndian.Uint32(data[propsLenStart : propsLenStart+4])

	start := int64(propsLenStart + 4)
	end := start + int64(propsLen)

	return data[start:end], nil
}

const discardBytesPreVector = 1 + 8 + 1 + 16 + 8 + 8
