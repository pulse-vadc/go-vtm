// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package vtm

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	TEXT_ONLY_OBJ = true
	STANDARD_OBJ  = false
)

type vtmObjectChild struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type vtmObjectChildren struct {
	Children []vtmObjectChild `json:"children"`
}

type vtmErrorResponse struct {
	ErrorId   string `json:"error_id"`
	ErrorText string `json:"error_text"`
	ErrorInfo string `json:"error_info"`
}

//********************************************

type vtmConnector struct {
	url           string
	client        *http.Client
	username      string
	password      string
	verifySslCert bool
	textOnly      bool
	contentType   string
	expectedCodes map[string][]int
	readOnly      bool
}

func (c vtmConnector) getChildConnector(path string) *vtmConnector {
	newUrl := c.url + path
	conn := newConnector(newUrl, c.username, c.password, c.verifySslCert)
	return conn
}

func (c vtmConnector) get() (io.ReadCloser, bool) {
	request, err := http.NewRequest("GET", c.url, nil)
	request.SetBasicAuth(c.username, c.password)
	response, err := c.client.Do(request)
	if err != nil {
		panic(err)
	}
	if response.StatusCode == 200 {
		return response.Body, true
	}
	return response.Body, false
}

func (c vtmConnector) put(body string, isTextObject bool) (io.ReadCloser, bool) {
	var contentType string
	if isTextObject == true {
		contentType = "application/octet-stream"
	} else {
		contentType = "application/json"
	}
	request, err := http.NewRequest("PUT", c.url, strings.NewReader(body))
	request.SetBasicAuth(c.username, c.password)
	request.Header.Set("Content-Type", contentType)
	response, err := c.client.Do(request)
	if err != nil {
		panic(err)
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return response.Body, true
	}
	return response.Body, false
}

func (c vtmConnector) delete() (io.ReadCloser, bool) {
	request, err := http.NewRequest("DELETE", c.url, nil)
	request.SetBasicAuth(c.username, c.password)
	response, err := c.client.Do(request)
	if err != nil {
		panic(err)
	}
	if response.StatusCode == 204 {
		return response.Body, true
	}
	return response.Body, false
}

func newConnector(url, username, password string, verifySslCert bool) *vtmConnector {
	conn := new(vtmConnector)
	conn.url = url
	conn.username = username
	conn.password = password
	conn.verifySslCert = verifySslCert
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifySslCert},
	}
	conn.client = &http.Client{Transport: tr, Timeout: 3 * time.Second}
	return conn
}

/*
VirtualTrafficManager is the central struct in the go-vtm library through which all tasks are performed.
*/
type VirtualTrafficManager struct {
	connector *vtmConnector
}

func (tm VirtualTrafficManager) testConnectivityOnce() (ok bool, err *url.Error) {
	ok = false
	defer func() {
		if r := recover(); r != nil {
			err = r.(*url.Error)
		}
	}()
	_, ok = tm.connector.get()
	return ok, nil
}

func (tm VirtualTrafficManager) testConnectivity() (bool, *url.Error) {
	var err *url.Error
	for i := 1; i <= 3; i++ {
		var ok bool
		ok, err = tm.testConnectivityOnce()
		if ok == true {
			return true, nil
		}
		time.Sleep(time.Duration(i) * time.Second)
	}
	return false, err
}

/*
NewVirtualTrafficManager creates an instance of VirtualTrafficManager and returns it, together with its reachability status.

Params:
	url		(string) The base URL of the target vTM, upto, but not including, the API verion portion.
				eg.	For direct connection to a vTM:
						https://my-vtm-1:9070/api
					For connection via a Services Director proxy:
						https://my-sd-1:8100/api/tmcm/<VER>/instances/<INSTANCE_ID>
	username	(string) Username to use for the connection.
				ie.	For direct connection to a vTM:
						Username on vTM with sufficient permissions to perform required operations.
					For connections via a Services Director proxy:
						Username on ServicesDirector with sufficient permissions to perform required operations.
	password	(string) Password to use for the connection.
				ie.	For direct connection to a vTM:
						vTM password for the user specified in the 'username' parameter.
					For connections via a Services Director proxy:
						Services Director password for the user specified in the 'username' parameter.

Returns:
	*VirtualTrafficManager		The newly-instantiated object
	bool						true if the target vTM is reachable with the provided parameters, else false
*/
func NewVirtualTrafficManager(url, username, password string, verifySslCert bool) (*VirtualTrafficManager, bool, *url.Error) {
	vtm := new(VirtualTrafficManager)
	conn := newConnector(url, username, password, verifySslCert)
	vtm.connector = conn
	contactable, contactErr := vtm.testConnectivity()
	return vtm, contactable, contactErr
}
