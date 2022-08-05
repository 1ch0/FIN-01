
app: {
	name: string
	tech: string
	mem:  int

	if tech == "react" {
		tier: "frontend"
	}
	if tech != "react" {
		tier: "backend"
	}

	if mem < 1Gi {
		footprint: "small"
	}
	if mem >= 1Gi && mem < 4Gi {
		footprint: "medium"
	}
	if mem >= 4Gi {
		footprint: "large"
	}
}

app: {
	name: "web1"
	tech: "react"
	mem: 2G
}