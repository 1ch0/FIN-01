//vela 里面的 cuelang 模板

output: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	spec: {
		selector: matchLabels: {
			"app.oam.dev/component": context.name
			"app":                   context.name
		}

		template: {
			metadata: labels: {
				"app.oam.dev/component": context.name
				"app":                   context.name
			}

			spec: {
				containers: [{
					name:  context.name
					image: parameter.image

					if parameter["cmd"] != _|_ {
						command: parameter.cmd
					}

					if parameter["env"] != _|_ {
						env: parameter.env
					}

					if context["config"] != _|_ {
						env: context.config
					}

					ports: [{
						containerPort: parameter.port
					}]

					if parameter["cpu"] != _|_ {
						resources: {
							limits:
								cpu: parameter.cpu
							memory: parameter.memory
							requests:
								cpu: parameter.cpu
							memory: parameter.memory
						}
					}

				}]
		}
		}
	}
}

parameter: {
	// +usage=Which image would you like to use for your service
	// +short=i
	image: string

	// +usage=Commands to run in the container
	cmd?: [...string]

	// +usage=Which port do you want customer traffic sent to
	// +short=p
	port: *80 | int
	// +usage=Define arguments by using environment variables
	env?: [...{
		// +usage=Environment variable name
		name: string
		// +usage=The value of the environment variable
		value?: string
		// +usage=Specifies a source the value of this var should come from
		valueFrom?: {
			// +usage=Selects a key of a secret in the pod's namespace
			secretKeyRef: {
				// +usage=The name of the secret in the pod's namespace to select from
				name: string
				// +usage=The key of the secret to select from. Must be a valid secret key
				key: string
			}
		}
	}]
	// +usage=Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	cpu?: string
	// +usage=Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	memory?: string
}

parameter: {
	image: "test"
	cmd: ["nginx"]
}

context: {
    name: "test"
}