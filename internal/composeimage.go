package gdeyamloperator

func ComposeImageName(mode string, hubdomain string, stage string, module string, tag string) string {

	var complete_image string
	if mode == "fqdn" {
		complete_image = hubdomain + "/" + stage + "/" + module + ":" + tag
	} else if mode == "nexus" {
		complete_image = hubdomain + "/" + module + ":" + tag
	}
	return complete_image

}
