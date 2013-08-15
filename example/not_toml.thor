# This is a thor document.  Boom.
# Based on the example at https://github.com/mojombo/toml

title: "thor example"

owner: {
	name: "Me Myself"
	organization: "MeCo"
	bio: "Some dude."
	dob: "2999-12-31T23:59:59" # First class dates?  Nah.  Why not?  Idunno.
}

database: {
	server: "192.168.1.1"
	ports: { 8001 8001 8002 }
	connection_max: 5000
	enabled: true
}

servers: {
	# You can indent as you please. Tabs or spaces. thor don't care.
	alpha: {
		ip: "10.0.0.1"
		dc: "eqdc10"
	}

	beta: {
		ip: "10.0.0.2"
		dc: "eqdc10"
	}

	clients: {
		data: { { "gamma" "delta" } { 1 2 } }
	}

	# Line breaks are OK when inside lists
	hosts: {
		"alpha"
		"omega"
	}
}
