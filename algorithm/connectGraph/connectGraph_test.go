package connectGraph

import (
	"testing"
	"io/ioutil"
	"bufio"
	"bytes"
	"strconv"
	"regexp"
	"github.com/myteksi/go/commons/util/log/logging"
	"fmt"
	"strings"
	"reflect"
)

const connectGraphTestStr = "connectGraphTest"
func TestTestCases(t *testing.T)  {
	testConnectGraph(t, "testcases.in", "")
}

func testConnectGraph(t *testing.T, actionFile string, expectedRes string)  {
	actionData, err := ioutil.ReadFile(actionFile)
	if err != nil {
		t.Error(err, " when reading file '", actionFile, "'")
	}

	actionDataReader := bufio.NewReader(bytes.NewBuffer(actionData))
	nodeCountStr, err := actionDataReader.ReadString('\n')
	if err != nil {
		t.Error(err, " when reading the first line")
	}
	nodeCountStr = strings.TrimRight(nodeCountStr, "\n")
	nodeCount, err := strconv.Atoi(nodeCountStr)
	if err != nil {
		t.Error(err, " when converting ", nodeCountStr)
	}
	fmt.Println("nodeCount:", nodeCount)
	g := NewGraph(nodeCount)
	v := reflect.ValueOf(g)
	re := regexp.MustCompile("(\\w+?)\\((\\w+?)\\s*?,\\s*?(\\w+?)\\)")
	for {
		actionStr, err := actionDataReader.ReadString('\n')
		if err != nil {
			logging.Debug(connectGraphTestStr, err.Error(), " happens when reading line")
			break
		}
		segs := re.FindStringSubmatch(actionStr)
		arg0, _ := strconv.Atoi(segs[2])
		arg1, _ := strconv.Atoi(segs[3])
		funName := string(bytes.ToUpper([]byte(segs[1][:1]))) + segs[1][1:]
		funcVariable := v.MethodByName(funName)
		var expectedRes string
		for _, v := range funcVariable.Call([]reflect.Value{reflect.ValueOf(arg0), reflect.ValueOf(arg1)}) {
			expectedRes += fmt.Sprint(v)
		}
		if len(expectedRes) > 0 {
			fmt.Println("Expected Result:", expectedRes)
		}
	}

}
