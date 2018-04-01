package main_test

import (
	"github.com/tsongpon/listener/data"
	"github.com/tsongpon/listener/route"
	"gopkg.in/resty.v1"
	"net/http/httptest"
	"os"
	"testing"

	"encoding/json"
	"github.com/fsouza/go-dockerclient"
	"github.com/tsongpon/listener/transport"
	"log"
	"time"
)

const integrationMongoContainerName = "rp_integrationTest_mongo"

var (
	server           *httptest.Server
	baseUrl          string
	mongoContainerId string
	dockerClient     *docker.Client
)

func TestMain(m *testing.M) {
	setupTest()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func TestPing(t *testing.T) {
	resp, err := resty.R().Get(baseUrl + "/ping")

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if resp.StatusCode() != 200 {
		t.Errorf("Success expected: %d", resp.StatusCode()) //Uh-oh this means our test failed
	}
}

func TestFacebookHookGetWithValidToken(t *testing.T) {
	os.Setenv("TOKEN", "validToken")
	resp, err := resty.R().Get(baseUrl + "/facebookhook?hub.verify_token=validToken&hub.challenge=123")
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if resp.StatusCode() != 200 {
		t.Errorf("Success expected: %d", resp.StatusCode())
	}

	if resp.String() != "123" {
		t.Errorf("hub.challenge not valid")
	}
}

func TestFacebookHookGetWithInValidToken(t *testing.T) {
	os.Setenv("TOKEN", "validToken")
	resp, err := resty.R().Get(baseUrl + "/facebookhook?hub.verify_token=faketoken&hub.challenge=123")

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if resp.StatusCode() != 400 {
		t.Errorf("Success expected: %d", resp.StatusCode())
	}
}

func TestGetActivities(t *testing.T) {
	resp, err := resty.R().Get(baseUrl + "/useractivities")

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if resp.StatusCode() != 200 {
		t.Errorf("Success expected: %d", resp.StatusCode())
	}

	responseTransport := new(transport.UserActivities)
	json.Unmarshal([]byte(resp.String()), &responseTransport)
}

func TestGetActivitiesWithNotExistUser(t *testing.T) {
	notExistUserId := time.Now().Nanosecond()
	resp, err := resty.R().Get(baseUrl + "/useractivities?" + string(notExistUserId))

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode() != 200 {
		t.Errorf("Success expected: %d", resp.StatusCode())
	}

	responseTransport := new(transport.UserActivities)
	json.Unmarshal([]byte(resp.String()), &responseTransport)

	if responseTransport.Total != 0 || responseTransport.Size != 0 {
		t.Error("Not exist useractivities should return empty response")
	}
}

func TestPostUserUpdate(t *testing.T) {
	jsonPayload := `{
					"entry": [
						{
						  "time": 1521910231,
						  "changes": [
							{
							  "field": "pic_square_with_logo",
							  "value": "http://somepic/io"
							}
						  ],
						  "id": "TestPostUserUpdate",
						  "uid": "TestPostUserUpdate"
						}
					  ],
					  "object": "user"
 					}`

	res, _ := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonPayload).
		Post(baseUrl + "/facebookhook")

	if res.StatusCode() != 200 {
		t.Error("Post webhook with valid payload should return status 200")
	}
}

func TestPostUserUpdateAndQuery(t *testing.T) {
	jsonPayload := `{
					"entry": [
						{
						  "time": 1521910231,
						  "changes": [
							{
							  "field": "pic_square_with_logo",
							  "value": "http://somepic/io"
							}
						  ],
						  "id": "TestPostUserUpdateAndQuery",
						  "uid": "TestPostUserUpdateAndQuery"
						}
					  ],
					  "object": "user"
 					}`

	resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonPayload).
		Post(baseUrl + "/facebookhook")

	getResponse, getErr := resty.R().Get(baseUrl + "/useractivities?userid=TestPostUserUpdateAndQuery")

	if getErr != nil {
		panic(getErr)
	}

	responseTransport := new(transport.UserActivities)
	json.Unmarshal([]byte(getResponse.String()), &responseTransport)

	if responseTransport.Total != 1 {
		t.Error("Response should contain 1 item")
	}

	if responseTransport.Data[0].UserId != "TestPostUserUpdateAndQuery" {
		t.Error("Response has wrong userId")
	}

	if responseTransport.Data[0].Field != "pic_square_with_logo" {
		t.Error("Response has wrong field value")
	}

	if responseTransport.Data[0].Value != "http://somepic/io" {
		t.Error("Response has wrong value")
	}
}

