package inject

// #tag() is how we inject data
env:      *"dev" | string @tag(env)
database: string          @tag(database)

#DB: {
	host: #hosts[env]
	port: string | *"5432"
	db:   database

	conn: "postgres://\(host):\(port)/\(db)"
}

#hosts: [string]: string

#hosts: {
	dev: "postgres.dev"
	stg: "postgres.stg"
	prd: "postgres.prd"
}

// # -t key=value  -e to eval a specific value
// cue eval tags.cue -t database="foo" -e "#DB.conn"
// cue eval tags.cue -t database="foo" -t env="prd" -e "#DB.conn"
//"postgres://postgres.dev:5432/foo"
