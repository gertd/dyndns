package namecheap_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/gertd/dyndns/pkg/namecheap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSuccessResponse(t *testing.T) {
	var i namecheap.InterfaceResponse

	buf := []byte(success)
	err := xml.Unmarshal(buf, &i)

	require.NoError(t, err)
	assert.Equal(t, "SETDNSHOST", i.Command)
	assert.Equal(t, "eng", i.Language)
	assert.True(t, i.Done)
	assert.Equal(t, 0, i.ErrCount)
	assert.Empty(t, i.Errors)
	assert.Equal(t, 0, i.ResponseCount)
	assert.Empty(t, i.Responses)
	assert.Equal(t, "52.32.92.192", i.IP)

	t.Logf("%v\n", i)
}

func TestFailureResponse(t *testing.T) {
	var i namecheap.InterfaceResponse

	buf := []byte(failure)
	err := xml.Unmarshal(buf, &i)

	require.NoError(t, err)
	assert.Equal(t, "SETDNSHOST", i.Command)
	assert.Equal(t, "eng", i.Language)
	assert.True(t, i.Done)
	assert.Equal(t, 1, i.ErrCount)
	assert.Len(t, i.Errors, 1)
	assert.Equal(t, 1, i.ResponseCount)
	assert.Len(t, i.Responses, 1)
	assert.Empty(t, i.IP)

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
