// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package v1beta2

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// KafkaConnector
type KafkaConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The specification of the Kafka Connector.
	Spec *KafkaConnectorSpec `json:"spec,omitempty"`

	// The status of the Kafka Connector.
	Status *KafkaConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// KafkaConnectorList contains a list of instances.
type KafkaConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// A list of Kafka objects.
	Items []KafkaConnector `json:"items,omitempty"`
}

func init() {
	SchemeBuilder.Register(&KafkaConnector{}, &KafkaConnectorList{})
}

// The specification of the Kafka Connector.
type KafkaConnectorSpec struct {
	// The Class for the Kafka Connector.
	Class *string `json:"class,omitempty"`

	// The Kafka Connector configuration. The following properties cannot be set:
	// connector.class, tasks.max.
	Config *apiextensions.JSON `json:"config,omitempty"`

	// Whether the connector should be paused. Defaults to false.
	Pause *bool `json:"pause,omitempty"`

	// The maximum number of tasks for the Kafka Connector.
	TasksMax *int32 `json:"tasksMax,omitempty"`
}

// The Kafka Connector configuration. The following properties cannot be set:
// connector.class, tasks.max.
//type KafkaConnectorSpecConfig map[string]interface{}

// The status of the Kafka Connector.
type KafkaConnectorStatus struct {
	// List of status conditions.
	Conditions []KafkaConnectorStatusConditionsElem `json:"conditions,omitempty"`

	// The connector status, as reported by the Kafka Connect REST API.
	ConnectorStatus *apiextensions.JSON `json:"connectorStatus,omitempty"`

	// The generation of the CRD that was last reconciled by the operator.
	ObservedGeneration *int32 `json:"observedGeneration,omitempty"`

	// The maximum number of tasks for the Kafka Connector.
	TasksMax *int32 `json:"tasksMax,omitempty"`

	// The list of topics used by the Kafka Connector.
	Topics []string `json:"topics,omitempty"`
}

type KafkaConnectorStatusConditionsElem struct {
	// Last time the condition of a type changed from one status to another. The
	// required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about the condition's last
	// transition.
	Message *string `json:"message,omitempty"`

	// The reason for the condition's last transition (a single word in CamelCase).
	Reason *string `json:"reason,omitempty"`

	// The status of the condition, either True, False or Unknown.
	Status *string `json:"status,omitempty"`

	// The unique identifier of a condition, used to distinguish between other
	// conditions in the resource.
	Type *string `json:"type,omitempty"`
}

// The connector status, as reported by the Kafka Connect REST API.
//type KafkaConnectorStatusConnectorStatus map[string]interface{}
