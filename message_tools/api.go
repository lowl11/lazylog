package message_tools

func Build(message string, layers ...string) string {
	var layersBlock string
	if len(layers) > 0 {
		layersBlock += "["
		for index, layer := range layers {
			layersBlock += layer
			if index != len(layers)-1 {
				layersBlock += ", "
			}
		}
		layersBlock += "] "
	}

	return layersBlock + message
}
