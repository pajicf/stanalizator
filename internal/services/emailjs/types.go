package emailjs

type StanalizatorTemplateParams struct {
	ToMail   string `json:"to_email"`
	MailHTML string `json:"mail_html"`
}

type SendMailRequestBody struct {
	UserID         string      `json:"user_id"`
	ServiceID      string      `json:"service_id"`
	TemplateID     string      `json:"template_id"`
	TemplateParams interface{} `json:"template_params"`
}
