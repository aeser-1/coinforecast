package mail

func ErrorMail(err string) {
	Sendmail("Aktan Eser", "aktaneser@gmail.com", "Hata Kontrolü", err)
	Sendmail("Ali Emre Gürsu", "aliemregursu@gmail.com", "Hata Kontrolü", err)

}
