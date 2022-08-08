// Code generated by cue get go. DO NOT EDIT.

//cue:generate_config cue get go k8s.io/apimachinery/pkg/apis/meta/v1

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

// Event represents a single event to a watched resource.
//
// +protobuf=true
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
#WatchEvent: {
	type: string @go(Type) @protobuf(1,bytes,opt)

	// Object is:
	//  * If Type is Added or Modified: the new state of the object.
	//  * If Type is Deleted: the state of the object immediately before deletion.
	//  * If Type is Error: *Status is recommended; other types may make sense
	//    depending on context.
	object: runtime.#RawExtension @go(Object) @protobuf(2,bytes,opt)
}

// InternalEvent makes watch.Event versioned
// +protobuf=false
#InternalEvent: watch.#Event
