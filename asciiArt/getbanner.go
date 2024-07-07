package asciiArt

func BannerFile(option string) string {
	switch option {
	case "standard":
		return "standard.txt"
	case "shadow":
		return "shadow.txt"
	case "thinkertoy":
		return "thinkertoy.txt"
	default:
		return "invalid bannerfile name"
	}
}
