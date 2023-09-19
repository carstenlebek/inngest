// Code generated by "enumer -trimprefix=Period -type=Period -json -gql -sql -text -transform=snake"; DO NOT EDIT.

package enums

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

const _PeriodName = "noneminutehourdayweekmonth"

var _PeriodIndex = [...]uint8{0, 4, 10, 14, 17, 21, 26}

func (i Period) String() string {
	if i < 0 || i >= Period(len(_PeriodIndex)-1) {
		return fmt.Sprintf("Period(%d)", i)
	}
	return _PeriodName[_PeriodIndex[i]:_PeriodIndex[i+1]]
}

var _PeriodValues = []Period{0, 1, 2, 3, 4, 5}

var _PeriodNameToValueMap = map[string]Period{
	_PeriodName[0:4]:   0,
	_PeriodName[4:10]:  1,
	_PeriodName[10:14]: 2,
	_PeriodName[14:17]: 3,
	_PeriodName[17:21]: 4,
	_PeriodName[21:26]: 5,
}

// PeriodFromString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PeriodFromString(s string) (Period, error) {
	if val, ok := _PeriodNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Period values", s)
}

// PeriodValues returns all values of the enum
func PeriodValues() []Period {
	return _PeriodValues
}

// IsAPeriod returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Period) IsAPeriod() bool {
	for _, v := range _PeriodValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Period
func (i Period) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Period
func (i *Period) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Period should be a string, got %s", data)
	}

	var err error
	*i, err = PeriodFromString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for Period
func (i Period) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Period
func (i *Period) UnmarshalText(text []byte) error {
	var err error
	*i, err = PeriodFromString(string(text))
	return err
}

func (i Period) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Period) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := PeriodFromString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

func (i *Period) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		bytes, ok := v.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := PeriodFromString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

func (i Period) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(i.String())))
}