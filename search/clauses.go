package search

import (
	"encoding/json"
	"fmt"

	"github.com/chanced/dynamic"
)

type Clauses []Clause

func (c Clauses) MarshalJSON() ([]byte, error) {
	u, err := c.unpack()
	if err != nil {
		return nil, err
	}
	s := []Clause(*u)
	return json.Marshal(s)
}

func (c *Clauses) UnmarshalJSON(data []byte) error {
	*c = Clauses{}

	r := dynamic.RawJSON(data)
	var cm []map[Type]dynamic.RawJSON
	if r.IsObject() {
		ce := map[Type]dynamic.RawJSON{}
		err := json.Unmarshal(r, &ce)
		if err != nil {
			return err
		}
		cm = []map[Type]dynamic.RawJSON{ce}
	} else {
		err := json.Unmarshal(r, &cm)
		if err != nil {
			return err
		}
	}

	for _, cd := range cm {
		for t, d := range cd {
			handler, ok := clauseHandlers[t]
			if !ok {
				return fmt.Errorf("%w <%s>", ErrUnsupportedType, t)
			}
			ce := handler()
			err := json.Unmarshal(d, &ce)
			if err != nil {
				return err
			}
			fmt.Println(ce)
			*c = append(*c, ce)
		}
	}
	return nil
}

func (c *Clauses) Validate() error {
	if c == nil {
		*c = Clauses{}
	}
	_, err := c.unpack()
	return err
}

func (c *Clauses) unpack() (*Clauses, error) {
	if c == nil {
		*c = Clauses{}
	}
	for i, ce := range *c {
		if v, ok := ce.(Clauser); ok {
			r, err := unpackClause(v)
			if err != nil {
				return nil, err
			}
			(*c)[i] = r
		}
	}
	return c, nil

}
func (c *Clauses) RemoveByName(name string) {
	if c == nil {
		*c = Clauses{}
	}
	for i, v := range *c {
		if v.Name() == name {
			*c = append((*c)[:i], (*c)[i+1:]...)
		}
	}
	return
}
func (c *Clauses) Add(clause Clause) error {

	var err error
	clause, err = unpackClause(clause)
	if err != nil {
		return err
	}
	if c == nil {
		*c = Clauses{clause}
		return nil
	}
	*c = append(*c, clause)
	return nil
}

// func (c Clause) MarshalJSON() ([]byte, error) {
// 	r, err := json.Marshal(c.Rule)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return json.Marshal(map[Type]dynamic.RawJSON{
// 		c.Rule.Type(): dynamic.RawJSON(r),
// 	})
// }
// func (c *Clause) UnmarshalJSON(data []byte) error {
// 	var m map[Type]dynamic.RawJSON
// 	err := json.Unmarshal(data, &m)
// 	if err != nil {
// 		return err
// 	}
// 	for t, d := range m {
// 		handler, ok := typeHandlers[t]
// 		if !ok {
// 			return ErrUnsupportedType
// 		}
// 		r := handler()
// 		err := json.Unmarshal(d, &r)
// 		if err != nil {
// 			return err
// 		}
// 		c.Rule = r
// 		c.Type = t
// 		return nil
// 	}
// 	return nil
// }

// func (c *Clause) UnmarshalBSON(data []byte) error {
// 	return c.UnmarshalJSON(data)
// }

// func (c Clause) MarshalBSON() ([]byte, error) {
// 	return c.MarshalJSON()
// }

func unpackClause(clause Clause) (Clause, error) {
	var err error
	if v, ok := clause.(Clauser); ok {
		clause, err = v.Clause()
		if err != nil {
			return nil, err
		}
	}
	return clause, nil
}
