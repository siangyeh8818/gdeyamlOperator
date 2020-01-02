package utility

import (
	"fmt"
	"strings"
)

func ImagenameSplit(rawimage string) (string, string) {
	var parsername []string
	var tag []string

	var return_image string

	parsername = strings.Split(rawimage, "/")
	if len(parsername) == 2 {
		for i := range parsername {
			//	fmt.Println(parsername[i])
			if i == 1 {
				tag = strings.Split(parsername[i], ":")
				/*      for j := range tag {
				                fmt.Println(tag[j])
							              }
				*/
			}
		}
		return_image = tag[0]
	} else if len(parsername) == 3 {
		for i := range parsername {
			//	fmt.Println(parsername[i])
			if i == 2 {
				tag = strings.Split(parsername[i], ":")
			}
		}
		return_image = parsername[1] + "/" + tag[0]

	}
	return parsername[0], return_image
}

func ImagenameSplitReturnTag(rawimage string) (string, string, string) {
	var parsername []string
	var tag []string
	var tagstr string

	var return_image string
	var return_image_tag string
	var return_image_stage string

	parsername = strings.Split(rawimage, "/")
	if len(parsername) == 1 {
		for i := range parsername {
			//fmt.Println(parsername[i])
			if i == 0 {
				if GetSubstringColonIndex(parsername[i]) != -1 {
					tagstr = GetSubstring(parsername[i], GetSubstringColonIndex(parsername[i]), len(parsername[i]))
				} else if GetSubstringColonIndex(parsername[i]) == -1 {
					fmt.Println("%s has't witten tag  ", rawimage)
					tagstr = "latest"
				}

				tag = strings.Split(parsername[i], ":")
				/*
					if len(tag) == 1 {
						fmt.Println("%s has't witten tag  ", rawimage)
						tag = append(tag, "latest")
					}
				*/
			}
		}
		return_image_stage = ""
		return_image = tag[0]
		return_image_tag = tagstr
	} else if len(parsername) == 2 {
		for i := range parsername {
			//fmt.Println(parsername[i])
			if i == 1 {
				if GetSubstringColonIndex(parsername[i]) != -1 {
					tagstr = GetSubstring(parsername[i], GetSubstringColonIndex(parsername[i]), len(parsername[i]))
				} else if GetSubstringColonIndex(parsername[i]) == -1 {
					fmt.Println("%s has't witten tag  ", rawimage)
					tagstr = "latest"
				}

				tag = strings.Split(parsername[i], ":")
				/*
					if len(tag) == 1 {
						fmt.Println("%s has't witten tag  ", rawimage)
						tag = append(tag, "latest")
					}
				*/

			}
		}

		return_image = tag[0]
		return_image_tag = tagstr
		return_image_stage = GetStage(parsername[0])
	} else if len(parsername) == 3 {
		for i := range parsername {
			//fmt.Println(parsername[i])
			if i == 2 {
				if GetSubstringColonIndex(parsername[i]) != -1 {
					tagstr = GetSubstring(parsername[i], GetSubstringColonIndex(parsername[i]), len(parsername[i]))
				} else if GetSubstringColonIndex(parsername[i]) == -1 {
					fmt.Println("%s has't witten tag  ", rawimage)
					tagstr = "latest"
				}

				tag = strings.Split(parsername[i], ":")
				/*
					if len(tag) == 1 {
						fmt.Println("%s has't witten tag  ", rawimage)
						tag = append(tag, "latest")
					}
				*/
			}
		}
		return_image = tag[0]
		return_image_tag = tagstr
		return_image_stage = GetStage(parsername[1])
	} else if len(parsername) == 4 {
		for i := range parsername {
			//fmt.Println(parsername[i])
			if i == 3 {
				if GetSubstringColonIndex(parsername[i]) != -1 {
					tagstr = GetSubstring(parsername[i], GetSubstringColonIndex(parsername[i]), len(parsername[i]))
				} else if GetSubstringColonIndex(parsername[i]) == -1 {
					fmt.Println("%s has't witten tag  ", rawimage)
					tagstr = "latest"
				}

				tag = strings.Split(parsername[i], ":")
				/*
					if len(tag) == 1 {
						fmt.Println("%s has't witten tag  ", rawimage)
						tag = append(tag, "latest")
					}
				*/
			}
		}
		return_image = tag[0]
		return_image_tag = tagstr
		return_image_stage = GetStage(parsername[1])
	}

	return return_image_stage, return_image, return_image_tag
}

func GetStage(name string) string {
	var f_stage string
	if strings.Contains(name, "dev") {
		f_stage = "dev"
	} else if strings.Contains(name, "qa") {
		f_stage = "qa"
	} else if strings.Contains(name, "preview") {
		f_stage = "preview"
	} else if strings.Contains(name, "stable") {
		f_stage = "stable"
	}
	return f_stage
}

func GetSubstring(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
func GetSubstringColonIndex(str string) int {
	colon := strings.Index(str, ":")
	return colon + 1
}
