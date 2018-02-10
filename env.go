package kredivo

type Env int

const (
	//SandBox Env, config for SandBox environment
	SandBox Env = iota

	//Production Env, config for Production environment
	Production

	//Test Env, config for Test environment
	Test
)

//String function, convert Env to String
func (e Env) String() string {
	switch e {
	case SandBox:
		return "https://sandbox.kredivo.com/kredivo/v2"
	case Production:
		return "https://kredivo.com/kredivo/v2"
	default:
		return "https://sandbox.kredivo.com/kredivo/v2"
	}
}
