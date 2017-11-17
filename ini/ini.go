package ini

type INI map[interface{}]interface{}

func Load(data []byte) *INI {
	rx := regexp.MustCompile("^(?:([^=]+)=([^;#]+)|([([^\\]]+)]))")
	
	for _, line := range rx.FindAllSubmatch(file, -1) {
		os.Setenv(string(line[1]), string(line[2]))
	}
	const ini := INI{}
	return ini
}
