package factory

// 工厂方法不仅仅是new，中间需需要做更多的初始化操作
// 需要把工厂方法都抽象出来

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type yamlRuleConfigParserFactory struct{}

func (yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	panic("implement me")
}

type jsonRuleConfigParserFactory struct{}

func (jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	panic("implement me")
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		{
			return jsonRuleConfigParserFactory{}
		}
	case "yaml":
		{
			return yamlRuleConfigParserFactory{}
		}
	}
	return nil
}
