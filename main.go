package gojebug

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
)

var (
	should = assert.New(TLogger{})
)

func CheckErr(err error) {
	should.NoError(err)
}

func print(in string) string {
	fmt.Println(in)
	return in
}

type TLogger struct{}

func (t TLogger) Errorf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func prettyJsonPrint(something interface{}) string {
	j, err := json.MarshalIndent(something, "", "\t")
	CheckErr(err)
	return string(j)
}

func PrettyJsonPrint(something interface{}) string {
	return print(prettyJsonPrint(something))
}

func JsonPrint(something interface{}) string {
	j, err := json.Marshal(something)
	CheckErr(err)
	res := string(j)
	fmt.Println(res)
	return res
}

func Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return should.Equal(expected, actual, msgAndArgs...)
}

func printReaderContent(reader io.Reader) string {
	b, err := ioutil.ReadAll(reader)
	CheckErr(err)
	return string(b)
}

func PrintReaderContent(reader io.Reader) string {
	return print(printReaderContent(reader))
}
func PrintReaderContentJSON(reader io.Reader) string {
	var res map[string]interface{}
	err := json.Unmarshal([]byte(printReaderContent(reader)), &res)
	CheckErr(err)
	return PrettyJsonPrint(res)
}

func PrintRequest(r http.Request) string {
	return print(printRequest(r))
}

func printRequest(r http.Request) string {
	var res string
	res += "METHOD = " + r.Method + "\n"
	res += fmt.Sprintf("======%s================================\n", "URL")
	res += r.URL.String() + "\n"
	res += fmt.Sprintf("======%s================================\n", "QUERY PARAMS")
	res += prettyJsonPrint(r.URL.Query()) + "\n"
	res += fmt.Sprintf("======%s================================\n", "BODY")
	res += printReaderContent(r.Body) + "\n"
	return res
}
