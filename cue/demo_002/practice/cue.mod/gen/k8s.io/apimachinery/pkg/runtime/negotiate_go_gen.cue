// Code generated by cue get go. DO NOT EDIT.

//cue:generate_config cue get go k8s.io/apimachinery/pkg/runtime

package runtime

// NegotiateError is returned when a ClientNegotiator is unable to locate
// a serializer for the requested operation.
#NegotiateError: {
	ContentType: string
	Stream:      bool
}