func TestPostUserUpdateAndQueryByUser(t *testing.T) {
	jsonPayloadUser1 := `{
					"entry": [
						{
						  "time": 1521910231,
						  "changes": [
							{
							  "field": "pic_square_with_logo",
							  "value": "http://somepic/io"
							}
						  ],
						  "id": "TestPostUserUpdateAndQueryWithParam_1",
						  "uid": "TestPostUserUpdateAndQueryWithParam_1"
						}
					  ],
					  "object": "user"
 					}`
	resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonPayloadUser1).
		Post(baseUrl + "/facebookhook")

	jsonPayloadUser2 := `{
					"entry": [
						{
						  "time": 1521910231,
						  "changes": [
							{
							  "field": "pic_square_with_logo",
							  "value": "http://user2/io"
							}
						  ],
						  "id": "TestPostUserUpdateAndQueryWithParam_2",
						  "uid": "TestPostUserUpdateAndQueryWithParam_2"
						}
					  ],
					  "object": "user"
 					}`
	resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonPayloadUser2).
		Post(baseUrl + "/facebookhook")

	getResponse, getErr := resty.R().Get(baseUrl + "/useractivities?userid=TestPostUserUpdateAndQueryWithParam_2")

	if getErr != nil {
		panic(getErr)
	}

	responseTransport := new(transport.UserActivities)
	json.Unmarshal([]byte(getResponse.String()), &responseTransport)

	if responseTransport.Size != 1 {
		t.Error("Response should contain 1 item")
	}

	if responseTransport.Data[0].UserId != "TestPostUserUpdateAndQueryWithParam_2" {
		t.Error("Response contain wrong userId")
	}

	if responseTransport.Data[0].Value != "http://user2/io" {
		t.Error("Response contain wrong value")
	}
}

func TestPostUserUpdateAndQueryWithParam(t *testing.T) {
	jsonPayloadUser1 := `{
					"entry": [
						{
						  "time": 1521910231,
						  "changes": [
							{
							  "field": "pic_square_with_logo",
							  "value": "http://somepic/io"
							}
						  ],
						  "id": "TestPostUserUpdateAndQueryWithParam",
						  "uid": "TestPostUserUpdateAndQueryWithParam"
						}
					  ],
					  "object": "user"
 					}`
	resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonPayloadUser1).
		Post(baseUrl + "/facebookhook")

	jsonPayloadUser2 := `{
					"entry": [
						{
						  "time": 1521910231,
						  "changes": [
							{
							  "field": "pic_square_with_logo",
							  "value": "http://somepic/io"
							}
						  ],
						  "id": "TestPostUserUpdateAndQueryWithParam",
						  "uid": "TestPostUserUpdateAndQueryWithParam"
						}
					  ],
					  "object": "user"
 					}`
	resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonPayloadUser2).
		Post(baseUrl + "/facebookhook")

	getResponse, getErr := resty.R().Get(baseUrl + "/useractivities?size=1")

	if getErr != nil {
		panic(getErr)
	}

	responseTransport := new(transport.UserActivities)
	json.Unmarshal([]byte(getResponse.String()), &responseTransport)

	if responseTransport.Size != 1 {
		t.Error("Response should contain 1 item")
	}

	if len(responseTransport.Data) != 1 {
		t.Error("Response data should contain 1 item")
	}
}

//Util function
func setupTest() {
	prepareDocker()
	prepareEnv()
	startContainer()
	dbHost := "localhost:28017"
	dbName := "RP_integrationTest"
	data.InitDB(dbHost, dbName)
	redPlanetRoute := route.NewRedPlanetRouter()
	server = httptest.NewServer(redPlanetRoute) //Creating new server with the user handlers

	baseUrl = server.URL //Grab the address for the API endpoint
}

func prepareDocker() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	dockerClient = client
}

func prepareEnv() {
	log.Println("prepareing test environment")
	mongoContainerIdFromHost := getMongoContainerId()
	if mongoContainerIdFromHost == "" {
		log.Println("pulling mongo docker image(if missing)")
		dockerClient.PullImage(docker.PullImageOptions{Repository: "mongo:3.6"}, docker.AuthConfiguration{})
		portBindings := map[docker.Port][]docker.PortBinding{
			"27017/tcp": {{HostPort: "28017"}}}

		createContHostConfig := docker.HostConfig{
			PortBindings:    portBindings,
			PublishAllPorts: true,
			Privileged:      false}

		container, err := dockerClient.CreateContainer(docker.CreateContainerOptions{
			Name: integrationMongoContainerName,
			Config: &docker.Config{
				Image: "mongo:3.6",
			},
			HostConfig: &createContHostConfig,
		})

		if err == nil {
			mongoContainerId = container.ID
		}
	} else {
		mongoContainerId = mongoContainerIdFromHost
	}
}

func getMongoContainerId() string {
	cons, _ := dockerClient.ListContainers(docker.ListContainersOptions{All: true})
	for _, con := range cons {
		if con.Names[0] == "/"+integrationMongoContainerName {
			return con.ID
		}
	}
	return ""
}

func startContainer() {
	portBindings := map[docker.Port][]docker.PortBinding{
		"27017/tcp": {{HostPort: "28017"}}}

	dockerClient.StartContainer(mongoContainerId, &docker.HostConfig{PortBindings: portBindings})
}

func tearDown() {
	dockerClient.StopContainer(mongoContainerId, 2)
	dockerClient.RemoveContainer(docker.RemoveContainerOptions{ID: mongoContainerId, Force: true})
}
