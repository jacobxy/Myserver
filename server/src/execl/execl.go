package execl

import ()

type ExeclLoad func(file string) (bool, error)

var _execlFile map[string]string = make(map[string]string, 100)

var _execl map[string]ExeclLoad = make(map[string]ExeclLoad)

func RegisterExecl(key string, fn ExeclLoad) {
	_, ok := _execl[key]
	if ok {
		panic(" function alread exist")
	}

	_execl[key] = fn
}

func RegisterExeclFile(key string, file string) {
	_, ok := _execlFile[key]
	if ok {
		panic(" alread exist")
	}

	_execlFile[key] = file
}

//加载没有固定顺序
func LoadExecl() {
	for k, v := range _execlFile {
		if f, ok := _execl[k]; ok {
			f(v)
		}
	}
}
