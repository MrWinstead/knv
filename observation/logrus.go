package observation

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

var (
	_ Observer = &logrusObserver{}
)

type logrusObserver struct {
	logger *logrus.Logger
}

func NewLogrusObserver(writer io.Writer) Observer {
	created := &logrusObserver{
		logger: logrus.New(),
	}
	created.logger.SetOutput(writer)
	return created
}

func (o *logrusObserver) Submit(schema interface{}, value interface{}) {
	o.logger.WithFields(logrus.Fields{
		"schema": fmt.Sprintf("%+v", schema),
		"value":  fmt.Sprintf("%+v", value),
	}).Println()
}
