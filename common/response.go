package common

// Response es la estructura estándar que se usa para devolver respuestas
type Response struct {
	Message    string `json:"message"`
	StatusCode string `json:"statusCode"`
}

// Ok Setea los campos a un estado de respuesta válida
func (s *Response) Ok(mssg string) {

	s.StatusCode = "200"
	s.Message = "ok."
	if mssg != "" {

		s.Message = mssg
	}
}

// Created Setea los campos a un estado de respuesta afirmativa de creación
func (s *Response) Created(mssg string) {

	s.StatusCode = "201"
	s.Message = "created"
	if mssg != "" {

		s.Message = mssg
	}
}

// BadRequest Setea los campos a un estado de respuesta de Error de parte del cliente
func (s *Response) BadRequest(mssg string) {

	s.StatusCode = "400"
	s.Message = mssg
}

// Unauthorized Setea los campos a un estado de respuesta de Error de parte del cliente
func (s *Response) Unauthorized(mssg string) {

	s.StatusCode = "401"
	s.Message = mssg
}

// Forbridden Setea los campos a un estado de respuesta de Error de parte del cliente
func (s *Response) Forbridden(mssg string) {

	s.StatusCode = "403"
	s.Message = mssg
}

// NotFound Setea los campos a un estado de respuesta de Error de parte del cliente
func (s *Response) NotFound(mssg string) {

	s.StatusCode = "404"
	s.Message = mssg
}

// InternalError Setea los campos a un estado de respuesta de error interno
func (s *Response) InternalError(mssg, code string) {

	s.StatusCode = "500"
	s.Message = mssg
	if code != "" {

		s.StatusCode = code
	}
}
