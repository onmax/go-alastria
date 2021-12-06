package hex

func Remove0x(input string) string {
	if len(input) >= 2 && input[0:2] == "0x" {
		return input[2:]
	}
	return input
}
