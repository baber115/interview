package main

/*
关于GetPodAction定义，下面赋值正确的是
A. var fragment Fragment = new(GetPodAction)
B. var fragment Fragment = GetPodAction
C. var fragment Fragment = &GetPodAction{}
D. var fragment Fragment = GetPodAction{}

答案：A C D
*/

type Fragment interface {
	Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	return nil
}
