package main

func LoginDockerHub(stage string, user string, password string) {
	var login_cmd string
	login_cmd = "docker login cr-" + stage + ".pentium.network -u " + user + " -p " + password
	RunCommand(login_cmd)
}
