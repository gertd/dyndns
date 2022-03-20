//nolint:dupl
package namecheap

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSuccessResponse(t *testing.T) {
	var i interfaceResponse

	buf := []byte(success)
	err := xml.Unmarshal(buf, &i)

	assert.Equal(t, nil, err)
	assert.Equal(t, "SETDNSHOST", i.Command)
	assert.Equal(t, "eng", i.Language)
	assert.Equal(t, true, i.Done)
	assert.Equal(t, 0, i.ErrCount)
	assert.Equal(t, 0, len(i.Errors))
	assert.Equal(t, 0, i.ResponseCount)
	assert.Equal(t, 0, len(i.Responses))
	assert.Equal(t, "52.32.92.192", i.IP)

	t.Logf("%v\n", i)
}

func TestFailureResponse(t *testing.T) {
	var i interfaceResponse

	buf := []byte(failure)
	err := xml.Unmarshal(buf, &i)

	assert.Equal(t, nil, err)
	assert.Equal(t, "SETDNSHOST", i.Command)
	assert.Equal(t, "eng", i.Language)
	assert.Equal(t, true, i.Done)
	assert.Equal(t, 1, i.ErrCount)
	assert.Equal(t, 1, len(i.Errors))
	assert.Equal(t, 1, i.ResponseCount)
	assert.Equal(t, 1, len(i.Responses))
	assert.Equal(t, "", i.IP)

	t.Logf("%v\n", i)

}

var success = `
<?xml version="1.0"?>
<interface-response>
    <Command>SETDNSHOST</Command>
    <Language>eng</Language>
    <IP>52.32.92.192</IP>
    <ErrCount>0</ErrCount>
    <ResponseCount>0</ResponseCount>
    <Done>true</Done>
    <debug>
        <![CDATA[]]>
    </debug>
</interface-response>`

var failure = `
<?xml version="1.0"?>
<interface-response>
    <Command>SETDNSHOST</Command>
    <Language>eng</Language>
    <ErrCount>1</ErrCount>
    <errors>
        <Err1>No Records updated. A record not Found;</Err1>
    </errors>
    <ResponseCount>1</ResponseCount>
    <responses>
        <response>
            <ResponseNumber>380091</ResponseNumber>
            <ResponseString>No updates; A record not Found;</ResponseString>
        </response>
    </responses>
    <Done>true</Done>
    <debug>
        <![CDATA[]]>
    </debug>
</interface-response>`
