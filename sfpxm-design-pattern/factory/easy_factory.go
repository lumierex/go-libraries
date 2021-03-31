package factory

type IRuleConfigParser interface {
	parse(data []byte)
}

type jsonRuleConfigParser struct{}

func (jsonRuleConfigParser) parse(data []byte) {

}

type yamlRuleConfigParser struct{}

func (yamlRuleConfigParser) parse(data []byte) {
}

func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		{
			return jsonRuleConfigParser{}
		}
	case "yaml":
		{
			return yamlRuleConfigParser{}
		}
	}
	return nil
}
