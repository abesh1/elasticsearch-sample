package elasticsearch

type Config map[string]*systemConf

type systemConf struct {
	URL string `yaml:"url"`
}
