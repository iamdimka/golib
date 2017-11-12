package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

func ExendFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	rx := regexp.MustCompile("\\s*(\\S*)\\s*=\\s*(\\S*)\\s*")

	for _, line := range rx.FindAllSubmatch(file, -1) {
		os.Setenv(string(line[1]), string(line[2]))
	}
	return nil
}

func Read(p interface{}, name string) error {
	raw := os.Getenv(name)

	switch p := p.(type) {
	case *bool:
		v, err := strconv.ParseBool(name)
		*p = v
		return err

	case *string:
		*p = raw
		return nil

	case *uint8:
		v, err := strconv.ParseUint(raw, 10, 8)
		*p = uint8(v)
		return err

	case *uint16:
		v, err := strconv.ParseUint(raw, 10, 16)
		*p = uint16(v)
		return err

	case *uint32:
		v, err := strconv.ParseUint(raw, 10, 32)
		*p = uint32(v)
		return err

	case *uint64:
		v, err := strconv.ParseUint(raw, 10, 64)
		*p = uint64(v)
		return err

	case *int8:
		v, err := strconv.ParseInt(raw, 10, 8)
		*p = int8(v)
		return err

	case *int16:
		v, err := strconv.ParseInt(raw, 10, 16)
		*p = int16(v)
		return err

	case *int32:
		v, err := strconv.ParseInt(raw, 10, 32)
		*p = int32(v)
		return err

	case *int64:
		v, err := strconv.ParseInt(raw, 10, 64)
		*p = int64(v)
		return err

	case *float32:
		v, err := strconv.ParseFloat(raw, 32)
		*p = float32(v)
		return err

	case *float64:
		v, err := strconv.ParseFloat(raw, 64)
		*p = float64(v)
		return err

	}

	return nil
}

func FillStruct(p interface{}) (err error) {
	rv := reflect.ValueOf(p)

	if rv.Kind() != reflect.Ptr {
		err = fmt.Errorf("Should be a pointer to struct")
		return
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		err = fmt.Errorf("Should be a pointer to struct")
		return
	}

	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		tag := rt.Field(i).Tag
		name := tag.Get("env")
		if name == "" {
			continue
		}

		fv := rv.Field(i)
		if fv.Kind() != reflect.Ptr {
			fv = fv.Addr()
		}

		err = Read(fv.Interface(), name)
		if err != nil {
			return
		}
	}

	return
}
