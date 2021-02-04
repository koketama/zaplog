package zaplog

import (
	"testing"

	"github.com/koketama/errors"
)

func TestJSONLogger(t *testing.T) {
	logger, err := NewJSONLogger(WithField("name", "minami"), WithFileP("/tmp/zap/dummy/zaplog.log"))
	if err != nil {
		t.Fatal(err)
	}
	defer logger.Sync()

	err = errors.New("pkg error")
	logger.Error("err occurs", WrapMeta(nil, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
	logger.Error("err occurs", WrapMeta(err, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)

}
