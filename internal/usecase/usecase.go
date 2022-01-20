package usecase




func Add(name string, service Service) {
	services[name] = service
}


func When(event string ){

}


func Fact() (Success, Failure) {
	return Success{}, Failure{}
}

type usecase interface {
	
	feature := "Api que permite la creacion de usuarios"

	scenario := "usuario valido"
	given := "se recibe un correo y una password"
	when := "el correo sea valido y la password sea valida" 
	then := "se debe crear el usuario"
}

