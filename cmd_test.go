package simpleexec

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCmd(t *testing.T) {
	require := require.New(t)

	testString := "this is a test string"

	cmd := ParseCmd(fmt.Sprintf("echo -n '%v'", testString))
	require.NotNil(cmd)
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf

	err := cmd.Start()
	require.Nil(err)

	err = cmd.Wait()
	require.Nil(err)

	require.Equal(testString, buf.String())
}

func TestBadCmd(t *testing.T) {
	require := require.New(t)

	cmd := ParseCmd(fmt.Sprintf("echoa 'Hello"))
	require.Nil(cmd)
}

func TestPipe(t *testing.T) {
	require := require.New(t)

	testString := "this is a test 'string'"
	cmd := ParseCmd(fmt.Sprintf(`echo -n "%v"`, testString))
	require.NotNil(cmd)

	pipedCmd := cmd.Pipe("sed -e 's/this is //g'")
	require.NotNil(pipedCmd)

	buf := bytes.NewBuffer(nil)
	pipedCmd.Stdout = buf

	err := pipedCmd.Start()
	require.Nil(err)

	err = pipedCmd.Wait()
	require.Nil(err)

	require.Equal("a test 'string'", buf.String())
}

func TestComplexPipe(t *testing.T) {
	require := require.New(t)

	testString := "this is a test 'string'"
	cmd := ParseCmd(fmt.Sprintf(`echo -n "%v"`, testString))
	require.NotNil(cmd)

	pipedCmd := cmd.Pipe("sed -e 's/this is //g'").
		Pipe(`sed -e "s/ 'string'//g"`).
		Pipe(`sed -e "s/a /ab /g"`).
		Pipe("wc -c")
	require.NotNil(pipedCmd)

	buf := bytes.NewBuffer(nil)
	pipedCmd.Stdout = buf

	err := pipedCmd.Start()
	require.Nil(err)

	err = pipedCmd.Wait()
	require.Nil(err)

	require.Equal("7\n", buf.String())
}
