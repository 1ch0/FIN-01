#labels: [string]: string

#metadata: {
	name:   string
	labels: #labels
	annotations?: [string]: string
}

#requiredLabels: #labels & {
	app:  string
	env:  string
	team: string
}

#defaultLabels: #requiredLabels & {
	env: *"dev" | "stg" | "prd"
}

#myLabels: #defaultLabels & {
	app:  "cuetorials"
	team: "hifstadter"
}

#Schema: #Deployment | #Service | #Ingress

#Schema: {
	metadata: #metadata & {
		labels: #myLabels
	}
}

#Deployment: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	...
}

#Service: {
	apiVersion: "v1"
	kind:       "Service"
	...
}

#Ingress: {
	apiVersion: "extensions/v1beta1"
	kind:       "Ingress"
	...
}
