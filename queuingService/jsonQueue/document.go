package jsonQueue

type Docs []interface{}

func NewJob() *Docs {
	return &Docs{}
}

//AppendNewJob ... append new message in the Docs array.
func (j *Docs) AppendNewJob(doc interface{}) {
	*j = append(*j, doc)
}

//Pop ... delete the consumed message from the Docs.
func (j *Docs) Pop() interface{} {
	if len(*j) > 0 {
		doc := (*j)[0]
		*j = (*j)[1:]
		return doc
	}
	return nil
}
