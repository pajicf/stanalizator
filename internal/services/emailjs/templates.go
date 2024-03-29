package emailjs

func SendStanalizatorTemplate(ejs *EmailJS, to string, html string) {
	p := StanalizatorTemplateParams{
		ToMail:   to,
		MailHTML: html,
	}

	ejs.SendMail("stanalizator", p)
}
