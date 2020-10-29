package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Metric struct {
	MatricName                                       string    `json:"__name__"`
	ContainerLabelAnnotationKubernetesIoConfigSeen   time.Time `json:"container_label_annotation_kubernetes_io_config_seen"`
	ContainerLabelAnnotationKubernetesIoConfigSource string    `json:"container_label_annotation_kubernetes_io_config_source"`
	ContainerLabelAppID                              string    `json:"container_label_app_id"`
	ContainerLabelIoKubernetesContainerName          string    `json:"container_label_io_kubernetes_container_name"`
	ContainerLabelIoKubernetesDockerType             string    `json:"container_label_io_kubernetes_docker_type"`
	ContainerLabelIoKubernetesPodName                string    `json:"container_label_io_kubernetes_pod_name"`
	ContainerLabelIoKubernetesPodNamespace           string    `json:"container_label_io_kubernetes_pod_namespace"`
	ContainerLabelIoKubernetesPodUID                 string    `json:"container_label_io_kubernetes_pod_uid"`
	ContainerLabelNodeID                             string    `json:"container_label_node_id"`
	ContainerLabelPodTemplateHash                    string    `json:"container_label_pod_template_hash"`
	ControllerRevisionHash                           string    `json:"controller_revision_hash"`
	ExportedName                                     string    `json:"exported_name"`
	ID                                               string    `json:"id"`
	Image                                            string    `json:"image"`
	Instance                                         string    `json:"instance"`
	Interface                                        string    `json:"interface"`
	Job                                              string    `json:"job"`
	KubernetesPodName                                string    `json:"kubernetes_pod_name"`
	Name                                             string    `json:"name"`
	PodTemplateGeneration                            string    `json:"pod_template_generation"`
}

type Data struct {
	ResultType string `json:"resultType"`
	Result     []struct {
		Metric Metric        `json:"metric"`
		Value  []interface{} `json:"value"`
	} `json:"result"`
}

type ContainerData struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

func main() {

	var structdata ContainerData
	response, err := http.Get("http://10.43.12.194:30000/api/v1/query?query=container_network_transmit_bytes_total%7Bcontainer_label_io_kubernetes_pod_namespace%3D%22default%22%2Ccontainer_label_io_kubernetes_pod_name%3D%22applblqj-9dd6c6465-lm7sh%22%7D")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := string(responseData)
	Data := []byte(jsonData)
	err = json.Unmarshal(Data, &structdata)

	if err != nil {
		fmt.Println(err)
	}

	//	fmt.Println(structdata.Data.Result[0].Value[1])

	//	matricValue := structdata.Data.Result[0].Value[1]

	str, _ := structdata.Data.Result[0].Value[1].(string)

	fmt.Println(str)

	//	fmt.Printf("%T", structdata.Data.Result[0].Value[1])

	//	matricValue, _ := strconv.Atoi(str)

	matricValue, _ := strconv.ParseInt(str, 10, 64)

	fmt.Println(matricValue)

	fmt.Printf("%T", matricValue)

}
