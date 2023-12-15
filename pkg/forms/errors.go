package forms

type errors map[string][]string

func (e errors) Add(feild,message string){

	e[feild] = append(e[feild],message)

}
//return the first error message
func(e errors) Get(feild string) string{

	es:= e[feild]

	if len(es) == 0 {
		return ""
	} 
	return es[0]



}