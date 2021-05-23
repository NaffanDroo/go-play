package main

func bold(s string) string {
	return _render(s, "b")
}

func paragraph(s string) string {
	return _render(s, "p")
}

func _render(s string, tag string) string {
	return "<" + tag + ">" + s + "</" + tag + ">"
}
