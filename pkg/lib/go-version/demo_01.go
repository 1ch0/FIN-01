package main

import (
	"fmt"
	"sort"

	"github.com/hashicorp/go-version"
)

func main() {

	v1, _ := version.NewVersion("1.2")
	v2, _ := version.NewVersion("1.5+metadata")

	// Comparison example. There is also GreaterThan, Equal, and just
	// a simple Compare that returns an int allowing easy >=, <=, etc.
	if v1.LessThan(v2) {
		fmt.Printf("%s is less than %s \n", v1, v2)
	}

	// Constraints example.
	constraints, _ := version.NewConstraint(">= 1.0, < 1.4")
	if constraints.Check(v1) {
		fmt.Printf("%s satisfies constraints %s \n", v1, constraints)
	}

	versionsRaw := []string{"1.1", "0.7.1", "1.4-beta", "1.4", "2"}
	versions := make([]*version.Version, len(versionsRaw))
	for i, raw := range versionsRaw {
		v, _ := version.NewVersion(raw)
		versions[i] = v
	}

	// After this, the versions are properly sorted
	sort.Sort(version.Collection(versions))
	fmt.Printf("Versions: %v\n", versions)
}
